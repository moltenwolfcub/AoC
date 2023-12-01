package main

import (
	"fmt"
	"unicode"

	helpers "github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	lines := helpers.ReadLines("input.txt")

	calibrationValue := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		var first, last string
		for _, r := range line {
			if unicode.IsDigit(r) {
				first = string(r)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			r := rune(line[i])
			if unicode.IsDigit(r) {
				last = string(r)
				break
			}
		}

		number := helpers.MustAtoi(first)*10 + helpers.MustAtoi(last)
		calibrationValue += number
	}
	fmt.Println(calibrationValue)
}
