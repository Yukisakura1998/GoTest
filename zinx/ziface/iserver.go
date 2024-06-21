package ziface

type IServer interface {
	Start()
	Stop()
	Server()
	// AddRouter 注册路由方法
	AddRouter(msgId uint32, router IRouter)
	GetConnManager() IConnectManager
	// SetOnConnStart 设置该Server的连接创建时Hook函数
	SetOnConnStart(func(IConnection))
	// SetOnConnStop 设置该Server的连接断开时的Hook函数
	SetOnConnStop(func(IConnection))
	// CallOnConnStart 调用连接OnConnStart Hook函数
	CallOnConnStart(conn IConnection)
	// CallOnConnStop 调用连接OnConnStop Hook函数
	CallOnConnStop(conn IConnection)
}
