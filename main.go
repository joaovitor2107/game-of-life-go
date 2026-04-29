package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := NewGrid(40)
	//g.RandomElements()

	println(BuildOffset(Glider))

	game := &Game {
		grid:  g,
		cellSize: 8,
	}

	ebiten.SetWindowSize(840, 840)
	ebiten.SetWindowTitle("Cellular Automata")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
