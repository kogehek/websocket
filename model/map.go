package model

import (
	"math/rand"

	guuid "github.com/google/uuid"
)

// min 4x4
const xsize = 8
const ysize = 8

type Map struct {
	Xsize int      `json:"xsize"`
	Ysize int      `json:"ysize"`
	UUID  string   `json:"uuid"`
	Level int      `json:"level"`
	Grid  [][]Cell `json:"grid"`
}

func NewMap(b Biome) *Map {
	var grid [][]Cell

	for i := 0; i < xsize; i++ {
		for j := 0; j < ysize; j++ {
			grid[i][j] = Cell{
				Pos: Pos{X: i, Y: j},
			}
			// grid = append(grid, Cell{
			// 	Pos: Pos{X: i, Y: j},
			// })
		}
	}

	return &Map{
		Xsize: xsize,
		Ysize: ysize,
		UUID:  guuid.New().String(),
		Level: 1,
		Grid:  grid,
	}
}

func setStartPos(grid [][]Cell) {
	x := rand.Intn(xsize-0) + 0
	y := 0
	if x != 0 {
		y = rand.Intn(ysize-0) + 0
	}
	grid[x][y].Start = Pos{X: x, Y: y}
	// return Pos{X: x, Y: y}
}
