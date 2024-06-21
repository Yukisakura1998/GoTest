package core

import "fmt"

// AOIManagement AOI管理模块
type AOIManagement struct {
	MinX   int
	MinY   int
	MaxX   int
	MaxY   int
	CountX int
	CountY int
	grids  map[int]*Grid
}

func NewAOIManagement(minX, minY, maxX, maxY, countX, countY int) *AOIManagement {
	aoi := &AOIManagement{
		MinX:   minX,
		MinY:   minY,
		MaxX:   maxX,
		MaxY:   maxY,
		CountX: countX,
		CountY: countY,
		grids:  make(map[int]*Grid),
	}
	for y := 0; y < countY; y++ {
		for x := 0; x < countX; x++ {
			gid := y*countX + x

			aoi.grids[gid] = NewGrid(gid,
				aoi.MinX+x*aoi.gridWidth(),
				aoi.MinX+(x+1)*aoi.gridWidth(),
				aoi.MinY+y*aoi.gridHeight(),
				aoi.MaxY+(y+1)*aoi.gridHeight())
		}
	}
	return aoi
}
func (m *AOIManagement) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CountX
}
func (m *AOIManagement) gridHeight() int {
	return (m.MaxY - m.MinY) / m.CountY
}
func (m *AOIManagement) String() string {
	s := fmt.Sprintf("AOIManagr:\r\nminX:%d, maxX:%d, cntsX:%d, minY:%d, maxY:%d, cntsY:%d\r\n Grids in AOI Manager:\r\n",
		m.MinX, m.MaxX, m.CountX, m.MinY, m.MaxY, m.CountY)
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}

	return s
}
func (m *AOIManagement) GetSurroundGridByGridID(gridID int) (grids []*Grid) {
	if _, ok := m.grids[gridID]; ok {

		grids = append(grids, m.grids[gridID])

		idx := gridID % m.CountX

		if idx > 0 {
			grids = append(grids, m.grids[gridID-1])
		}
		if idx < m.CountX-1 {
			grids = append(grids, m.grids[gridID+1])
		}

		gidsX := make([]int, 0, len(grids))

		for _, gridX := range grids {
			gidsX = append(gidsX, gridX.GID)
		}

		for _, gridX := range gidsX {
			idy := gridX / m.CountX
			if idy > 0 {
				grids = append(grids, m.grids[gridX-m.CountX])
			}
			if idy < m.CountY-1 {
				grids = append(grids, m.grids[gridX+m.CountX])
			}
		}
		return
	}
	return
}

func (m *AOIManagement) GetPlayerIdsByPosition(x, y float32) (playerIDs []int) {
	gridID := m.GetGridByPosition(x, y)
	grids := m.GetSurroundGridByGridID(gridID)
	for _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetPlayerIDs()...)
		fmt.Printf("<---gridID : %d, pids : %v--->\r\n", grid.GID, grid.GetPlayerIDs())
	}
	return
}
func (m *AOIManagement) GetGridByPosition(x, y float32) (grid int) {
	idx := (int(x) - m.MinX) / m.gridWidth()
	idy := (int(y) - m.MinY) / m.gridHeight()
	return idy*m.CountX + idx
}
func (m *AOIManagement) AddPlayerIdToGrid(pID, gID int) {
	m.grids[gID].Add(pID)
}

func (m *AOIManagement) RemovePidFromGrid(pID, gID int) {
	m.grids[gID].Remove(pID)
}

func (m *AOIManagement) GetPlayerIdsByGrid(gID int) (playerIDs []int) {
	playerIDs = m.grids[gID].GetPlayerIDs()
	return
}

func (m *AOIManagement) AddPlayerIdToGridByPosition(x, y float32, playerID int) {
	gID := m.GetGridByPosition(x, y)
	grid := m.grids[gID]
	grid.Add(playerID)
}

func (m *AOIManagement) RemovePlayerIdFromGridByPosition(x, y float32, playerID int) {
	gID := m.GetGridByPosition(x, y)
	grid := m.grids[gID]
	grid.Remove(playerID)
}
