package main

import (
	"fmt"
	"io"
	"net"
	"time"
	"zinx/zinx/znet"
)

func main() {
	fmt.Println("Client is starting...")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Client error...")
		fmt.Println(err)
		return
	}
	for {
		time.Sleep(10 * time.Second)
		pkg := znet.NewPackage()
		msg, err := pkg.Pack(znet.NewMessage(0, []byte("Client send \r\n")))
		if err != nil {
			fmt.Println("1:", err)
			return
		}
		if _, err := conn.Write(msg); err != nil {
			fmt.Println("2:", err)
			return
		}
		head := make([]byte, pkg.GetHeadLen())
		if _, err := io.ReadFull(conn, head); err != nil {
			fmt.Println("3:", err)
			return
		}
		msgHead, err := pkg.Unpack(head)
		if err != nil {
			fmt.Println("4:", err)
			return
		}
		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				fmt.Println("5:", err)
				return
			}
			fmt.Println("==> Receive Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
		}
	}

}
