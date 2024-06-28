package main

import (
	"zinx/mmo_game/api"
	"zinx/mmo_game/core"
	"zinx/zinx/ziface"
	"zinx/zinx/znet"
)

func main() {
	s := znet.NewServer()
	s.SetOnConnStart(OnConnectionAdd)
	s.AddRouter(2, &api.WorldChatAPI{})
	s.AddRouter(3, &api.MoveAPI{})
	s.SetOnConnStop(OnConnectionLost)
	s.Server()
}

func OnConnectionAdd(conn ziface.IConnection) {
	player := core.NewPlayer(conn)
	//ID = 1
	player.SyncPlayerId()
	//ID = 200
	player.BroadCastStartPosition()
	//add to management
	core.ThisWorldManagement.AddPlayer(player)
	//key playerId
	conn.SetProp("pid", player.PlayerId)
	//同步周围环境
	player.SyncSurrounding()
}

func OnConnectionLost(conn ziface.IConnection) {
	playerID, _ := conn.GetProp("pid")
	player := core.ThisWorldManagement.GetPlayerByPlayerId(playerID.(int32))
	//触发玩家下线业务
	if playerID != nil {
		player.Logout()
	}
	//key playerId
	conn.RemoveProp("pid")
}
