package core

import (
	"fmt"
	"testing"
)

func TestNewAOIManagement(t *testing.T) {
	aoiManagement := NewAOIManagement(0, 250, 0, 250, 5, 5)
	fmt.Println(aoiManagement)
}

func TestGetSurroundGridByGridID(t *testing.T) {
	aoiManagement := NewAOIManagement(0, 250, 0, 250, 5, 5)
	for grid, _ := range aoiManagement.grids {
		surroundGrids := aoiManagement.GetSurroundGridByGridID(grid)
		fmt.Println("grid : ", grid, " grids len : ", len(surroundGrids))
		gids := make([]int, 0, len(surroundGrids))
		for _, surroundGrid := range surroundGrids {
			gids = append(gids, surroundGrid.GID)
		}
		fmt.Println(gids)
	}
}

// protoc --go_out=. --go_opt=paths=source_relative
