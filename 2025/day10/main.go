package main

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

type LightArray int

func LightArrayFromStr(s string) LightArray {
	r := 0
	for i := len(s) - 1; i >= 0; i-- {
		r <<= 1
		ch := s[i]
		if ch == '#' {
			r++
		}
	}
	return LightArray(r)
}

func (l LightArray) String() string {
	return strconv.FormatInt(int64(l), 2)
}

type Button struct {
	toggleMap LightArray
}

func (b Button) String() string {
	return b.toggleMap.String()
}

func main() {
	input := helpers.ReadLines("input.txt")

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

var targetRe = regexp.MustCompile(`\[([\.|#]+)\]`)
var buttonRe = regexp.MustCompile(`\(((?:\d,?)+)\)`)

func part1(input []string) int {
	runningTotal := 0
	for _, line := range input {
		if line == "" {
			continue
		}

		inputTarget := targetRe.FindStringSubmatch(line)
		target := LightArrayFromStr(inputTarget[1])

		inputButtons := buttonRe.FindAllStringSubmatch(line, -1)

		buttons := make([]Button, len(inputButtons))
		for i, b := range inputButtons {
			lights := strings.Split(b[1], ",")

			button := Button{}
			for _, l := range lights {
				lightPos := helpers.MustAtoi(l)
				button.toggleMap ^= (1 << lightPos)
			}
			buttons[i] = button
		}
		sol := findSolution(target, buttons)
		runningTotal += sol
	}
	return runningTotal
}

func findSolution(target LightArray, buttons []Button) int {
	if target == 0 {
		return 0
	}

	candidates := []*Move{}
	for _, b := range buttons {
		candidates = append(candidates, &Move{
			state:          LightArray(0),
			Button:         b,
			alreadyPressed: make([]Button, 0),
			presses:        0,
		})
	}

	for len(candidates) > 0 {
		m := candidates[0]
		candidates = candidates[1:]

		if slices.Contains(m.alreadyPressed, m.Button) {
			continue
		}

		m.press()

		if m.state^target == 0 {
			return m.presses
		}
		for _, b := range buttons {
			candidates = append(candidates, &Move{
				state:          m.state,
				Button:         b,
				alreadyPressed: append(m.alreadyPressed, m.Button),
				presses:        m.presses,
			})
		}

		if m.presses > len(buttons) {
			panic("no possible solution")
		}
	}

	panic("no solution found")
}

type Move struct {
	state          LightArray
	Button         Button
	alreadyPressed []Button
	presses        int
}

func (m *Move) press() {
	m.state ^= m.Button.toggleMap
	m.presses++
}
