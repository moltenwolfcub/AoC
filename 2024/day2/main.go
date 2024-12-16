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

	safe := 0

outer:
	for _, line := range input {
		if line == "" {
			continue
		}

		rawLevels := strings.Split(line, " ")
		levels := make([]int, len(rawLevels))
		for i, l := range rawLevels {
			levels[i] = helpers.MustAtoi(l)
		}

		var dir bool // descending = true; ascending = false
		if levels[0] < levels[1] {
			dir = false
		} else if levels[0] > levels[1] {
			dir = true
		} else {
			// same so unstable
			continue
		}

		last := -1
		for _, this := range levels {

			if last == -1 {
				last = this
				continue
			}

			if last == this {
				// same so unstable
				continue outer
			}

			if dir {
				if this > last {
					continue outer
				}
			} else {
				if this < last {
					continue outer
				}
			}

			if helpers.IntAbs(this-last) > 3 {
				continue outer
			}

			last = this
		}
		safe++
	}
	return safe
}

func part2(input []string) int {
	safe := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		rawLevels := strings.Split(line, " ")
		levels := make([]int, len(rawLevels))
		for i, l := range rawLevels {
			levels[i] = helpers.MustAtoi(l)
		}

		dampenedLevels := make([][]int, len(levels)+1)
		dampenedLevels[0] = levels

		for i := 0; i < len(levels); i++ {
			first := levels[:i]
			second := levels[i+1:]

			new := []int{}

			new = append(new, first...)
			new = append(new, second...)
			dampenedLevels[i+1] = new
		}

	nextSet:
		for _, levelset := range dampenedLevels {
			last := -1
			dir := 0 // descending = 1; ascending = -1

			for _, this := range levelset {

				if last == -1 {
					last = this
					continue
				}
				if dir == 0 {
					if last < this {
						dir = -1
					} else if last > this {
						dir = 1
					}
				}

				if last == this {
					// same so unstable
					continue nextSet
				}

				if dir == 1 {
					if this > last {
						continue nextSet
					}
				} else {
					if this < last {
						continue nextSet
					}
				}

				if helpers.IntAbs(this-last) > 3 {
					continue nextSet
				}

				last = this
			}
			safe++
			break
		}
	}
	return safe
}
