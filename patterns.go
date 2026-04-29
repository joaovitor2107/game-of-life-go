package main

type Pattern [][2] int

// implementation of different patterns for initial states
var(
	// {x, y}
	Glider = Pattern{
		{0, 1},
		{1, 2},
		{2, 0},
		{2, 1},
		{2, 2},
	}

	Block = Pattern{
		{0,0},
		{0,1},
		{1,0},
		{1,1},
	}

	Boat = Pattern{
		{0,0},
		{0,1},
		{1,0},
		{1,2},
		{2,1},
	}

	Tub = Pattern{
		{0,1},
		{1,0},
		{1,2},
		{2,1},
	}

	Toad = Pattern{
		{0,1},
		{0,2},
		{1,3},
		{2,0},
		{3,1},
		{3,2},
	}

	Blinker = Pattern{
		{0,0},
		{0,1},
		{0,2},
	}
)

var Patterns = []Pattern{Glider, Block, Boat, Tub, Toad, Blinker}

func BuildOffset (p Pattern) (int, int) {
	Xmax, Ymax := 0, 0

	for _, pair := range p {
		Xmax = max(Xmax, pair[0])
		Ymax = max(Ymax, pair[1])
	}

	return Xmax + 1, Ymax + 1
}
