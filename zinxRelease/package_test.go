package znet

import (
	"fmt"
	"io"
	"net"
	"testing"
)

func TestDataPack(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		return
	}

	go func() {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		go func(conn net.Conn) {
			pkg := NewPackage()
			for {
				head := make([]byte, pkg.GetHeadLen())
				_, err := io.ReadFull(conn, head)
				if err != nil {
					break
				}
				msgHead, err := pkg.Unpack(head)
				if err != nil {
					return
				}
				if msgHead.GetMsgLen() > 0 {
					msg := msgHead.(*Message)
					msg.Data = make([]byte, msg.GetMsgLen())
					_, err := io.ReadFull(conn, msg.Data)
					if err != nil {
						return
					}
					fmt.Println("==> Receive Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
				}

			}
		}(conn)

	}()
	conn, err := net.Dial("tcp", "127.0.0.1:8001")
	if err != nil {
		return
	}
	pkg := NewPackage()

	msg1 := &Message{
		Id:      1,
		Data:    []byte{'a', 'b', 'c', 'd'},
		DataLen: 4,
	}
	sendMsg1, err := pkg.Pack(msg1)

	if err != nil {
		return
	}

	msg2 := &Message{
		Id:      2,
		Data:    []byte{'1', '2', '3', '4'},
		DataLen: 4,
	}
	sendMsg2, err := pkg.Pack(msg2)

	if err != nil {
		return
	}

	sendMsg := append(sendMsg1, sendMsg2...)

	_, err2 := conn.Write(sendMsg)
	if err2 != nil {
		return
	}

	select {}

}
