package main

import (
	"fmt"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("example1.txt")

	arrays := [][][]bool{}

	var currentArray [][]bool
	for _, line := range input {
		if line == "" {
			arrays = append(arrays, currentArray)
			currentArray = nil
			continue
		}

		yArray := make([]bool, len(line))
		for x, c := range line {
			if c == '#' {
				yArray[x] = true
				continue
			}
			if c == '.' {
				yArray[x] = false
				continue
			}
			panic("unknown character")
		}
		currentArray = append(currentArray, yArray)
	}
	fmt.Println(arrays)
}
