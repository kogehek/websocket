package model

import (
	"math/rand"
	"time"
)

type Enemy struct {
	Name   string `json:"name"`
	Atack  int    `json:"atack"`
	Health int    `json:"health"`
}

var Enemys []Enemy = []Enemy{
	{
		Name: "skeleton", Atack: 1, Health: 2,
	},
	{
		Name: "mole", Atack: 3, Health: 1,
	},
	{
		Name: "vampire", Atack: 2, Health: 4,
	},
}

func getEnemy() Enemy {
	rand.Seed(time.Now().Unix())
	return Enemys[rand.Intn(len(Enemys))]
}
