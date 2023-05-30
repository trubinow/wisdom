package servers

import (
	"fmt"
	"net"
	"time"
)

type Handler interface {
	Handle(c net.Conn)
}

type TCPServer struct {
	handler               Handler
	tcpConnectionDeadline time.Duration
}

func NewTCPServer(handler Handler, tcpConnectionDeadline time.Duration) TCPServer {
	return TCPServer{
		handler:               handler,
		tcpConnectionDeadline: tcpConnectionDeadline,
	}
}

func (s TCPServer) Run(port int) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	defer listener.Close()

	fmt.Println("Listening on port:", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		if err = conn.SetDeadline(time.Now().Add(s.tcpConnectionDeadline)); err != nil {
			return err
		}

		go s.handler.Handle(conn)
	}
}

func (s TCPServer) Shutdown() error {
	return nil
}
