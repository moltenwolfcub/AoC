package main

import (
	"fmt"
	"image"
	"slices"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	galaxies := []image.Point{}

	for y, line := range input {
		if line == " " {
			continue
		}
		for x, c := range line {
			if c == '#' {
				galaxies = append(galaxies, image.Pt(x, y))
			}
		}
	}

	maxX := len(input[0])
	maxY := len(input) - 1 // -1 cause empty line on the end

	emptyRows := []int{}
	emptyCols := []int{}

	for i := 0; i < maxY; i++ {
		hasGalaxy := false
		for _, g := range galaxies {
			if g.Y == i {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			emptyRows = append(emptyRows, i)
		}
	}
	for i := 0; i < maxX; i++ {
		hasGalaxy := false
		for _, g := range galaxies {
			if g.X == i {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			emptyCols = append(emptyCols, i)
		}
	}
	slices.Reverse(emptyRows)
	slices.Reverse(emptyCols)

	for _, row := range emptyRows {
		for i, g := range galaxies {
			if row < g.Y {
				galaxies[i] = g.Add(image.Pt(0, 1))
			}
		}
	}
	for _, col := range emptyCols {
		for i, g := range galaxies {
			if col < g.X {
				galaxies[i] = g.Add(image.Pt(1, 0))
			}
		}
	}

	total := 0

	toCheck := galaxies

	for _, this := range toCheck {
		for _, other := range toCheck[1:] {
			dx := helpers.IntAbs(this.X - other.X)
			dy := helpers.IntAbs(this.Y - other.Y)
			dist := dx + dy
			total += dist
		}
		toCheck = toCheck[1:]
	}

	fmt.Printf("Part 1: %d\n", total)
}
