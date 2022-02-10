package inbound

import (
	"context"

	"github.com/vmessocket/vmessocket/app/proxyman"
	"github.com/vmessocket/vmessocket/common"
	"github.com/vmessocket/vmessocket/common/dice"
	"github.com/vmessocket/vmessocket/common/errors"
	"github.com/vmessocket/vmessocket/common/mux"
	"github.com/vmessocket/vmessocket/common/net"
	"github.com/vmessocket/vmessocket/proxy"
	"github.com/vmessocket/vmessocket/transport/internet"
)

type AlwaysOnInboundHandler struct {
	proxy   proxy.Inbound
	workers []worker
	tag     string
}

func NewAlwaysOnInboundHandler(ctx context.Context, tag string, receiverConfig *proxyman.ReceiverConfig, proxyConfig interface{}) (*AlwaysOnInboundHandler, error) {
	rawProxy, err := common.CreateObject(ctx, proxyConfig)
	if err != nil {
		return nil, err
	}
	p, ok := rawProxy.(proxy.Inbound)
	if !ok {
		return nil, newError("not an inbound proxy.")
	}
	h := &AlwaysOnInboundHandler{
		proxy: p,
		tag:   tag,
	}
	nl := p.Network()
	pr := receiverConfig.PortRange
	address := receiverConfig.Listen.AsAddress()
	if address == nil {
		address = net.AnyIP
	}
	mss, err := internet.ToMemoryStreamConfig(receiverConfig.StreamSettings)
	if err != nil {
		return nil, newError("failed to parse stream config").Base(err).AtWarning()
	}
	if receiverConfig.ReceiveOriginalDestination {
		if mss.SocketSettings == nil {
			mss.SocketSettings = &internet.SocketConfig{}
		}
		if mss.SocketSettings.Tproxy == internet.SocketConfig_Off {
			mss.SocketSettings.Tproxy = internet.SocketConfig_Redirect
		}
		mss.SocketSettings.ReceiveOriginalDestAddress = true
	}
	if pr == nil {
		if net.HasNetwork(nl, net.Network_UNIX) {
			newError("creating unix domain socket worker on ", address).AtDebug().WriteToLog()
			worker := &dsWorker{
				address:        address,
				proxy:          p,
				stream:         mss,
				tag:            tag,
				sniffingConfig: receiverConfig.GetEffectiveSniffingSettings(),
				ctx:            ctx,
			}
			h.workers = append(h.workers, worker)
		}
	}
	if pr != nil {
		for port := pr.From; port <= pr.To; port++ {
			if net.HasNetwork(nl, net.Network_TCP) {
				newError("creating stream worker on ", address, ":", port).AtDebug().WriteToLog()
				worker := &tcpWorker{
					address:        address,
					port:           net.Port(port),
					proxy:          p,
					stream:         mss,
					recvOrigDest:   receiverConfig.ReceiveOriginalDestination,
					tag:            tag,
					sniffingConfig: receiverConfig.GetEffectiveSniffingSettings(),
					ctx:            ctx,
				}
				h.workers = append(h.workers, worker)
			}
			if net.HasNetwork(nl, net.Network_UDP) {
				worker := &udpWorker{
					ctx:            ctx,
					tag:            tag,
					proxy:          p,
					address:        address,
					port:           net.Port(port),
					sniffingConfig: receiverConfig.GetEffectiveSniffingSettings(),
					stream:         mss,
				}
				h.workers = append(h.workers, worker)
			}
		}
	}
	return h, nil
}

func (h *AlwaysOnInboundHandler) Close() error {
	var errs []error
	for _, worker := range h.workers {
		errs = append(errs, worker.Close())
	}
	if err := errors.Combine(errs...); err != nil {
		return newError("failed to close all resources").Base(err)
	}
	return nil
}

func (h *AlwaysOnInboundHandler) GetInbound() proxy.Inbound {
	return h.proxy
}

func (h *AlwaysOnInboundHandler) GetRandomInboundProxy() (interface{}, net.Port, int) {
	if len(h.workers) == 0 {
		return nil, 0, 0
	}
	w := h.workers[dice.Roll(len(h.workers))]
	return w.Proxy(), w.Port(), 9999
}

func (h *AlwaysOnInboundHandler) Start() error {
	for _, worker := range h.workers {
		if err := worker.Start(); err != nil {
			return err
		}
	}
	return nil
}

func (h *AlwaysOnInboundHandler) Tag() string {
	return h.tag
}
