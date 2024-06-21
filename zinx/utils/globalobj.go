package utils

import (
	"encoding/json"
	"os"
	"zinx/zinx/ziface"
)

type GlobalObj struct {
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string

	Version        string
	MaxConn        int
	MaxPackageSize uint32
	//工作数量
	WorkerPoolSize uint32
	//系统允许的最大值
	MaxWorkerPoolSize uint32

	MaxMsgChanLen uint32
}

var GlobalObject *GlobalObj

func (g *GlobalObj) LoadJson() {
	file, err := os.ReadFile("config/zinx.json")
	if err != nil {
		return
	}
	err2 := json.Unmarshal(file, &g)
	if err2 != nil {
		return
	}
}
func init() {
	//默认值
	GlobalObject = &GlobalObj{
		Name:              "zinx",
		Version:           "1.0.0",
		TcpPort:           8080,
		Host:              "0.0.0.0",
		MaxConn:           10,
		MaxPackageSize:    1024,
		WorkerPoolSize:    10,
		MaxWorkerPoolSize: 1024,
		MaxMsgChanLen:     1024,
	}
	//加载自定义参数
	GlobalObject.LoadJson()
}
