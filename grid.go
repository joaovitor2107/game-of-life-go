package main

import "sync"

type CellState int
const(
	Dead CellState = 0
	Alive CellState = 1
)

type Grid struct {
	size int
	cells [][] CellState
	nextgen [][] CellState
}

func newGrid (size int) *Grid {
	cells := make([][]CellState, size)
	nextgen := make([][]CellState, size)
	for i := range cells {
		cells[i] = make([]CellState, size)
		nextgen[i] = make([]CellState, size)
	}
	return &Grid{
		size: size,
		cells: cells,
		nextgen: nextgen,
	}
}

func (g *Grid) NextGeneration() {
	var wg sync.WaitGroup
	for i := range g.cells {
		for j:= range g.cells[i] {
			wg.Add(1)
			go func (i, j int){
				g.UpdateCell(i, j)
				defer wg.Done()
			} (i, j)
		}
	}

	wg.Wait()
	g.cells, g.nextgen = g.nextgen, g.cells
}

// rules:
// if a cell is alive it survives if it has 2 or 3 live neighbors
// if a cell is dead it comes to life if it has 3 live neighbors
//
// edge cases:
// wraps around edges so the grid behaves well
func (g *Grid) UpdateCell(i, j int) {
	neighbors := g.AliveCells(i, j)

	// the rules collapse to the cells is going to be alive if it has
	// 3 living neighbors or if it is alive and have 2 living neighbors
	if (g.cells[i][j] == Alive && neighbors == 2) || neighbors == 3 {
		g.nextgen[i][j] = Alive
		return
	}
	g.nextgen[i][j] = Dead
}

func (g *Grid) AliveCells (i, j int) int {
	cellsalive := 0

	for a := i-1; a <= i+1; a++ {
		for b := j-1; b <= j+1; b++ {
			if a == i && b == j {
				continue
			}
			na := (g.size + a) % g.size
			nb := (g.size + b) % g.size
			if g.cells[na][nb] == Alive {
				cellsalive++
			}
		}
	}

	return cellsalive
}
