package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	seeds := []int{}
	propertyMaps := [][]mapping{{}, {}, {}, {}, {}, {}, {}}

	inputSeg := 0
	for _, line := range input {
		if line == "" {
			inputSeg++
			continue
		}

		if inputSeg == 0 {
			_, seedData, ok := strings.Cut(line, ": ")
			if !ok {
				panic("couldn't cut seeds list")
			}
			seedNums := strings.Split(seedData, " ")

			for _, seed := range seedNums {
				seeds = append(seeds, helpers.MustAtoi(seed))
			}
			continue
		}

		if strings.Contains(line, ":") {
			continue
		}

		mapData := strings.Split(line, " ")

		lineMapping := mapping{
			valStart: helpers.MustAtoi(mapData[0]),
			keyStart: helpers.MustAtoi(mapData[1]),
			length:   helpers.MustAtoi(mapData[2]),
		}
		propertyMaps[inputSeg-1] = append(propertyMaps[inputSeg-1], lineMapping)
	}

	locations := []int{}
	for _, seed := range seeds {
		oldProp := seed
		for _, propertyMap := range propertyMaps {
			newProp := getValue(propertyMap, oldProp)
			if newProp < 0 {
				panic("Couldn't map property")
			}

			oldProp = newProp
		}
		locations = append(locations, oldProp)
	}
	slices.Sort(locations)
	fmt.Printf("Part 1: %v\n", locations[0])
}

func getValue(propertyMap []mapping, key int) int {
	for _, mapping := range propertyMap {
		v, ok := mapping.getValue(key)
		if ok {
			return v
		}
	}
	return key
}

type mapping struct {
	keyStart int
	valStart int
	length   int
}

func (m mapping) getValue(key int) (int, bool) {
	if m.keyStart <= key && key < m.keyStart+m.length {
		offset := key - m.keyStart
		value := m.valStart + offset
		return value, true
	} else {
		return -1, false
	}
}
