package tcp

import (
	"context"
	gotls "crypto/tls"
	"strings"
	"time"

	"github.com/vmessocket/vmessocket/common"
	"github.com/vmessocket/vmessocket/common/net"
	"github.com/vmessocket/vmessocket/common/session"
	"github.com/vmessocket/vmessocket/transport/internet"
)

type Listener struct {
	listener   net.Listener
	tlsConfig  *gotls.Config
	config     *Config
	addConn    internet.ConnHandler
}

func ListenTCP(ctx context.Context, address net.Address, port net.Port, streamSettings *internet.MemoryStreamConfig, handler internet.ConnHandler) (internet.Listener, error) {
	l := &Listener{
		addConn: handler,
	}
	tcpSettings := streamSettings.ProtocolSettings.(*Config)
	l.config = tcpSettings
	var listener net.Listener
	var err error
	listener, err = internet.ListenSystem(ctx, &net.TCPAddr{
		IP:   address.IP(),
		Port: int(port),
	}, streamSettings.SocketSettings)
	if err != nil {
		return nil, newError("failed to listen TCP on ", address, ":", port).Base(err)
	}
	newError("listening TCP on ", address, ":", port).WriteToLog(session.ExportIDToError(ctx))
	l.listener = listener
	go l.keepAccepting()
	return nil, nil
}

func (v *Listener) Addr() net.Addr {
	return v.listener.Addr()
}

func (v *Listener) keepAccepting() {
	for {
		conn, err := v.listener.Accept()
		if err != nil {
			errStr := err.Error()
			if strings.Contains(errStr, "closed") {
				break
			}
			newError("failed to accepted raw connections").Base(err).AtWarning().WriteToLog()
			if strings.Contains(errStr, "too many") {
				time.Sleep(time.Millisecond * 500)
			}
			continue
		}
		v.addConn(internet.Connection(conn))
	}
}

func init() {
	common.Must(internet.RegisterTransportListener(protocolName, ListenTCP))
}
