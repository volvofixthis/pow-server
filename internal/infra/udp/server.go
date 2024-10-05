package udp

import (
	"fmt"
	"net"
)

// UDPConnWrapper wraps a UDP connection to implement Conn interface
type UDPConnWrapper struct {
	conn *net.UDPConn
	addr *net.UDPAddr
}

func (w *UDPConnWrapper) Read(p []byte) (int, error) {
	n, addr, err := w.conn.ReadFromUDP(p)
	w.addr = addr
	return n, err
}

func (w *UDPConnWrapper) Write(p []byte) (int, error) {
	if w.addr == nil {
		return 0, fmt.Errorf("no address to write to")
	}
	return w.conn.WriteToUDP(p, w.addr)
}

func (w *UDPConnWrapper) Close() error {
	return w.conn.Close()
}
