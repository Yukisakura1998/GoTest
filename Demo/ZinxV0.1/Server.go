package main

import "zinx/zinx/znet"

func main() {
	s := znet.NewServer("[ZinxV0.1]")
	s.Server()
}
