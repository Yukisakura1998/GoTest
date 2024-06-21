package main

import (
	"fmt"
	"zinx/zinx/ziface"
	"zinx/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
	//IPAddress net.IPAddr
}
type HelloRouter struct {
	znet.BaseRouter
	//IPAddress net.IPAddr
}

func (r *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("call Hello")
	//先读数据
	err := request.GetConnection().Send(201, []byte("Hello\r\n"))
	if err != nil {
		return
	}

}
func (r *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call ping")
	//先读数据
	err := request.GetConnection().Send(200, []byte("ping\r\n"))
	if err != nil {
		return
	}
}

func main() {
	s := znet.NewServer()
	//添加自定义router
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	//启动
	s.Server()
}
