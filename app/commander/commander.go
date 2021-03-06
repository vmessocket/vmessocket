package commander

import (
	"context"
	"net"
	"sync"

	"google.golang.org/grpc"

	"github.com/vmessocket/vmessocket/common"
	"github.com/vmessocket/vmessocket/common/signal/done"
	"github.com/vmessocket/vmessocket/features/outbound"
)

type Commander struct {
	sync.Mutex
	ohm      outbound.Manager
	server   *grpc.Server
	services []Service
}

func NewCommander(ctx context.Context, config *Config) (*Commander, error) {
	c := &Commander{}
	for _, rawConfig := range config.Service {
		config, err := rawConfig.GetInstance()
		if err != nil {
			return nil, err
		}
		rawService, err := common.CreateObject(ctx, config)
		if err != nil {
			return nil, err
		}
		service, ok := rawService.(Service)
		if !ok {
			return nil, newError("not a Service.")
		}
		c.services = append(c.services, service)
	}
	return c, nil
}

func (c *Commander) Close() error {
	c.Lock()
	defer c.Unlock()
	if c.server != nil {
		c.server.Stop()
		c.server = nil
	}
	return nil
}

func (c *Commander) Start() error {
	c.Lock()
	c.server = grpc.NewServer()
	for _, service := range c.services {
		service.Register(c.server)
	}
	c.Unlock()
	listener := &OutboundListener{
		buffer: make(chan net.Conn, 4),
		done:   done.New(),
	}
	go func() {
		if err := c.server.Serve(listener); err != nil {
			newError("failed to start grpc server").Base(err).AtError().WriteToLog()
		}
	}()
	return c.ohm.AddHandler(context.Background(), &Outbound{
		listener: listener,
	})
}

func (c *Commander) Type() interface{} {
	return (*Commander)(nil)
}

func init() {
	common.Must(common.RegisterConfig((*Config)(nil), func(ctx context.Context, cfg interface{}) (interface{}, error) {
		return NewCommander(ctx, cfg.(*Config))
	}))
}
