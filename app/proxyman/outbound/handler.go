package outbound

import (
	"context"

	"github.com/vmessocket/vmessocket/app/proxyman"
	"github.com/vmessocket/vmessocket/common"
	"github.com/vmessocket/vmessocket/common/net"
	"github.com/vmessocket/vmessocket/common/session"
	"github.com/vmessocket/vmessocket/core"
	"github.com/vmessocket/vmessocket/features/outbound"
	"github.com/vmessocket/vmessocket/proxy"
	"github.com/vmessocket/vmessocket/transport"
	"github.com/vmessocket/vmessocket/transport/internet"
)

type Handler struct {
	senderSettings  *proxyman.SenderConfig
	streamSettings  *internet.MemoryStreamConfig
	proxy           proxy.Outbound
	outboundManager outbound.Manager
}

func NewHandler(ctx context.Context, config *core.OutboundHandlerConfig) (outbound.Handler, error) {
	v := core.MustFromContext(ctx)
	h := &Handler{
		outboundManager: v.GetFeature(outbound.ManagerType()).(outbound.Manager),
	}
	if config.SenderSettings != nil {
		senderSettings, err := config.SenderSettings.GetInstance()
		if err != nil {
			return nil, err
		}
		switch s := senderSettings.(type) {
		case *proxyman.SenderConfig:
			h.senderSettings = s
			mss, err := internet.ToMemoryStreamConfig(s.StreamSettings)
			if err != nil {
				return nil, newError("failed to parse stream settings").Base(err).AtWarning()
			}
			h.streamSettings = mss
		default:
			return nil, newError("settings is not SenderConfig")
		}
	}
	proxyConfig, err := config.ProxySettings.GetInstance()
	if err != nil {
		return nil, err
	}
	rawProxyHandler, err := common.CreateObject(ctx, proxyConfig)
	if err != nil {
		return nil, err
	}
	proxyHandler, ok := rawProxyHandler.(proxy.Outbound)
	if !ok {
		return nil, newError("not an outbound handler")
	}
	h.proxy = proxyHandler
	return h, nil
}

func (h *Handler) Address() net.Address {
	if h.senderSettings == nil || h.senderSettings.Via == nil {
		return nil
	}
	return h.senderSettings.Via.AsAddress()
}

func (h *Handler) Close() error {
	return nil
}

func (h *Handler) Dial(ctx context.Context, dest net.Destination) (internet.Connection, error) {
	if h.senderSettings != nil {
		if h.senderSettings.Via != nil {
			outbound := session.OutboundFromContext(ctx)
			if outbound == nil {
				outbound = new(session.Outbound)
				ctx = session.ContextWithOutbound(ctx, outbound)
			}
			outbound.Gateway = h.senderSettings.Via.AsAddress()
		}
	}
	conn, err := internet.Dial(ctx, dest, h.streamSettings)
	return conn, err
}

func (h *Handler) Dispatch(ctx context.Context, link *transport.Link) {
	if err := h.proxy.Process(ctx, link, h); err != nil {
		err := newError("failed to process outbound traffic").Base(err)
		session.SubmitOutboundErrorToOriginator(ctx, err)
		err.WriteToLog(session.ExportIDToError(ctx))
		common.Interrupt(link.Writer)
	} else {
		common.Must(common.Close(link.Writer))
	}
	common.Interrupt(link.Reader)
}

func (h *Handler) GetOutbound() proxy.Outbound {
	return h.proxy
}

func (h *Handler) Start() error {
	return nil
}
