package main

import (
	"fmt"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	times := parseData(input[0])
	records := parseData(input[1])

	if len(times) != len(records) {
		panic("Failed to parse data. Lists weren't same length")
	}

	total := 1
	for i := 0; i < len(times); i++ {
		time := times[i]
		record := records[i]

		margin := 0
		for i := 1; i < time; i++ {
			speed := i
			duration := time - i
			distTravelled := speed * duration

			if distTravelled > record {
				margin++
			}
		}
		total *= margin
	}

	fmt.Printf("Part 1: %d\n", total)
}

func parseData(data string) []int {
	_, cutData, ok := strings.Cut(data, ":")
	if !ok {
		panic("Couldn't cut data")
	}
	splitData := strings.Split(strings.TrimSpace(cutData), " ")

	finalData := []int{}
	for _, d := range splitData {
		if d == "" {
			continue
		}
		finalData = append(finalData, helpers.MustAtoi(d))
	}

	return finalData
}
