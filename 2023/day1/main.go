package main

import (
	"fmt"
	"regexp"
	"unicode"

	helpers "github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	lines := helpers.ReadLines("input.txt")
	fmt.Println("Part 1: " + part1(lines))
	fmt.Println("Part 2: " + part2(lines))
}

func part1(input []string) string {
	calibrationValue := 0
	for lineNum, line := range input {
		if line == "" {
			continue
		}

		var first, last string
		for _, r := range line {
			if unicode.IsDigit(r) {
				first = string(r)
				break
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			r := rune(line[i])
			if unicode.IsDigit(r) {
				last = string(r)
				break
			}
		}
		if first == "" {
			return fmt.Sprintf("No numbers found on line %v", lineNum+1)
		}

		number := helpers.MustAtoi(first)*10 + helpers.MustAtoi(last)
		calibrationValue += number
	}
	return fmt.Sprintf("%v", calibrationValue)
}

func part2(input []string) string {
	calibrationValue := 0

	firstRegex := regexp.MustCompile(`^.*?(\d|one|two|three|four|five|six|seven|eight|nine)`)
	lastRegex := regexp.MustCompile(`.*(\d|one|two|three|four|five|six|seven|eight|nine).*?$`)

	for lineNum, line := range input {
		if line == "" {
			continue
		}

		first := getDigit(firstRegex, line)
		last := getDigit(lastRegex, line)

		if first < 0 || last < 0 {
			return fmt.Sprintf("Couldn't find number on line %d", lineNum)
		}

		calibrationValue += first*10 + last
	}
	return fmt.Sprintf("%v", calibrationValue)
}

func getDigit(regex *regexp.Regexp, line string) int {
	match := regex.FindStringSubmatch(line)
	if match == nil {
		return -1
	}

	strNum := match[1]
	if len(strNum) == 1 {
		return helpers.MustAtoi(strNum)
	}

	switch strNum {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	default:
		return 0
	}
}
