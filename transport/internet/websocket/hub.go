package websocket

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"

	"github.com/vmessocket/vmessocket/common"
	"github.com/vmessocket/vmessocket/common/net"
	http_proto "github.com/vmessocket/vmessocket/common/protocol/http"
	"github.com/vmessocket/vmessocket/common/session"
	"github.com/vmessocket/vmessocket/transport/internet"
)

var upgrader = &websocket.Upgrader{
	ReadBufferSize:   4 * 1024,
	WriteBufferSize:  4 * 1024,
	HandshakeTimeout: time.Second * 4,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Listener struct {
	sync.Mutex
	server   http.Server
	listener net.Listener
	config   *Config
	addConn  internet.ConnHandler
}

type requestHandler struct {
	path                string
	ln                  *Listener
	earlyDataEnabled    bool
	earlyDataHeaderName string
}

func ListenWS(ctx context.Context, address net.Address, port net.Port, streamSettings *internet.MemoryStreamConfig, addConn internet.ConnHandler) (internet.Listener, error) {
	l := &Listener{
		addConn: addConn,
	}
	wsSettings := streamSettings.ProtocolSettings.(*Config)
	l.config = wsSettings
	var listener net.Listener
	var err error
	listener, err = internet.ListenSystem(ctx, &net.TCPAddr{
		IP:   address.IP(),
		Port: int(port),
	}, streamSettings.SocketSettings)
	if err != nil {
		return nil, newError("failed to listen TCP(for WS) on ", address, ":", port).Base(err)
	}
	newError("listening TCP(for WS) on ", address, ":", port).WriteToLog(session.ExportIDToError(ctx))
	l.listener = listener
	useEarlyData := false
	earlyDataHeaderName := ""
	l.server = http.Server{
		Handler: &requestHandler{
			path:                wsSettings.GetNormalizedPath(),
			ln:                  l,
			earlyDataEnabled:    useEarlyData,
			earlyDataHeaderName: earlyDataHeaderName,
		},
		ReadHeaderTimeout: time.Second * 4,
		MaxHeaderBytes:    http.DefaultMaxHeaderBytes,
	}
	return nil, err
}

func (ln *Listener) Addr() net.Addr {
	return ln.listener.Addr()
}

func (h *requestHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	var earlyData io.Reader
	if !h.earlyDataEnabled {
		if request.URL.Path != h.path {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
	} else if h.earlyDataHeaderName != "" {
		if request.URL.Path != h.path {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
		earlyDataStr := request.Header.Get(h.earlyDataHeaderName)
		earlyData = base64.NewDecoder(base64.RawURLEncoding, bytes.NewReader([]byte(earlyDataStr)))
	} else {
		if strings.HasPrefix(request.URL.RequestURI(), h.path) {
			earlyDataStr := request.URL.RequestURI()[len(h.path):]
			earlyData = base64.NewDecoder(base64.RawURLEncoding, bytes.NewReader([]byte(earlyDataStr)))
		} else {
			writer.WriteHeader(http.StatusNotFound)
			return
		}
	}
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		newError("failed to convert to WebSocket connection").Base(err).WriteToLog()
		return
	}
	forwardedAddrs := http_proto.ParseXForwardedFor(request.Header)
	remoteAddr := conn.RemoteAddr()
	if len(forwardedAddrs) > 0 && forwardedAddrs[0].Family().IsIP() {
		remoteAddr = &net.TCPAddr{
			IP:   forwardedAddrs[0].IP(),
			Port: int(0),
		}
	}
	if earlyData == nil {
		h.ln.addConn(newConnection(conn, remoteAddr))
	} else {
		h.ln.addConn(newConnectionWithEarlyData(conn, remoteAddr, earlyData))
	}
}

func init() {
	common.Must(internet.RegisterTransportListener(protocolName, ListenWS))
}
