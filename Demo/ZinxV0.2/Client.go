package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("Client is starting...")

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}

	for {
		_, err := conn.Write([]byte("go 0.2/r/n"))
		if err != nil {
			return
		}

		buff := make([]byte, 512)
		cnt, err2 := conn.Read(buff)
		if err2 != nil {
			return
		}

		fmt.Printf("%s,%d", buff, cnt)
		time.Sleep(1 * time.Second)
	}

}
