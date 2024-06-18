package znet

import "zinx/zinx/ziface"

type Request struct {
	//建立的连接
	conn ziface.IConnection
	//请求的数据
	data []byte
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}
func (r *Request) GetData() []byte {
	return r.data
}
