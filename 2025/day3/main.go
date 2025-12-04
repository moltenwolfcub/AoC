package main

import (
	"fmt"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	total := 0
	for _, line := range input {
		if line == "" {
			continue
		}

		first, remainingBank := getLargestFromBank(line, 1)
		second, _ := getLargestFromBank(remainingBank, 0)

		value := first*10 + second
		total += value
	}
	fmt.Println(total)
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
