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
	total := 0
	for _, line := range input {
		if line == "" {
			continue
		}

		first, remainingBank := getLargestFromBank(line, 1)
		second, _ := getLargestFromBank(remainingBank, 0)

		value := formNumber([]int{first, second})
		total += value
	}
	return total
}

func part2(input []string) int {
	joltageLen := 12

	total := 0
	for _, line := range input {
		if line == "" {
			continue
		}

		bank := line
		digits := []int{}
		for i := 0; i < joltageLen; i++ {
			digit, remainingBank := getLargestFromBank(bank, joltageLen-1-i)
			digits = append(digits, digit)
			bank = remainingBank
		}

		value := formNumber(digits)
		total += value
	}
	return total
}

func formNumber(digits []int) int {
	number := 0
	for _, d := range digits {
		number *= 10
		number += d
	}
	return number
}

func getLargestFromBank(bank string, endPadding int) (int, string) {
	largest := 0
	largestId := -1
	for i, battery := range bank {
		if i == len(bank)-endPadding {
			break
		}

		joltage := helpers.MustAtoi(string(battery))
		if joltage > largest {
			largest = joltage
			largestId = i
		}
	}

	if largestId < 0 {
		panic("ID was less than 0")
	}

	remainingLine := bank[largestId+1:]

	return largest, remainingLine
}
