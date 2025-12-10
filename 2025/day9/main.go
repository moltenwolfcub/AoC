package main

import (
	"fmt"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

type Tile struct {
	X, Y int
}

func (t Tile) Size() int {
	return t.X * t.Y
}

func main() {
	input := helpers.ReadLines("input.txt")

	tiles := make([]Tile, len(input)-1)
	for i, line := range input {
		if line == "" {
			continue
		}

		coords := strings.Split(line, ",")
		tiles[i] = Tile{helpers.MustAtoi(coords[0]), helpers.MustAtoi(coords[1])}
	}

	greatestArea := 0
	for f := 0; f < len(tiles); f++ {
		first := tiles[f]
		for s := f + 1; s < len(tiles); s++ {
			second := tiles[s]

			a := area(first, second)
			greatestArea = max(greatestArea, a)
		}
	}
	fmt.Println(greatestArea)
}

func area(first, second Tile) int {
	w := helpers.IntAbs(second.X-first.X) + 1
	h := helpers.IntAbs(second.Y-first.Y) + 1
	return w * h
}
