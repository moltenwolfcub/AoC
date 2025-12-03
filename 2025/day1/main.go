package main

import (
	"fmt"
	"regexp"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	re := regexp.MustCompile(`([LR])(\d+)`)

	zeros := 0

	safeIndex := 50
	for _, line := range input {
		if line == "" {
			continue
		}

		matches := re.FindStringSubmatch(line)

		direction := matches[1]
		amount := helpers.MustAtoi(matches[2])

		switch direction {
		case "R":
			safeIndex += amount
		case "L":
			safeIndex -= amount
		default:
			panic(fmt.Sprintf("Unknown instruction: '%x'", direction))
		}

		safeIndex = makePrinciple(safeIndex)

		if safeIndex == 0 {
			zeros++
		}
	}
	fmt.Println(zeros)
}

func makePrinciple(rot int) int {
	for rot >= 100 || rot < 0 {
		if rot >= 100 {
			rot -= 100
		}
		if rot < 0 {
			rot += 100
		}
	}
	return rot
}
