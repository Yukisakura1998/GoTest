package znet

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"zinx/zinx/utils"
	"zinx/zinx/ziface"
)

// Connection 连接模块
type Connection struct {
	Server ziface.IServer
	//socket
	Conn *net.TCPConn
	//ID
	ConnID uint32
	// status
	isClosed bool
	//exit
	ExitChan chan bool
	//msgChan
	msgChan chan []byte
	//msgBuffChan
	msgBuffChan chan []byte
	//MsgHandler
	MessageHandler ziface.IMassageHandle
	//
	property map[string]interface{}

	propertyLock sync.RWMutex
}

// Start : start connection
func (c *Connection) Start() {
	fmt.Println("Conn Start() ..ConnID =", c.ConnID)
	//启动业务
	go c.Write()
	go c.Reader()

	c.Server.CallOnConnStart(c)

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

	c.Server.CallOnConnStop(c)

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

	c.Server.GetConnManager().Remove(c)

	close(c.ExitChan)
	close(c.msgChan)

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
func (c *Connection) Send(msgId uint32, data []byte) error {

	if c.isClosed == true {
		return errors.New("connection is closed")
	}
	pkg := NewPackage()

	sendMsg, err := pkg.Pack(NewMessage(msgId, data))

	if err != nil {
		return err
	}

	//if _, err := c.Conn.Write(sendMsg); err != nil {
	//	return err
	//}
	c.msgChan <- sendMsg

	return nil
}
func (c *Connection) SendBuff(msgId uint32, data []byte) error {
	if c.isClosed == true {
		return errors.New("connection is closed")
	}
	pkg := NewPackage()

	sendMsg, err := pkg.Pack(NewMessage(msgId, data))

	if err != nil {
		return err
	}

	c.msgBuffChan <- sendMsg

	return nil
}

// NewConnection : new a connection
func NewConnection(s ziface.IServer, conn *net.TCPConn, connId uint32, msgHand ziface.IMassageHandle) *Connection {
	c := &Connection{
		Server:         s,
		Conn:           conn,
		ConnID:         connId,
		MessageHandler: msgHand,
		isClosed:       false,
		msgChan:        make(chan []byte),
		ExitChan:       make(chan bool, 1),
		msgBuffChan:    make(chan []byte, utils.GlobalObject.MaxMsgChanLen),
		property:       map[string]interface{}{},
	}
	c.Server.GetConnManager().Add(c)
	return c
}

func (c *Connection) Write() {
	fmt.Println("Write is running ..ConnID =", c.ConnID)
	defer fmt.Println("Write is exit ..ConnID =", c.ConnID, ", remote address is ", c.RemoteAddr())
	defer c.Stop()

	for {
		select {
		case data := <-c.msgChan:
			if _, err := c.Conn.Write(data); err != nil {
				fmt.Println("Send is error ..ConnID =", c.ConnID, " error: ", err)
			}
		case data, ok := <-c.msgBuffChan:
			if ok {
				//有数据要写给客户端
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send Buff Data error:, ", err, " Conn Writer exit")
					return
				}
			} else {
				fmt.Println("msgBuffChan is Closed")
				break
			}
		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) Reader() {
	fmt.Println("Read is running ..ConnID =", c.ConnID)
	defer fmt.Println("Read is exit ..ConnID =", c.ConnID, ", remote address is ", c.RemoteAddr())
	defer c.Stop()
	for {
		pkg := NewPackage()
		head := make([]byte, pkg.GetHeadLen())
		_, err := io.ReadFull(c.GetTCPConnection(), head)
		if err != nil {
			fmt.Println("read head err:", err.Error())
			c.ExitChan <- true
			continue
		}
		msg, err := pkg.Unpack(head)
		if err != nil {
			fmt.Println("unpack head err:", err.Error())
			c.ExitChan <- true
			continue
		}
		var data []byte
		if msg.GetMsgLen() > 0 {
			data = make([]byte, msg.GetMsgLen())
			_, err := io.ReadFull(c.GetTCPConnection(), data)
			if err != nil {
				fmt.Println("read data err:", err.Error())
				c.ExitChan <- true
				continue
			}
		}
		msg.SetMsgData(data)

		//获取request
		req := Request{
			conn:    c,
			message: msg,
		}

		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MessageHandler.SendMsgToTaskQueue(&req)
		} else {
			go c.MessageHandler.DoMessageHandler(&req)
		}

		////通过handAPI操作这个buff，这个API就是New的时候传进来的API，现在new的时候传进来的是，server里的CallBackHandleFunc
		//if err := c.handAPI(c.Conn, buff, cnt); err != nil {
		//	fmt.Println("Handler error : ", err, " ,ConnID =", c.ConnID)
		//	c.ExitChan <- true
		//	break
		//}
	}
}
func (c *Connection) GetProp(key string) (interface{}, error) {
	c.propertyLock.RLock()
	defer c.propertyLock.RUnlock()

	if value, ok := c.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("no property found")
	}
}

// SetProp :set prop
func (c *Connection) SetProp(key string, value interface{}) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	c.property[key] = value
}

// RemoveProp :remove prop
func (c *Connection) RemoveProp(key string) {
	c.propertyLock.Lock()
	defer c.propertyLock.Unlock()

	delete(c.property, key)
}
