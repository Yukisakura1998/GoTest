package ziface

type IServer interface {
	Start()
	Stop()
	Server()
	// AddRouter 注册路由方法
	AddRouter(router IRouter)
}
