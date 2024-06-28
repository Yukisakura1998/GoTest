package api

import (
	"google.golang.org/protobuf/proto"
	"zinx/mmo_game/core"
	"zinx/mmo_game/pb"
	"zinx/zinx/ziface"
	"zinx/zinx/znet"
)

type MoveAPI struct {
	znet.BaseRouter
}

func (API *MoveAPI) Handle(request ziface.IRequest) {
	msg := &pb.Position{}
	err := proto.Unmarshal(request.GetMsgData(), msg)
	if err != nil {
		return
	}

	pid, err := request.GetConnection().GetProp("pid")
	if err != nil {
		request.GetConnection().Stop()
		return
	}
	player := core.ThisWorldManagement.GetPlayerByPlayerId(pid.(int32))

	player.Move(msg.X, msg.Y, msg.Z, msg.V)

}
