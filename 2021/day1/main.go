package main

import (
	"fmt"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")
	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []string) int {
	lastValue := -1
	runningTotal := 0

	for _, line := range input {
		if line == "" {
			continue
		}
		num := helpers.MustAtoi(line)
		if lastValue != -1 {
			if num > lastValue {
				runningTotal++
			}
		}
		lastValue = num
	}
	return runningTotal
}

func part2(input []string) int {
	lastValues := make([]int, 3)
	runningTotal := 0

	for _, line := range input {
		if line == "" {
			continue
		}
		num := helpers.MustAtoi(line)

		theseValues := append(lastValues[1:3], num)
		if listFilled(lastValues) {
			if sum(theseValues) > sum(lastValues) {
				runningTotal++
			}
		}
		lastValues = theseValues
	}
	return runningTotal
}

func sum(l []int) (total int) {
	for _, i := range l {
		total += i
	}
	return
}

func listFilled(l []int) bool {
	filled := 0
	for _, i := range l {
		if i != 0 {
			filled++
		}
	}
	return filled >= 3
}
