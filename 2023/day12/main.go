package main

import (
	"fmt"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	total1, total2 := 0, 0
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

		total1 += calculateRecord(record, groups)
		total2 += calculateRecord(unfold(record, groups, 5))
	}
	fmt.Printf("Part 1: %d\n", total1)
	fmt.Printf("Part 2: %d\n", total2)
}

func unfold(record string, groups []int, foldCount int) (string, []int) {
	unfoldedRecord := ""
	unfoldedGroups := make([]int, 0, len(groups)*5)
	for i := 0; i < foldCount; i++ {
		unfoldedRecord += record
		unfoldedRecord += "?"
		unfoldedGroups = append(unfoldedGroups, groups...)
	}
	unfoldedRecord = unfoldedRecord[:len(unfoldedRecord)-1]
	return unfoldedRecord, unfoldedGroups
}

var recordCache = map[string]int{}

func genKey(record string, desiredGroups []int) string {
	key := ""
	key += record
	key += "|"
	for _, g := range desiredGroups {
		key += fmt.Sprintf("%d", g)
		key += ","
	}
	return key
}

func calculateRecord(record string, desiredGroups []int) int {
	key := genKey(record, desiredGroups)
	combinations, ok := recordCache[key]
	if ok {
		return combinations
	}

	if len(record) == 0 {
		if len(desiredGroups) == 0 {
			retVal := 1
			recordCache[key] = retVal
			return retVal
		}
		retVal := 0
		recordCache[key] = retVal
		return retVal
	}

	switch record[0] {
	case '.':
		retVal := calculateRecord(record[1:], desiredGroups)
		recordCache[key] = retVal
		return retVal
	case '?':
		retVal := calculateRecord("#"+record[1:], desiredGroups) + calculateRecord("."+record[1:], desiredGroups)
		recordCache[key] = retVal
		return retVal
	case '#':
		if len(desiredGroups) == 0 {
			// have more broken springs than the group specifies
			retVal := 0
			recordCache[key] = retVal
			return retVal
		}
		if len(record) < desiredGroups[0] {
			// group speficies a bigger run than remains in the record
			retVal := 0
			recordCache[key] = retVal
			return retVal
		}

		for i := 1; i < desiredGroups[0]; i++ {
			if record[i] == '.' {
				retVal := 0
				recordCache[key] = retVal
				return retVal
			}
		}

		if len(record) == desiredGroups[0] {
			// end of string so don't worry about handling the following characters
			retVal := calculateRecord(record[desiredGroups[0]:], desiredGroups[1:])
			recordCache[key] = retVal
			return retVal
		}

		if record[desiredGroups[0]] == '#' {
			retVal := 0 //the run would be longer than the desired value
			recordCache[key] = retVal
			return retVal
		}

		// return the rest of the string after the found run. And replace the character
		// after the run with an operational to ensure the length of the run isn't more
		// than the required one
		retVal := calculateRecord("."+record[desiredGroups[0]+1:], desiredGroups[1:])
		recordCache[key] = retVal
		return retVal

	default:
		panic(fmt.Errorf("unknown character: %v", record[0]))
	}
}
