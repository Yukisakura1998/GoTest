package api

import (
	"google.golang.org/protobuf/proto"
	"zinx/mmo_game/core"
	"zinx/mmo_game/pb"
	"zinx/zinx/ziface"
	"zinx/zinx/znet"
)

type WorldChatAPI struct {
	znet.BaseRouter
}

func (API *WorldChatAPI) Handle(request ziface.IRequest) {
	msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetMsgData(), msg)
	if err != nil {
		return
	}

	pid, err := request.GetConnection().GetProp("pid")
	if err != nil {
		return
	}
	player := core.ThisWorldManagement.GetPlayerByPlayerId(pid.(int32))

	player.Talk(msg.Content)

}
