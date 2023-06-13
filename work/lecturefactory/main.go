package main

import (
	"lecturefactory/factorys"
)

func main() {
	/*
		var enemy factorym.EnemyFactory = factorym.GoblinFactory{}

		goblin := enemy.CreateEnemy()

		goblin.Attack()
	*/

	game := factorys.NewGame(2)
	game.StartGame()
}
