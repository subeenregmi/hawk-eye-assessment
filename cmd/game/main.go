package main

import (
	"github.com/subeenregmi/hawk-eye-assessment/internal/game"
)

func main() {
	g := game.Game{WinThreshold: 8}
	if err := g.Start(); err != nil {
		panic(err)
	}
}
