package main

import (
	"fmt"
	"image"
	"math"
	"slices"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	seeds := []int{}
	seedRanges := []seedRange{}
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
			for i := 0; i < len(seedNums); i += 2 {
				sRange := seedRange{
					start:       helpers.MustAtoi(seedNums[i]),
					length:      helpers.MustAtoi(seedNums[i+1]),
					originStart: helpers.MustAtoi(seedNums[i]),
				}
				seedRanges = append(seedRanges, sRange)
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
	fmt.Printf("Part 1: %v\n", part1(seeds, propertyMaps))
	fmt.Printf("Part 2: %v\n", part2(seedRanges, propertyMaps))
}

func part1(seeds []int, Maps [][]mapping) int {
	locations := []int{}
	for _, seed := range seeds {
		oldProp := seed
		for _, propertyMap := range Maps {
			oldProp = getValue(propertyMap, oldProp)
		}
		locations = append(locations, oldProp)
	}
	slices.Sort(locations)
	return locations[0]
}

func part2(seedRanges []seedRange, Maps [][]mapping) int {
	/*
		Go through each range:

		Go through each mappingLayer:

		Check how many of the mappings it intersects{
			if 1 then just raw map and keep new range
			if more than 1 then split into sections so that
				each section only intersects 1 mapping range and
				continue processing with multiple maps
		}
	*/

	allLocs := []seedRange{}

	for _, sRange := range seedRanges {

		ranges := []seedRange{sRange}
		mapped := []seedRange{}

		for _, mappingLayer := range Maps {

			for len(ranges) > 0 {
				found, rest, overlaps := split(mappingLayer, ranges[0])

				if overlaps {
					ranges = append(ranges, rest...)

					mapped = append(mapped, found.mapRange(mappingLayer))
				} else {
					// if the segment overlaps none then move it into mapped
					// for next mapping layer
					mapped = append(mapped, rest...)
				}

				ranges = ranges[1:]
			}
			ranges = mapped
			mapped = []seedRange{}
		}
		allLocs = append(allLocs, ranges...)
	}

	minLoc := math.MaxInt

	for _, loc := range allLocs {
		minLoc = int(math.Min(float64(minLoc), float64(loc.start)))
	}

	return minLoc
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

func split(mappingSet []mapping, sRange seedRange) (found seedRange, rest []seedRange, overlapsAny bool) {
	hasFound := false
	for _, m := range mappingSet {
		intersects, intersection := seedRangeFromRect(m.toRect()).intersects(sRange)
		if intersects {
			found = intersection
			hasFound = true
			break
		}
	}
	if !hasFound {
		return seedRange{}, []seedRange{sRange}, false
	}

	rest = []seedRange{}

	preLen := found.start - sRange.start
	if preLen > 0 {
		found.originStart += preLen

		pre := seedRange{
			start:       sRange.start,
			length:      preLen,
			originStart: sRange.originStart,
		}
		rest = append(rest, pre)
	}

	postLen := (sRange.start + sRange.length) - (found.start + found.length)
	if postLen > 0 {
		post := seedRange{
			start:       found.start + found.length,
			length:      postLen,
			originStart: sRange.originStart + (found.start - sRange.start) + found.length,
		}
		rest = append(rest, post)
	}

	return found, rest, true
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

func (m mapping) toRect() image.Rectangle {
	return image.Rect(m.keyStart, -1, m.keyStart+m.length, 1)
}

type seedRange struct {
	start       int
	length      int
	originStart int
}

// assume only intersecting one of the maps in m
func (s seedRange) mapRange(m []mapping) seedRange {
	for _, mapping := range m {
		if mapping.keyStart <= s.start && s.start < mapping.keyStart+mapping.length {
			offset := s.start - mapping.keyStart

			newStart := mapping.valStart + offset

			return seedRange{
				start:       newStart,
				length:      s.length,
				originStart: s.originStart,
			}
		}
	}
	panic("Couldn't mapRange. Make sure map only intersects 1 section")
}

func seedRangeFromRect(r image.Rectangle) seedRange {
	return seedRange{
		start:  r.Min.X,
		length: r.Max.X - r.Min.X,
	}
}

func (s seedRange) toRect() image.Rectangle {
	return image.Rect(s.start, -1, s.start+s.length, 1)
}

func (s seedRange) intersects(o seedRange) (bool, seedRange) {
	sRect := s.toRect()
	oRect := o.toRect()

	doesIntersect := sRect.Overlaps(oRect)

	intersectRect := sRect.Intersect(oRect)
	intersection := seedRangeFromRect(intersectRect)
	intersection.originStart = o.originStart

	return doesIntersect, intersection
}
