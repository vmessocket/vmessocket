//go:build !linux && !freebsd
// +build !linux,!freebsd

package udp

import "github.com/vmessocket/vmessocket/common/net"

func ReadUDPMsg(conn *net.UDPConn, payload []byte, oob []byte) (int, int, int, *net.UDPAddr, error) {
	nBytes, addr, err := conn.ReadFromUDP(payload)
	return nBytes, 0, 0, addr, err
}

func RetrieveOriginalDest(oob []byte) net.Destination {
	return net.Destination{}
}
