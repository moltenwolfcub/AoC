package main

import (
	"fmt"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

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

	fmt.Println(epsilon * gamma)
}
