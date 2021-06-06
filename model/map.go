package model

import (
	"math/rand"

	"time"

	guuid "github.com/google/uuid"
)

// min 4x4
const xsize = 8
const ysize = 8

type Map struct {
	Xsize int      `json:"x_size"`
	Ysize int      `json:"y_size"`
	UUID  string   `json:"uuid"`
	Level int      `json:"level"`
	Grid  [][]Cell `json:"grid"`
}

func NewMap() *Map {

	var grid [][]Cell

	for i := 0; i < xsize; i++ {
		grid = append(grid, []Cell{})
		for j := 0; j < ysize; j++ {
			grid[i] = append(grid[i], Cell{
				Pos: Pos{X: i, Y: j},
			})
		}
	}
	setStartPos(grid)
	setEnemyPos(grid, 15)

	return &Map{
		Xsize: xsize,
		Ysize: ysize,
		UUID:  guuid.New().String(),
		Level: 1,
		Grid:  grid,
	}
}

func setStartPos(grid [][]Cell) {
	rand.Seed(time.Now().UnixNano())
	var edge bool = rand.Intn(2) == 0
	x := 0
	y := rand.Intn(ysize)
	if edge {
		x = rand.Intn(xsize)
		y = 0
	}
	grid[x][y].Start = true
}

func setEnemyPos(grid [][]Cell, count int) {

	for i := 0; i < count; i++ {
		rand.Seed(time.Now().UnixNano() + int64(i))
		pos := free(grid, int64(count))
		grid[pos.X][pos.Y].Enemy = Enemys[rand.Intn(len(Enemys))]
	}
}

func free(grid [][]Cell, seed int64) Pos {
	randx := rand.Intn(xsize)
	randy := rand.Intn(ysize)

	if !grid[randx][randy].IsFree() {
		rand.Seed(seed)
		free(grid, seed+1)
	}

	return Pos{X: randx, Y: randy}
}
