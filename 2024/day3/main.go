package main

import (
	"fmt"
	"regexp"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	total := 0

	for _, line := range input {
		if line == "" {
			continue
		}
		r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			total += helpers.MustAtoi(match[1]) * helpers.MustAtoi(match[2])
		}
	}

	fmt.Println(total)
}
