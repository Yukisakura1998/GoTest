package znet

import (
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
	Router    ziface.IRouter
}

//// CallBackHandleFunc 定义一个API，后改自定义
//func CallBackHandleFunc(conn *net.TCPConn, data []byte, length int) error {
//	fmt.Println("[Connection HandelFunc callback] ")
//	if _, err := conn.Write(data[:length]); err != nil {
//		fmt.Println("callback error")
//		return errors.New("CallBack error")
//	}
//	return nil
//}

func (s *Server) Start() {
	fmt.Println("Server is starting!")
	go func() {
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IPAddress, s.Port))
		if err != nil {
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start Zinx server  ", s.Name, " success, now listening...")

		var cid uint32
		cid = 0
		//阻塞
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}

			//主要的入口
			delConn := NewConnection(conn, cid, s.Router)
			cid++
			go delConn.Start()
			/*
				go func() {
					for {
						buff := make([]byte, 512)
						read, err := conn.Read(buff)
						if err != nil {
							continue
						}
						fmt.Printf("receive:%s", buff)
						if _, err := conn.Write(buff[:read]); err != nil {
							continue
						}

					}
				}()*/
			//s.conn = conn
		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)
}

func (s *Server) Server() {
	s.Start()
	select {}
}
func (s *Server) AddRouter(router ziface.IRouter) {
	s.Router = router
}
func NewServer(name string) ziface.IServer {
	s := &Server{
		IPVersion: "tcp4",
		IPAddress: "0.0.0.0",
		Port:      8080,
		Name:      name,
		Router:    nil,
	}
	return s
}
