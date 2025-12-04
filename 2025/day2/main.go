package main

import (
	"fmt"
	"regexp"

	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")[0]

	re := regexp.MustCompile(`(\d+)-(\d+)`)

	matches := re.FindAllStringSubmatch(input, -1)

	total1 := 0
	total2 := 0
	for _, match := range matches {
		rangeStart := helpers.MustAtoi(match[1])
		rangeEnd := helpers.MustAtoi(match[2])

		for i := rangeStart; i <= rangeEnd; i++ {
			result1, result2 := testID(i)
			total1 += result1
			total2 += result2
		}
	}

	fmt.Printf("Part 1: %d\n", total1)
	fmt.Printf("Part 2: %d\n", total2)
}

// returns ID if invalid or 0 if valid
func testID(id int) (result1, result2 int) {
	re1 := pcre.MustCompile(`^(\d+)\1$`, 0) //different regex library to support backreferencing
	hasMatch1 := re1.MatcherString(fmt.Sprint(id), 0).Matches()
	if hasMatch1 {
		result1 = id
	}

	re2 := pcre.MustCompile(`^(\d+)\1+$`, 0) //different regex library to support backreferencing
	hasMatch2 := re2.MatcherString(fmt.Sprint(id), 0).Matches()
	if hasMatch2 {
		result2 = id
	}

	return
}
