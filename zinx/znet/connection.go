package znet

import (
	"fmt"
	"net"
	"zinx/zinx/ziface"
)

// Connection 连接模块
type Connection struct {
	//socket
	Conn *net.TCPConn
	//ID
	ConnID uint32
	// status
	isClosed bool
	//func
	handAPI ziface.HandleFunc
	//exit
	ExitChan chan bool
}

func (c *Connection) Reader() {
	fmt.Println("Read is running ..ConnID =", c.ConnID)
	defer fmt.Println("Read is exit ..ConnID =", c.ConnID, ", remote address is ", c.RemoteAddr())
	defer c.Stop()
	for {
		//read in buff
		buff := make([]byte, 512)
		read, err := c.Conn.Read(buff)
		if err != nil {
			fmt.Println("Read error", err)
			continue
		}

		if err := c.handAPI(c.Conn, buff, read); err != nil {
			fmt.Println("Handler error : ", err, " ,ConnID =", c.ConnID)
			break
		}
	}
}

// Start : start connection
func (c *Connection) Start() {
	fmt.Println("Conn Start() ..ConnID =", c.ConnID)
	//启动业务
	go c.Reader()
}

// Stop : stop connection
func (c *Connection) Stop() {
	fmt.Println("Conn Stop() ..ConnID =", c.ConnID)
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	err := c.Conn.Close()
	if err != nil {
		return
	}

	close(c.ExitChan)
}

// GetTCPConnection : get connection socket
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnectID : get connection id
func (c *Connection) GetConnectID() uint32 {
	return c.ConnID
}

// RemoteAddr : get client status
func (c *Connection) RemoteAddr() net.Addr {
	return c.RemoteAddr()
}

// Send : send message
func (c *Connection) Send(data []byte) error {
	return nil
}

func NewConnection(conn *net.TCPConn, connId uint32, callbackAPI ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connId,
		handAPI:  callbackAPI,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}
