package model

import (
	guuid "github.com/google/uuid"
)

type Game struct {
	UUID string `json:"uuid"`
	Map  *Map   `json:"map"`
}

func NewMGame() *Game {
	return &Game{
		UUID: guuid.New().String(),
		Map:  NewMap(),
	}
}
