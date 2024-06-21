package main

import (
	"fmt"
	"io"
	"net"
	"zinx/zinx/znet"
)

func main() {
	fmt.Println("Client is starting...")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}

	for {
		//_, err := conn.Write([]byte("go 0.5\r\n"))
		//if err != nil {
		//	return
		//}
		//
		//buff := make([]byte, 512)
		//_, err2 := conn.Read(buff)
		//if err2 != nil {
		//	return
		//}
		//
		//fmt.Printf("%s", buff)
		//time.Sleep(1 * time.Second)
		pkg := znet.NewPackage()
		msg, err := pkg.Pack(znet.NewMessage(0, []byte("Client send \r\n")))
		if err != nil {
			return
		}
		if _, err := conn.Write(msg); err != nil {
			return
		}

		head := make([]byte, pkg.GetHeadLen())
		if _, err := io.ReadFull(conn, head); err != nil {
			return
		}
		msgHead, err := pkg.Unpack(head)
		if err != nil {
			return
		}
		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetMsgLen())
			if _, err := io.ReadFull(conn, msg.Data); err != nil {
				return
			}
			fmt.Println("==> Receive Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
		}
	}

}
