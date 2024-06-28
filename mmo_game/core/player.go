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
			continue
		}
	}

}

func (p *Player) GetSurrounding() []*Player {
	surroundPIds := ThisWorldManagement.AoiManagement.GetPlayerIdsByPosition(p.X, p.Z)
	surroundPlayers := make([]*Player, 0, len(surroundPIds))
	for _, pid := range surroundPIds {
		surroundPlayers = append(surroundPlayers, ThisWorldManagement.GetPlayerByPlayerId(int32(pid)))
	}
	return surroundPlayers
}
func (p *Player) SyncSurrounding() {
	//surroundPIds := ThisWorldManagement.AoiManagement.GetPlayerIdsByPosition(p.X, p.Z)
	//surroundPlayers := make([]*Player, 0, len(surroundPIds))
	//for _, pid := range surroundPIds {
	//	surroundPlayers = append(surroundPlayers, ThisWorldManagement.GetPlayerByPlayerId(int32(pid)))
	//}
	surroundPlayers := p.GetSurrounding()
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
	for _, player := range surroundPlayers {
		err := player.SendMsg(200, msg)
		if err != nil {
			continue
		}
	}

	playersData := make([]*pb.Player, 0, len(surroundPlayers))
	for _, player := range surroundPlayers {
		//有必要吗？有，要从core.player转为pb.player
		p := &pb.Player{
			Pid: player.PlayerId,
			P: &pb.Position{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
				V: player.V,
			},
		}
		//
		playersData = append(playersData, p)
	}
	SyncPlayersMsg := &pb.SyncPlayers{
		Ps: playersData[:],
	}
	err := p.SendMsg(202, SyncPlayersMsg)
	if err != nil {
		return
	}
}

func (p *Player) Move(x float32, y float32, z float32, v float32) {
	p.X = x
	p.Y = y
	p.Z = z
	p.V = v
	//send to surrounding
	msg := &pb.BroadCast{
		Pid: p.PlayerId,
		Tp:  4,
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	players := p.GetSurrounding()
	for _, player := range players {
		err := player.SendMsg(200, msg)
		if err != nil {
			continue
		}
	}
}

func (p *Player) Logout() {
	players := p.GetSurrounding()
	msg := &pb.SyncPlayerId{
		PlayerId: p.PlayerId,
	}
	for _, player := range players {
		err := player.SendMsg(201, msg)
		if err != nil {
			continue
		}
	}
	//remove to management
	ThisWorldManagement.AoiManagement.RemovePlayerIdFromGridByPosition(p.X, p.Z, int(p.PlayerId))
	ThisWorldManagement.RemovePlayer(p.PlayerId)
}
