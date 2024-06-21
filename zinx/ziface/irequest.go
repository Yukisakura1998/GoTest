package ziface

type IRequest interface {
	//获取当前连接
	GetConnection() IConnection
	//获取请求数据
	GetMsgData() []byte
	//
	GetMsgId() uint32
}
