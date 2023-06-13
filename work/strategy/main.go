package main

import (
	"strategy/game"
)

func main() { // Caller
	queen := game.Character{}
	queen.SetBehavior(game.BowAndArrowBehavior{})
	queen.Fight()
}
