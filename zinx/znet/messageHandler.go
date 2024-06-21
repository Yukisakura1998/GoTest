package znet

import (
	"fmt"
	"zinx/zinx/utils"
	"zinx/zinx/ziface"
)

type MessageHandler struct {
	Apis map[uint32]ziface.IRouter
	//消息请求列表
	TaskQueue []chan ziface.IRequest
	//工作数量
	WorkerPoolSize uint32
}

func (msgHandler *MessageHandler) DoMessageHandler(request ziface.IRequest) {
	handler, ok := msgHandler.Apis[request.GetMsgId()]
	if !ok {
		fmt.Println("api msg id '", request.GetMsgId(), " is not register")
		return
	}
	handler.PreHandle(request)
	handler.Handle(request)
	handler.PostHandle(request)
}
func (msgHandler *MessageHandler) AddRouter(msgId uint32, router ziface.IRouter) {
	if _, ok := msgHandler.Apis[msgId]; ok {
		fmt.Println("id exist:", msgId)
	}
	msgHandler.Apis[msgId] = router
	fmt.Println("add router id:", msgId)
}
func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		Apis:           make(map[uint32]ziface.IRouter),
		TaskQueue:      make([]chan ziface.IRequest, utils.GlobalObject.WorkerPoolSize),
		WorkerPoolSize: utils.GlobalObject.WorkerPoolSize,
	}
}

func (msgHandler *MessageHandler) StartWorkerPool() {
	for i := 0; i < int(utils.GlobalObject.WorkerPoolSize); i++ {
		msgHandler.TaskQueue[i] = make(chan ziface.IRequest, utils.GlobalObject.MaxWorkerPoolSize)
		go msgHandler.StartWorker(i, msgHandler.TaskQueue)
	}
}

func (msgHandler *MessageHandler) StartWorker(id int, taskQueue []chan ziface.IRequest) {
	fmt.Println("work id is:", id, "is start")

	for {
		select {
		case request := <-taskQueue[id]:
			msgHandler.DoMessageHandler(request)
		}
	}
}

func (msgHandler *MessageHandler) SendMsgToTaskQueue(request ziface.IRequest) {
	//分配方式
	workID := request.GetConnection().GetConnectID() % msgHandler.WorkerPoolSize
	fmt.Println("Add ConnID=", request.GetConnection().GetConnectID(), " request msgID=", request.GetMsgId(), "to workerID=", workID)
	//Send
	msgHandler.TaskQueue[workID] <- request
}
