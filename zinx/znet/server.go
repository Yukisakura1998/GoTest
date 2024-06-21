package znet

import (
	"fmt"
	"net"
	"zinx/zinx/utils"
	"zinx/zinx/ziface"
)

type Server struct {
	IPVersion      string
	IPAddress      string
	Port           int
	Name           string
	conn           net.Conn
	MessageHandler ziface.IMassageHandle
	ConnManager    ziface.IConnectManager
	OnConnStart    func(conn ziface.IConnection)
	OnConnStop     func(conn ziface.IConnection)
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
		s.MessageHandler.StartWorkerPool()

		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IPAddress, s.Port))
		if err != nil {
			return
		}

		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen", s.IPVersion, "err", err)
			return
		}

		fmt.Println("start Zinx server ", addr, " success, now listening...")

		var cid uint32
		cid = 0
		//阻塞
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err ", err)
				continue
			}
			//判断连接最大值
			if s.ConnManager.Count() >= utils.GlobalObject.MaxConn {
				err := conn.Close()
				if err != nil {
					continue
				}
				continue
			}

			//主要的入口
			delConn := NewConnection(s, conn, cid, s.MessageHandler)
			cid++
			go delConn.Start()

		}
	}()
}

func (s *Server) Stop() {
	fmt.Println("[STOP] Zinx server , name ", s.Name)
	s.ConnManager.Clear()
}

func (s *Server) Server() {
	s.Start()
	select {}
}
func (s *Server) AddRouter(msgId uint32, router ziface.IRouter) {
	s.MessageHandler.AddRouter(msgId, router)
	fmt.Println("[ADD] Zinx router , name ", s.Name, " success, id ", msgId)
}
func (s *Server) GetConnManager() ziface.IConnectManager {
	return s.ConnManager
}

func (s *Server) SetOnConnStart(funcName func(ziface.IConnection)) {
	s.OnConnStart = funcName
}

// SetOnConnStop 设置该Server的连接断开时的Hook函数
func (s *Server) SetOnConnStop(funcName func(ziface.IConnection)) {
	s.OnConnStop = funcName
}

// CallOnConnStart 调用连接OnConnStart Hook函数
func (s *Server) CallOnConnStart(conn ziface.IConnection) {
	if s.OnConnStart != nil {
		s.OnConnStart(conn)
	}
}

// CallOnConnStop 调用连接OnConnStop Hook函数
func (s *Server) CallOnConnStop(conn ziface.IConnection) {
	if s.OnConnStop != nil {
		s.OnConnStop(conn)
	}
}

func NewServer() ziface.IServer {
	utils.GlobalObject.LoadJson()
	s := &Server{
		IPVersion:      "tcp4",
		IPAddress:      utils.GlobalObject.Host,
		Port:           utils.GlobalObject.TcpPort,
		Name:           utils.GlobalObject.Name,
		MessageHandler: NewMessageHandler(),
		ConnManager:    NewConnectManager(),
	}
	return s
}
