package main

import (
	"fmt"
	"image"

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

	calcDist := func(distBtwn int) int {
		scaledGalaxies := make([]image.Point, len(galaxies))
		copy(scaledGalaxies, galaxies)

		for _, row := range emptyRows {
			for i, g := range galaxies {
				if row < g.Y {
					scaledGalaxies[i] = scaledGalaxies[i].Add(image.Pt(0, distBtwn))
				}
			}
		}
		for _, col := range emptyCols {
			for i, g := range galaxies {
				if col < g.X {
					scaledGalaxies[i] = scaledGalaxies[i].Add(image.Pt(distBtwn, 0))
				}
			}
		}

		total := 0

		toCheck := make([]image.Point, len(scaledGalaxies))
		copy(toCheck, scaledGalaxies)

		for _, this := range toCheck {
			for _, other := range toCheck[1:] {
				dx := helpers.IntAbs(this.X - other.X)
				dy := helpers.IntAbs(this.Y - other.Y)
				dist := dx + dy
				total += dist
			}
			toCheck = toCheck[1:]
		}
		return total
	}

	fmt.Printf("Part 1: %d\n", calcDist(1))
	fmt.Printf("Part 2: %d\n", calcDist(999999))
}
