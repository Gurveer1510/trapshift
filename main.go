package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := newGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("TrapShift")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
