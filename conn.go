package fdp

import "net"

const (
	numRight  = 32
	sizeRight = numRight + 64

	queueSize = 1024
	mtu       = 1280
)

type UDPConn struct {
	conn *net.UDPConn
}

func newUDPConn(conn *net.UDPConn) *UDPConn {
	return &UDPConn{
		conn: conn,
	}
}

func (c *UDPConn) Read(p []byte) (n int, err error) {

	return
}

func (c *UDPConn) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (c *UDPConn) Close() error {
	return c.conn.Close()
}

func (c *UDPConn) recvLoop() {

}

func (c *UDPConn) sendLoop() {

}
