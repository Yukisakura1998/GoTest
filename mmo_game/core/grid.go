package core

import (
	"fmt"
	"sync"
)

type Grid struct {
	//格子ID
	GID int

	MinX int
	MinY int
	MaxX int
	MaxY int
	//格子内对象集合
	playerIDs map[int]bool
	//读写锁
	pIDLock sync.RWMutex
}

// NewGrid 初始化
func NewGrid(gid int, minX, minY, maxX, maxY int) *Grid {
	return &Grid{
		GID:       gid,
		MinX:      minX,
		MinY:      minY,
		MaxX:      maxX,
		MaxY:      maxY,
		playerIDs: make(map[int]bool),
	}
}

func (g *Grid) Add(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	g.playerIDs[playerID] = true
}

func (g *Grid) Remove(playerID int) {
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	delete(g.playerIDs, playerID)
}

func (g *Grid) Contains(playerID int) bool {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	_, ok := g.playerIDs[playerID]
	return ok
}

func (g *Grid) GetPlayerIDs() (playerIDs []int) {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	for playerID := range g.playerIDs {
		playerIDs = append(playerIDs, playerID)
	}
	return playerIDs
}

func (g *Grid) GetPlayerIDByIndex(index int) (playerID int) {
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()
	for playerID := range g.playerIDs {
		if playerID == index {
			return playerID
		}
	}
	return -1
}

func (g *Grid) String() string {
	return fmt.Sprintf("Grid id: %d, minX:%d, maxX:%d, minY:%d, maxY:%d, playerIDs:%v",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIDs)
}
