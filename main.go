package main

import (
	"log"

	"github.com/dn46/goRayLib/game"
)

func main() {
	g := game.NewGame()

	err := g.StartWindow()
	if err != nil {
		log.Fatalf("failed to start game window: %v", err)
	}
}
