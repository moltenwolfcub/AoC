package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	fmt.Printf("Part 1: %v\n", part1(input, r))
	fmt.Printf("Part 2: %v\n", part2(input, r))
}

func part1(input []string, r *regexp.Regexp) int {
	total := 0

	for _, line := range input {
		if line == "" {
			continue
		}

		matches := r.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			total += helpers.MustAtoi(match[1]) * helpers.MustAtoi(match[2])
		}
	}

	return total
}

func part2(input []string, r *regexp.Regexp) int {
	total := 0

	combinedStr := ""
	for _, line := range input {
		if line == "" {
			continue
		}
		combinedStr += line
	}

	dontParts := strings.Split(combinedStr, "don't()")

	filteredLine := ""
	for i, section := range dontParts {
		if i == 0 {
			filteredLine += section
			continue
		}
		for i, doPart := range strings.Split(section, "do()") {
			if i == 0 {
				continue
			}
			filteredLine += doPart
		}
	}

	matches := r.FindAllStringSubmatch(filteredLine, -1)
	for _, match := range matches {
		total += helpers.MustAtoi(match[1]) * helpers.MustAtoi(match[2])
	}

	return total
}
