package main

import (
	"fmt"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")
	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []string) int {
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
	return depth * hPos
}

func part2(input []string) int {
	depth := 0
	hPos := 0

	aim := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		instruction := strings.Split(line, " ")
		amount := helpers.MustAtoi(instruction[1])

		switch instruction[0] {
		case "forward":
			hPos += amount
			depth += amount * aim
		case "up":
			aim -= amount
		case "down":
			aim += amount
		}
	}
	return depth * hPos
}
