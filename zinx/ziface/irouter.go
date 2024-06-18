package ziface

type IRouter interface {
	// PreHandle 处理之前操作
	PreHandle(request IRequest)
	// MainHandle 处理操作
	MainHandle(request IRequest)
	// PostHandle 处理之后操作
	PostHandle(request IRequest)
}
