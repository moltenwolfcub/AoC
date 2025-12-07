package main

import (
	"fmt"

	"github.com/moltenwolfcub/AoC/helpers"
)

type Vec2D struct {
	X, Y int
}

var grid [][]bool

func main() {
	input := helpers.ReadLines("input.txt")

	grid = make([][]bool, 0)
	rolls := make([]Vec2D, 0)

	for y, line := range input {
		if line == "" {
			continue
		}
		gridRow := make([]bool, len(line))

		for x, pos := range line {
			if pos == '@' {
				gridRow[x] = true
				rolls = append(rolls, Vec2D{x, y})
			}
		}
		grid = append(grid, gridRow)
	}

	accessibles := 0
	for _, roll := range rolls {
		adjacents := checkAdjacent(roll)

		if adjacents < 4 {
			accessibles++
		}
	}
	fmt.Println(accessibles)
}

func checkAdjacent(pos Vec2D) int {
	count := 0
	for x := pos.X - 1; x <= pos.X+1; x++ {
		if x < 0 || x >= len(grid[0]) {
			continue
		}
		for y := pos.Y - 1; y <= pos.Y+1; y++ {
			if y < 0 || y >= len(grid) {
				continue
			}
			if x == pos.X && y == pos.Y {
				continue
			}

			if grid[y][x] {
				count++
			}
		}
	}
	return count
}
