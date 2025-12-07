package main

import (
	"fmt"

	"github.com/moltenwolfcub/AoC/helpers"
)

type Vec2D struct {
	X, Y int
}

func main() {
	input := helpers.ReadLines("input.txt")

	grid := make([][]bool, 0)
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

	fmt.Printf("Part 1: %v\n", part1(rolls, grid))
	fmt.Printf("Part 2: %v\n", part2(rolls, grid))

}

func part1(initialRolls []Vec2D, initialGrid [][]bool) int {
	accessibles := 0
	for _, roll := range initialRolls {
		adjacents := checkAdjacent(roll, initialGrid)

		if adjacents < 4 {
			accessibles++
		}
	}

	return accessibles
}

func part2(initialRolls []Vec2D, initialGrid [][]bool) int {
	accessibles := 0

	currentRolls := append(make([]Vec2D, 0), initialRolls...)
	nextRolls := make([]Vec2D, 0)
	wasChange := true

	for wasChange {
		wasChange = false
		for _, roll := range currentRolls {
			adjacents := checkAdjacent(roll, initialGrid)

			if adjacents < 4 {
				accessibles++
				initialGrid[roll.Y][roll.X] = false
				wasChange = true
			} else {
				nextRolls = append(nextRolls, roll)
			}
		}
		currentRolls = nextRolls
		nextRolls = make([]Vec2D, 0)
	}

	return accessibles
}

func checkAdjacent(pos Vec2D, grid [][]bool) int {
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
