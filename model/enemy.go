package model

import (
	"math/rand"
	"time"
)

type Enemy struct {
	Name   string `json:"name"`
	Attack int    `json:"attack"`
	Health int    `json:"health"`
}

var Enemys []Enemy = []Enemy{
	{
		Name: "skeleton", Attack: 1, Health: 2,
	},
	{
		Name: "mole", Attack: 3, Health: 1,
	},
	{
		Name: "vampire", Attack: 2, Health: 4,
	},
}

func getEnemy() Enemy {
	rand.Seed(time.Now().Unix())
	return Enemys[rand.Intn(len(Enemys))]
}
