package ziface

import "net"

// IConnection 连接模块抽象层
type IConnection interface {
	// Start : start connection
	Start()
	// Stop : stop connection
	Stop()
	// GetTCPConnection : get connection socket
	GetTCPConnection() *net.TCPConn
	// GetConnectID : get connection id
	GetConnectID() uint32
	// RemoteAddr : get client status
	RemoteAddr() net.Addr
	// Send : send message
	Send(msgId uint32, data []byte) error
	// SendBuff : have buffer
	SendBuff(msgId uint32, data []byte) error

	// GetProp :get prop
	GetProp(key string) (interface{}, error)
	// SetProp :set prop
	SetProp(key string, value interface{})
	// RemoveProp :remove prop
	RemoveProp(key string)
}

// HandleFunc resolve function
type HandleFunc func(*net.TCPConn, []byte, int) error
