package core

import "sync"

type WorldManagement struct {
	AoiManagement *AOIManagement
	Players       map[int32]*Player
	pLock         sync.RWMutex
}

var ThisWorldManagement *WorldManagement

func init() {
	ThisWorldManagement = &WorldManagement{
		AoiManagement: NewAOIManagement(0, 0, 250, 250, 5, 5),
		Players:       make(map[int32]*Player),
	}
}
func (m *WorldManagement) AddPlayer(player *Player) {
	m.pLock.Lock()
	defer m.pLock.Unlock()
	m.Players[player.PlayerId] = player
	m.AoiManagement.AddPlayerIdToGridByPosition(player.X, player.Z, int(player.PlayerId))
}
func (m *WorldManagement) RemovePlayer(PlayerId int32) {
	player := m.Players[PlayerId]
	m.pLock.Lock()
	defer m.pLock.Unlock()
	delete(m.Players, PlayerId)
	m.AoiManagement.RemovePlayerIdFromGridByPosition(player.X, player.Z, int(PlayerId))
}
func (m *WorldManagement) GetPlayerByPlayerId(playerId int32) *Player {
	m.pLock.RLock()
	defer m.pLock.RUnlock()
	return m.Players[playerId]
}
func (m *WorldManagement) GetAllPlayers() []*Player {
	m.pLock.RLock()
	defer m.pLock.RUnlock()
	players := make([]*Player, len(m.Players))
	for _, player := range m.Players {
		players = append(players, player)
	}
	return players
}
