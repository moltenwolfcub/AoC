package main

import (
	"fmt"
	"slices"

	"github.com/moltenwolfcub/AoC/helpers"
)

type Beam struct {
	Pos    int
	Weight int
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
				start = Beam{x, 1}
			}
		}
		grid[y] = row
	}

	fmt.Printf("Part 1: %v\n", part1(grid, start))
	fmt.Printf("Part 2: %v\n", part2(grid, start))
}

func part1(grid [][]bool, start Beam) int {
	currentBeams := []Beam{start}
	nextBeams := make([]Beam, 0)

	splits := 0
	for _, row := range grid {
		for _, beam := range currentBeams {
			if row[beam.Pos] {
				splits++
				left := Beam{beam.Pos - 1, 1}
				right := Beam{beam.Pos + 1, 1}

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
		currentBeams = make([]Beam, 0)
		currentBeams = append(currentBeams, nextBeams...)
		nextBeams = make([]Beam, 0)
	}

	return splits
}

func part2(grid [][]bool, start Beam) int {
	currentBeams := []Beam{start}
	nextBeams := make([]Beam, 0)

	for _, row := range grid {
		for _, beam := range currentBeams {
			if row[beam.Pos] {
				left := Beam{beam.Pos - 1, beam.Weight}
				right := Beam{beam.Pos + 1, beam.Weight}

				l := slices.Index(nextBeams, left)
				if l < 0 {
					nextBeams = append(nextBeams, left)
				} else {
					nextBeams[l].Weight += left.Weight
				}

				r := slices.Index(nextBeams, right)
				if r < 0 {
					nextBeams = append(nextBeams, right)
				} else {
					nextBeams[r].Weight += right.Weight
				}
			} else {
				b := slices.Index(nextBeams, beam)
				if b < 0 {
					nextBeams = append(nextBeams, beam)
				} else {
					nextBeams[b].Weight += beam.Weight
				}
			}
		}
		currentBeams = make([]Beam, 0)
		currentBeams = append(currentBeams, nextBeams...)
		nextBeams = make([]Beam, 0)
	}

	universes := 0
	for _, b := range currentBeams {
		universes += b.Weight
	}

	return universes
}
