package main

import (
	"fmt"
	"net"
	"zinx/zinx/ziface"
	"zinx/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
	IPAddress net.IPAddr
}

func (r *PingRouter) PreHandle(request ziface.IRequest) {
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping before\r\n"))
	if err != nil {
		fmt.Println("call back before error : ", err)
		return
	}
}

func (r *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call ping")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping\r\n"))
	if err != nil {
		fmt.Println("call ping error : ", err)
		return
	}
}

func (r *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("call ping after")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping after\r\n"))
	if err != nil {
		fmt.Println("call ping after error : ", err)
		return
	}
}

func main() {
	s := znet.NewServer()
	//添加自定义router
	s.AddRouter(&PingRouter{})
	//启动
	s.Server()
}
