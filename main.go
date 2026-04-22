package main

import (
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func randomElements (g *Grid) {
	for i := range g.cells {
		for j := range g.cells[i] {
			g.cells[i][j] = CellState(rand.Intn(2))
		}
	}
}

func main() {
	g := newGrid(100)
	randomElements(g)

	game := &Game {
		grid:  g,
		cellSize: 8,
	}

	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Cellular Automata")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
