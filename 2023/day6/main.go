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

	fmt.Printf("Part 1: %d\n", part1(times, records))

	strTime := ""
	for _, i := range times {
		strTime += fmt.Sprintf("%d", i)
	}
	time := helpers.MustAtoi(strTime)

	strRecord := ""
	for _, i := range records {
		strRecord += fmt.Sprintf("%d", i)
	}
	record := helpers.MustAtoi(strRecord)

	fmt.Printf("Part 2: %d\n", part2(time, record))
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

func part1(times []int, records []int) int {
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
	return total
}

func part2(time int, record int) int {

	margin := 0
	for i := 1; i < time; i++ {
		speed := i
		duration := time - i
		distTravelled := speed * duration

		if distTravelled > record {
			margin++
		}
	}

	return margin
}
