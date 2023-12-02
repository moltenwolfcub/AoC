package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	runningTotal := 0
	runingPower := 0

	regex := regexp.MustCompile(`\d+`)

	for gameId, game := range input {
		if game == "" {
			continue
		}

		possible := true
		miniumNeeded := cubeHandful{}

		gameData := strings.Split(game, ":")[1]
		handfuls := strings.Split(gameData, ";")

		for _, stringHandful := range handfuls {
			cubes := strings.Split(stringHandful, ",")
			handful := cubeHandful{}

			for _, cube := range cubes {
				num := helpers.MustAtoi(regex.FindStringSubmatch(cube)[0])

				if strings.Contains(cube, "red") {
					handful.red = num
					miniumNeeded.red = int(math.Max(float64(miniumNeeded.red), float64(num)))
					continue
				}
				if strings.Contains(cube, "green") {
					handful.green = num
					miniumNeeded.green = int(math.Max(float64(miniumNeeded.green), float64(num)))
					continue
				}
				if strings.Contains(cube, "blue") {
					handful.blue = num
					miniumNeeded.blue = int(math.Max(float64(miniumNeeded.blue), float64(num)))
					continue
				}
			}
			if !handful.isPossible() {
				possible = false
			}
		}
		if possible {
			runningTotal += gameId + 1
		}
		runingPower += miniumNeeded.calcPower()
	}
	fmt.Printf("Part 1: %d\n", runningTotal)
	fmt.Printf("Part 2: %d\n", runingPower)
}

type cubeHandful struct {
	red   int
	green int
	blue  int
}

func (c cubeHandful) isPossible() bool {
	return c.red <= 12 && c.green <= 13 && c.blue <= 14
}

func (c cubeHandful) calcPower() int {
	return c.red * c.green * c.blue
}
