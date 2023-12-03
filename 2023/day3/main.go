package main

import (
	"fmt"
	"image"
	"unicode"

	"github.com/moltenwolfcub/AoC/helpers"
)

/*
	2 passes

	1st pass:
	Go through each line and load the schematic into a data-strucutre
	This will probably be made of:
		- a list of numbers which will be stored as a rectangle and a value
		- a list of points for the symbols

	2nd pass:
	Go through each symbol in the list and construct a rectangle containing
	all of the adjacent points (+1 & -1 on x&y) and use Overlaps() on each
	number rectangle in that list. If they overlap then add their value to
	the total.
*/

func main() {
	input := helpers.ReadLines("input.txt")

	numbers := []Number{}
	symbols := []Symbol{}

	for y, line := range input {
		var currentDigits string
		var currentNumStart image.Point
		for x := 0; x < len(line); x++ {
			c := rune(line[x])
			if len(currentDigits) > 0 && !unicode.IsDigit(c) {
				found := Number{
					pos: image.Rectangle{
						Min: currentNumStart,
						Max: image.Pt(x-1, y),
					},
					value: helpers.MustAtoi(currentDigits),
				}
				numbers = append(numbers, found)
				currentDigits = ""
			}

			if c == '.' {
				continue
			}

			if unicode.IsDigit(c) {
				if len(currentDigits) == 0 {
					currentNumStart = image.Pt(x, y)
				}
				currentDigits += string(c)
				continue
			}
			found := Symbol{
				pos: image.Pt(x, y),
			}

			symbols = append(symbols, found)
		}
		if len(currentDigits) > 0 {
			found := Number{
				pos: image.Rectangle{
					Min: currentNumStart,
					Max: image.Pt(len(line)-1, y),
				},
				value: helpers.MustAtoi(currentDigits),
			}
			numbers = append(numbers, found)
			currentDigits = ""
		}
	}

	sum := 0
	for _, num := range numbers {
		isPartNum := false
		for _, symbol := range symbols {
			if num.touching(symbol.pos) {
				isPartNum = true
				break
			}
		}
		if isPartNum {
			sum += num.value
		}
	}
	fmt.Println(sum)
}

type Number struct {
	pos   image.Rectangle
	value int
}

func (n Number) touching(pt image.Point) bool {
	scaledRect := image.Rectangle{
		Min: n.pos.Min.Sub(image.Pt(1, 1)),
		Max: n.pos.Max.Add(image.Pt(2, 2)), //2 cause check is < instead of <=
	}

	return pt.In(scaledRect)
}

type Symbol struct {
	pos image.Point
}
