package ziface

type IMassageHandle interface {
	DoMessageHandler(request IRequest)
	AddRouter(msgId uint32, router IRouter)
	StartWorkerPool()
	SendMsgToTaskQueue(request IRequest)
}
