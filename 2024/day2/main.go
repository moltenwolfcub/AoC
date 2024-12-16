package main

import (
	"fmt"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

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
	fmt.Println(safe)
}
