package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/zinx/ziface"
)

type Server struct {
	IPVersion string
	IPAddress string
	Port      int
	Name      string
	conn      net.Conn
}

// CallBackHandleFunc 定义一个API，后改自定义
func CallBackHandleFunc(conn *net.TCPConn, data []byte, length int) error {
	fmt.Println("[Connection HandelFunc callback] ")
	if _, err := conn.Write(data[:length]); err != nil {
		fmt.Println("callback error")
		return errors.New("CallBack error")
	}
	fmt.Printf("[Connection HandelFunc callback] data :%s", data)

	return nil
}

func (s *Server) Start() {
	fmt.Println("Server is starting!")
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IPAddress, s.Port))
		if err != nil {
			return
		}

		tcp, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			return
		}

		var cid uint32
		cid = 0
		//阻塞
		for {
			acceptTCP, err := tcp.AcceptTCP()
			if err != nil {
				continue
			}

			delConn := NewConnection(acceptTCP, cid, CallBackHandleFunc)
			cid++
			go delConn.Start()
			/*
				go func() {
					for {
						buff := make([]byte, 512)
						read, err := acceptTCP.Read(buff)
						if err != nil {
							continue
						}
						fmt.Printf("receive:%s", buff)
						if _, err := acceptTCP.Write(buff[:read]); err != nil {
							continue
						}

					}
				}()*/
			//s.conn = acceptTCP
		}
	}()
}

func (s *Server) Stop() {

}

func (s *Server) Server() {
	s.Start()
	select {}
}

func NewServer(name string) ziface.IServer {
	s := &Server{
		IPVersion: "tcp4",
		IPAddress: "0.0.0.0",
		Port:      8080,
		Name:      name,
	}
	return s
}
