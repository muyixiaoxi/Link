package transmit

import "net"

func InitConnect() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", "127.0.0.1:8333")
	return
}
