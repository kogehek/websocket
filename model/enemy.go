package model

import (
	"math/rand"
	"time"
)

type Enemy struct {
	name   string
	atack  int
	health int
}

var Enemys []Enemy = []Enemy{
	{
		name: "skeleton", atack: 1, health: 2,
	},
	{
		name: "mole", atack: 3, health: 1,
	},
	{
		name: "vampire", atack: 2, health: 4,
	},
}

func getEnemy() Enemy {
	rand.Seed(time.Now().Unix())
	return Enemys[rand.Intn(len(Enemys))]
}
