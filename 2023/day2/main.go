package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")
	fmt.Println(part1(input))
}

func part1(input []string) string {
	runningTotal := 0

	regex := regexp.MustCompile(`\d+`)

	for gameId, game := range input {
		if game == "" {
			continue
		}

		possible := true

		gameData := strings.Split(game, ":")[1]
		handfuls := strings.Split(gameData, ";")

		for _, stringHandful := range handfuls {
			cubes := strings.Split(stringHandful, ",")
			handful := cubeHandful{}

			for _, cube := range cubes {
				num := helpers.MustAtoi(regex.FindStringSubmatch(cube)[0])

				if strings.Contains(cube, "red") {
					handful.red = num
					continue
				}
				if strings.Contains(cube, "green") {
					handful.green = num
					continue
				}
				if strings.Contains(cube, "blue") {
					handful.blue = num
					continue
				}
			}
			if !handful.isPossible() {
				possible = false
				break
			}
		}
		if possible {
			runningTotal += gameId + 1
		}
	}
	return fmt.Sprintf("%d", runningTotal)
}

type cubeHandful struct {
	red   int
	green int
	blue  int
}

func (c cubeHandful) isPossible() bool {
	return c.red <= 12 && c.green <= 13 && c.blue <= 14
}
