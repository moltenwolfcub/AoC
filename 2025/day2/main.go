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

	total := 0
	for _, match := range matches {
		rangeStart := helpers.MustAtoi(match[1])
		rangeEnd := helpers.MustAtoi(match[2])

		for i := rangeStart; i <= rangeEnd; i++ {
			result := testID(i)
			total += result
		}
	}

	fmt.Println(total)
}

// returns ID if invalid or 0 if valid
func testID(id int) int {
	re := pcre.MustCompile(`^(\d+)\1$`, 0) //different regex library to support backreferencing

	hasMatch := re.MatcherString(fmt.Sprint(id), 0).Matches()

	if hasMatch {
		return id
	}
	return 0
}
