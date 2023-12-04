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
	cards := map[int]scratchCard{}

	for lineNum, line := range lines {
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

		cards[lineNum+1] = scratchCard{
			value: wins,
			index: lineNum + 1,
		}
	}
	fmt.Printf("Part 1: %d\n", worth)

	cardsToProcess := []scratchCard{}
	for _, v := range cards {
		cardsToProcess = append(cardsToProcess, v)
	}

	allCards := cardsToProcess
	for len(cardsToProcess) > 0 {
		current := cardsToProcess[0]

		if current.value > 0 {

			for i := current.index + 1; i <= current.index+current.value; i++ {
				cardsToProcess = append(cardsToProcess, cards[i])
				allCards = append(allCards, cards[i])
			}
		}

		cardsToProcess = cardsToProcess[1:]
	}
	fmt.Printf("Part 2: %d\n", len(allCards))
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

type scratchCard struct {
	index int
	value int
}
