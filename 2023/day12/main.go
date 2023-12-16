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

		splitLine := strings.Split(line, " ")

		groups := []int{}
		for _, g := range strings.Split(splitLine[1], ",") {
			groups = append(groups, helpers.MustAtoi(g))
		}

		record := splitLine[0]

		total += calculateRecord(record, groups)
	}
	fmt.Printf("Part 1: %d\n", total)
}

func calculateRecord(record string, desiredGroups []int) int {
	if len(record) == 0 {
		if len(desiredGroups) == 0 {
			return 1
		}
		return 0
	}

	switch record[0] {
	case '.':
		return calculateRecord(record[1:], desiredGroups)
	case '?':
		return calculateRecord("#"+record[1:], desiredGroups) + calculateRecord("."+record[1:], desiredGroups)
	case '#':
		if len(desiredGroups) == 0 {
			// have more broken springs than the group specifies
			return 0
		}
		if len(record) < desiredGroups[0] {
			// group speficies a bigger run than remains in the record
			return 0
		}

		for i := 1; i < desiredGroups[0]; i++ {
			if record[i] == '.' {
				return 0
			}
		}

		if len(record) == desiredGroups[0] {
			// end of string so don't worry about handling the following characters
			return calculateRecord(record[desiredGroups[0]:], desiredGroups[1:])
		}

		if record[desiredGroups[0]] == '#' {
			return 0 //the run would be longer than the desired value
		}

		// return the rest of the string after the found run. And replace the character
		// after the run with an operational to ensure the length of the run isn't more
		// than the required one
		return calculateRecord("."+record[desiredGroups[0]+1:], desiredGroups[1:])

	default:
		panic(fmt.Errorf("unknown character: %v", record[0]))
	}
}
