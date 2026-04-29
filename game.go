package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
    grid      *Grid
    cellSize  int
	tick      int
}

func (g *Game) Update() error {
	g.tick++
	if g.tick % 10 == 0{
		g.grid.NextGeneration()
	}
    return nil
}

// black backgroud, if the cell is alive it turns red
func (g *Game) Draw(screen *ebiten.Image) {
    for i := 0; i < g.grid.size; i++ {
        for j := 0; j < g.grid.size; j++ {
            if g.grid.cells[i][j] == Alive {
                for dx := 0; dx < g.cellSize; dx++ {
                    for dy := 0; dy < g.cellSize; dy++ {
                        screen.Set(
                            i*g.cellSize+dx,
                            j*g.cellSize+dy,
							color.RGBA{255, 0, 0, 255},
                        )
                    }
                }
            }
        }
    }
}

func (g *Game) Layout(w, h int) (int, int) {
    return g.grid.size * g.cellSize, g.grid.size * g.cellSize
}
