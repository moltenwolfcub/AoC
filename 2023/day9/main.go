package main

import (
	"fmt"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	totalNext := 0
	totalPrev := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		sequence := []int{}
		for _, i := range strings.Split(line, " ") {
			sequence = append(sequence, helpers.MustAtoi(i))
		}

		subSeqLast := []int{}
		subSeqFirst := []int{}

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
			subSeqFirst = append(subSeqFirst, sequence[0])

			sequence = findNextSubSeq(sequence)
		}

		next := 0
		for _, i := range subSeqLast {
			next += i
		}

		totalNext += next

		prev := subSeqFirst[len(subSeqFirst)-1]
		for i := len(subSeqFirst) - 2; i >= 0; i-- {
			subSeqFirst[i] -= prev
			prev = subSeqFirst[i]
		}

		totalPrev += prev
	}

	fmt.Printf("Part 1: %d\n", totalNext)
	fmt.Printf("Part 2: %d\n", totalPrev)
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
