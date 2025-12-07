package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

type IdRange struct {
	begin, end int
}

func (thisRange IdRange) contains(id int) bool {
	return id >= thisRange.begin && id <= thisRange.end
}

func (thisRange IdRange) intersects(otherRange IdRange) bool {
	if otherRange.end < thisRange.begin || otherRange.begin > thisRange.end {
		return false
	}
	return true
}

// will produce weird results if not checked for intersect first
func (thisRange IdRange) merge(otherRange IdRange) IdRange {
	newBegin := min(thisRange.begin, otherRange.begin)
	newEnd := max(thisRange.end, otherRange.end)

	return IdRange{newBegin, newEnd}
}

func (thisRange IdRange) size() int {
	return thisRange.end - thisRange.begin + 1 //+1 bc inclusive
}

func main() {
	input := helpers.ReadLines("input.txt")

	freshRanges := []IdRange{}
	availableIngredients := []int{}

	for _, line := range input {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "-")

		switch len(parts) {
		case 2:
			freshRanges = append(freshRanges, IdRange{
				helpers.MustAtoi(parts[0]),
				helpers.MustAtoi(parts[1]),
			})
		case 1:
			availableIngredients = append(availableIngredients, helpers.MustAtoi(parts[0]))
		default:
			panic("something went wrong parsing")
		}
	}

	fmt.Printf("Part 1: %v\n", part1(freshRanges, availableIngredients))
	fmt.Printf("Part 2: %v\n", part2(freshRanges))

}

func part1(freshRanges []IdRange, availableIngredients []int) int {
	freshCount := 0
	for _, i := range availableIngredients {
		for _, r := range freshRanges {
			if r.contains(i) {
				freshCount++
				break
			}
		}
	}
	return freshCount
}

func part2(freshRanges []IdRange) int {
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i].begin < freshRanges[j].begin
	})

	mergedRanges := []IdRange{}

	for _, r := range freshRanges {
		done := false
		for !done {
			intersects := false
			for i, m := range mergedRanges {
				if r.intersects(m) {
					intersects = true
					combined := r.merge(m)
					r = combined
					mergedRanges = append(mergedRanges[:i], mergedRanges[i+1:]...)
					break
				}
			}

			if !intersects {
				mergedRanges = append(mergedRanges, r)
				done = true
			}
		}
	}

	freshIds := 0
	for _, r := range mergedRanges {
		freshIds += r.size()
	}

	return freshIds
}
