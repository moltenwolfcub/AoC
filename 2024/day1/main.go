package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	left := []int{}
	right := []int{}

	for _, line := range input {
		if line == "" {
			continue
		}
		nums := strings.Split(line, "   ")
		left = append(left, helpers.MustAtoi(nums[0]))
		right = append(right, helpers.MustAtoi(nums[1]))
	}

	slices.Sort(left)
	slices.Sort(right)

	total := 0

	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}

		total += diff
	}
	fmt.Println(total)
}
