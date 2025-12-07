package main

import (
	"fmt"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

type IdRange struct {
	begin, end int
}

func (r IdRange) contains(id int) bool {
	return id >= r.begin && id <= r.end
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

	freshCount := 0
	for _, i := range availableIngredients {
		for _, r := range freshRanges {
			if r.contains(i) {
				freshCount++
				break
			}
		}
	}
	fmt.Println(freshCount)
}
