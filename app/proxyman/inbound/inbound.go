package inbound

import (
	"context"
	"sync"

	"github.com/vmessocket/vmessocket/app/proxyman"
	"github.com/vmessocket/vmessocket/common"
	"github.com/vmessocket/vmessocket/common/serial"
	"github.com/vmessocket/vmessocket/common/session"
	"github.com/vmessocket/vmessocket/core"
	"github.com/vmessocket/vmessocket/features/inbound"
)

type Manager struct {
	access          sync.RWMutex
	untaggedHandler []inbound.Handler
	taggedHandlers  map[string]inbound.Handler
	running         bool
}

func New(ctx context.Context, config *proxyman.InboundConfig) (*Manager, error) {
	m := &Manager{
		taggedHandlers: make(map[string]inbound.Handler),
	}
	return m, nil
}

func NewHandler(ctx context.Context, config *core.InboundHandlerConfig) (inbound.Handler, error) {
	rawReceiverSettings, err := config.ReceiverSettings.GetInstance()
	if err != nil {
		return nil, err
	}
	proxySettings, err := config.ProxySettings.GetInstance()
	if err != nil {
		return nil, err
	}
	receiverSettings, ok := rawReceiverSettings.(*proxyman.ReceiverConfig)
	if !ok {
		return nil, newError("not a ReceiverConfig").AtError()
	}
	allocStrategy := receiverSettings.AllocationStrategy
	if allocStrategy == nil || allocStrategy.Type == proxyman.AllocationStrategy_Always {
		return NewAlwaysOnInboundHandler(ctx, receiverSettings, proxySettings)
	}
	return nil, newError("unknown allocation strategy: ", receiverSettings.AllocationStrategy.Type).AtError()
}

func (m *Manager) AddHandler(ctx context.Context, handler inbound.Handler) error {
	m.access.Lock()
	defer m.access.Unlock()
	m.untaggedHandler = append(m.untaggedHandler, handler)
	if m.running {
		return handler.Start()
	}
	return nil
}

func (m *Manager) Close() error {
	m.access.Lock()
	defer m.access.Unlock()
	m.running = false
	var errors []interface{}
	for _, handler := range m.taggedHandlers {
		if err := handler.Close(); err != nil {
			errors = append(errors, err)
		}
	}
	for _, handler := range m.untaggedHandler {
		if err := handler.Close(); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) > 0 {
		return newError("failed to close all handlers").Base(newError(serial.Concat(errors...)))
	}
	return nil
}

func (m *Manager) GetHandler(ctx context.Context, tag string) (inbound.Handler, error) {
	m.access.RLock()
	defer m.access.RUnlock()

	handler, found := m.taggedHandlers[tag]
	if !found {
		return nil, newError("handler not found: ", tag)
	}
	return handler, nil
}

func (m *Manager) RemoveHandler(ctx context.Context, tag string) error {
	if tag == "" {
		return common.ErrNoClue
	}
	m.access.Lock()
	defer m.access.Unlock()
	if handler, found := m.taggedHandlers[tag]; found {
		if err := handler.Close(); err != nil {
			newError("failed to close handler ", tag).Base(err).AtWarning().WriteToLog(session.ExportIDToError(ctx))
		}
		delete(m.taggedHandlers, tag)
		return nil
	}
	return common.ErrNoClue
}

func (m *Manager) Start() error {
	m.access.Lock()
	defer m.access.Unlock()
	m.running = true
	for _, handler := range m.taggedHandlers {
		if err := handler.Start(); err != nil {
			return err
		}
	}
	for _, handler := range m.untaggedHandler {
		if err := handler.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (*Manager) Type() interface{} {
	return inbound.ManagerType()
}

func init() {
	common.Must(common.RegisterConfig((*proxyman.InboundConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return New(ctx, config.(*proxyman.InboundConfig))
	}))
	common.Must(common.RegisterConfig((*core.InboundHandlerConfig)(nil), func(ctx context.Context, config interface{}) (interface{}, error) {
		return NewHandler(ctx, config.(*core.InboundHandlerConfig))
	}))
}
