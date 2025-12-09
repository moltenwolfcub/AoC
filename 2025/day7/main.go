package main

import (
	"fmt"
	"slices"

	"github.com/moltenwolfcub/AoC/helpers"
)

type Beam struct {
	Pos int
}

func main() {
	input := helpers.ReadLines("input.txt")

	grid := make([][]bool, len(input)-1)

	var start Beam
	for y, line := range input {
		if line == "" {
			continue
		}

		row := make([]bool, len(line))
		for x, c := range line {
			switch c {
			case '^':
				row[x] = true
			case 'S':
				start = Beam{x}
			}
		}
		grid[y] = row
	}

	currentBeams := []Beam{start}
	nextBeams := make([]Beam, 0)

	splits := 0
	for _, row := range grid {
		for _, beam := range currentBeams {
			if row[beam.Pos] {
				splits++
				left := Beam{beam.Pos - 1}
				right := Beam{beam.Pos + 1}

				if !slices.Contains(nextBeams, left) {
					nextBeams = append(nextBeams, left)
				}

				if !slices.Contains(nextBeams, right) {
					nextBeams = append(nextBeams, right)
				}

			} else {
				if !slices.Contains(nextBeams, beam) {
					nextBeams = append(nextBeams, beam)
				}
			}
		}
		currentBeams = make([]Beam, len(nextBeams))
		currentBeams = append(currentBeams, nextBeams...)
		nextBeams = make([]Beam, 0)
	}

	fmt.Println(splits)
}
