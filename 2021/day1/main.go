package main

import (
	"fmt"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

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
	fmt.Println(runningTotal)
}
