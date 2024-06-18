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
	//exit
	ExitChan chan bool
	//router
	Router ziface.IRouter
}

// Start : start connection
func (c *Connection) Start() {
	fmt.Println("Conn Start() ..ConnID =", c.ConnID)
	//启动业务
	go c.Reader()

	for {
		select {
		case <-c.ExitChan:
			//得到退出消息，不再阻塞
			return
		}
	}
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

	c.ExitChan <- true

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
	return c.Conn.RemoteAddr()
}

// Send : send message
func (c *Connection) Send(data []byte) error {
	return nil
}

// NewConnection : new a connection
func NewConnection(conn *net.TCPConn, connId uint32, router ziface.IRouter) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnID:   connId,
		Router:   router,
		isClosed: false,
		ExitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) Reader() {
	fmt.Println("Read is running ..ConnID =", c.ConnID)
	defer fmt.Println("Read is exit ..ConnID =", c.ConnID, ", remote address is ", c.RemoteAddr())
	defer c.Stop()
	for {
		//cnt in buff,读取字符到buff中
		buff := make([]byte, 4096)
		_, err := c.Conn.Read(buff)
		if err != nil {
			fmt.Println("Read error", err)
			c.ExitChan <- true
			continue
		}

		//获取request
		req := Request{
			conn: c,
			data: buff,
		}

		go func(request ziface.IRequest) {
			//调用路由
			c.Router.PreHandle(request)
			c.Router.MainHandle(request)
			c.Router.PostHandle(request)
		}(&req)

		////通过handAPI操作这个buff，这个API就是New的时候传进来的API，现在new的时候传进来的是，server里的CallBackHandleFunc
		//if err := c.handAPI(c.Conn, buff, cnt); err != nil {
		//	fmt.Println("Handler error : ", err, " ,ConnID =", c.ConnID)
		//	c.ExitChan <- true
		//	break
		//}
	}
}
