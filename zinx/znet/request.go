package znet

import "zinx/zinx/ziface"

type Request struct {
	//建立的连接
	conn ziface.IConnection
	//请求的数据
	//data []byte
	message ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}
func (r *Request) GetMsgData() []byte {
	return r.message.GetMsgData()
}
func (r *Request) GetMsgId() uint32 {
	return r.message.GetMsgId()
}

//func (r *Request) GetData() []byte {
//	return r.data
//}
