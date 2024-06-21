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
func InitGlobalObj() {
	//默认值
	GlobalObject = &GlobalObj{
		Name:           "zinx",
		Version:        "0.0.5",
		TcpPort:        9000,
		Host:           "127.0.0.1",
		MaxConn:        10,
		MaxPackageSize: 1024,
	}
	//加载自定义参数
	GlobalObject.LoadJson()
}
