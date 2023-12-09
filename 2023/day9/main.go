package main

import (
	"fmt"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	total := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		sequence := []int{}
		for _, i := range strings.Split(line, " ") {
			sequence = append(sequence, helpers.MustAtoi(i))
		}

		subSeqLast := []int{}

		for {
			zero := true
			for _, i := range sequence {
				if i != 0 {
					zero = false
					break
				}
			}
			if zero {
				break
			}

			subSeqLast = append(subSeqLast, sequence[len(sequence)-1])

			sequence = findNextSubSeq(sequence)
		}

		next := 0
		for _, i := range subSeqLast {
			next += i
		}

		total += next
	}

	fmt.Printf("Part 1: %d\n", total)
}

func findNextSubSeq(seq []int) []int {
	ret := []int{}
	last := seq[0]
	for _, cur := range seq[1:] {
		ret = append(ret, cur-last)
		last = cur
	}
	return ret
}
