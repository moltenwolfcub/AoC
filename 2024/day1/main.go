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

	fmt.Printf("Part 1: %v\n", part1(left, right))
	fmt.Printf("Part 2: %v\n", part2(left, right))
}

func part1(left, right []int) int {

	total := 0

	for i := 0; i < len(left); i++ {
		diff := left[i] - right[i]
		if diff < 0 {
			diff = -diff
		}

		total += diff
	}
	return total
}

func part2(left, right []int) int {
	total := 0
	lastLeft := -1
	lastValue := -1

	for _, l := range left {
		if l == lastLeft {
			total += lastValue
			continue
		}

		similarity := 0

		redunant := 0

		for _, r := range right {
			if r < l {
				redunant++
				continue
			}
			if r == l {
				similarity++
				continue
			}
			break
		}
		right = right[similarity+redunant:]

		lastLeft = l
		lastValue = l * similarity
		total += lastValue
	}

	return total
}
