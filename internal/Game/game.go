package game

import world "proc-gen-world/internal/World"

type Game struct {
	world.World
}

func NewGame() (Game, error) {
	return Game{}, nil
}
