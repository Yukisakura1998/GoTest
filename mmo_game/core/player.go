package core

import (
	"errors"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"zinx/mmo_game/pb"
	"zinx/zinx/ziface"
)

type Player struct {
	PlayerId int32
	Conn     ziface.IConnection
	X        float32
	Y        float32
	Z        float32
	V        float32
}

func NewPlayer(conn ziface.IConnection) *Player {
	//用户验证可能需要放在这里，比如用户名和密码的验证，初始化玩家位置应该从数据库取出相应的坐标以及各种状态等。
	p := &Player{
		Conn:     conn,
		PlayerId: 0,
		X:        float32(160 + rand.Intn(10)), //随机在160坐标点 基于X轴偏移若干坐标
		Y:        0,
		Z:        float32(134 + rand.Intn(10)),
		V:        0,
	}
	return p
}
func (p *Player) SendMsg(msgId uint32, data proto.Message) error {
	msg, err := proto.Marshal(data)
	if err != nil {
		return err
	}
	if p.Conn == nil {
		return errors.New("NO CONNECTION FOUND")
	}

	err = p.Conn.Send(msgId, msg)
	if err != nil {
		return err
	}

	return nil
}
func (p *Player) SyncPlayerId() {
	//Id = 1
	data := &pb.SyncPlayerId{
		PlayerId: p.PlayerId,
	}
	err := p.SendMsg(1, data)

	if err != nil {
		return
	}
}
func (p *Player) BroadCastStartPosition() {
	msg := &pb.BroadCast{
		Pid: p.PlayerId,
		Tp:  2,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	err := p.SendMsg(200, msg)

	if err != nil {
		return
	}
}
func (p *Player) Talk(content string) {
	msg := &pb.BroadCast{
		Pid: p.PlayerId,
		Tp:  1,
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}
	players := ThisWorldManagement.GetAllPlayers()
	for _, player := range players {
		err := player.SendMsg(200, msg)
		if err != nil {
			return
		}
	}

}
