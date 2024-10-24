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
	bits := make([]int, len(input[0]))

	for _, line := range input {
		if line == "" {
			continue
		}
		for i, r := range line {
			c := string(r)
			n := helpers.MustAtoi(c)
			if n == 1 {
				bits[i] += 1
			} else {
				bits[i] -= 1
			}
		}
	}

	gamma := 0
	epsilon := 0
	for i, b := range bits {
		pos := len(bits) - i - 1

		var n int
		if b > 0 {
			n = 1
		}

		gamma += n << pos
		epsilon += (1 ^ n) << pos
	}

	return epsilon * gamma
}

func part2(input []string) int {
	oxygen := input
	for i := 0; i < len(input[0]); i++ {
		zeros := []string{}
		ones := []string{}
		for _, line := range oxygen {
			if line == "" {
				continue
			}
			c := helpers.MustAtoi(string(line[i]))
			if c == 1 {
				ones = append(ones, line)
			} else {
				zeros = append(zeros, line)
			}
		}
		if len(zeros) > len(ones) {
			oxygen = zeros
		} else {
			oxygen = ones
		}

		if len(oxygen) == 1 {
			break
		}
	}

	co2 := input
	for i := 0; i < len(input[0]); i++ {
		zeros := []string{}
		ones := []string{}
		for _, line := range co2 {
			if line == "" {
				continue
			}
			c := helpers.MustAtoi(string(line[i]))
			if c == 1 {
				ones = append(ones, line)
			} else {
				zeros = append(zeros, line)
			}
		}
		if len(ones) < len(zeros) {
			co2 = ones
		} else {
			co2 = zeros
		}

		if len(co2) == 1 {
			break
		}
	}

	return binToDec(co2[0]) * binToDec(oxygen[0])
}

func binToDec(bin string) (output int) {
	for i, b := range bin {
		pos := len(bin) - i - 1

		n := helpers.MustAtoi(string(b))

		output += n << pos
	}
	return
}
