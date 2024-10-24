package main

import (
	"fmt"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	depth := 0
	hPos := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		instruction := strings.Split(line, " ")
		amount := helpers.MustAtoi(instruction[1])

		switch instruction[0] {
		case "forward":
			hPos += amount
		case "up":
			depth -= amount
		case "down":
			depth += amount
		}
	}
	fmt.Println(depth * hPos)
}
