package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	lines := helpers.ReadLines("input.txt")

	worth := 0

	for _, line := range lines {
		if line == "" {
			continue
		}
		_, lineData, didCut := strings.Cut(line, ": ")
		if !didCut {
			panic("Couldn't properly cut string")
		}

		numberSets := strings.Split(lineData, "|")

		winningStrNums := strings.Split(numberSets[0], " ")
		myStrNums := strings.Split(numberSets[1], " ")

		winningNums := []int{}
		myNums := []int{}

		for _, strNum := range winningStrNums {
			if strNum == "" {
				continue
			}

			winningNums = append(winningNums, helpers.MustAtoi(strNum))
		}
		for _, strNum := range myStrNums {
			if strNum == "" {
				continue
			}

			myNums = append(myNums, helpers.MustAtoi(strNum))
		}

		wins := len(intersection(winningNums, myNums))

		cardWorth := math.Pow(2, float64(wins-1))
		worth += int(cardWorth)
	}

	fmt.Printf("Part 1: %d\n", worth)
}

func intersection(a []int, b []int) []int {
	ret := []int{}
	for _, aVal := range a {
		if slices.Contains(b, aVal) {
			ret = append(ret, aVal)
		}
	}
	return ret
}
