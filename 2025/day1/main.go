package main

import (
	"fmt"
	"regexp"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	re := regexp.MustCompile(`([LR])(\d+)`)

	part1Zeroes := 0
	part2Zeroes := 0

	safeIndex := 50
	for _, line := range input {
		if line == "" {
			continue
		}

		matches := re.FindStringSubmatch(line)

		direction := matches[1]
		amount := helpers.MustAtoi(matches[2])

		wasZero := safeIndex == 0

		switch direction {
		case "R":
			safeIndex += amount
		case "L":
			safeIndex -= amount
		default:
			panic(fmt.Sprintf("Unknown instruction: '%x'", direction))
		}

		originCrosses := 0
		safeIndex, originCrosses = makePrinciple(safeIndex, wasZero)

		part2Zeroes += originCrosses

		if safeIndex == 0 {
			part1Zeroes++
			part2Zeroes++
		}
	}

	fmt.Printf("Part 1: %d\n", part1Zeroes)
	fmt.Printf("Part 2: %d\n", part2Zeroes)
}

func makePrinciple(rot int, wasZero bool) (int, int) {
	originCrosses := 0
	if wasZero && rot < 0 {
		originCrosses-- // gone from zero so don't want to double count first move
	}

	for rot >= 100 || rot < 0 {
		if rot >= 100 {
			if rot != 100 {
				originCrosses++ //didn't cross origin it landed on origin
			}
			rot -= 100
		}
		if rot < 0 {
			rot += 100
			originCrosses++
		}
	}

	return rot, originCrosses
}
