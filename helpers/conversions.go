package helpers

import "strconv"

func MustAtoi(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		panic("Failed to parse number: " + str)
	}
	return value
}
