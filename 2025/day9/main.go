package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

type Tile struct {
	X, Y int
}

func (t Tile) Size() int {
	return t.X * t.Y
}

func main() {
	input := helpers.ReadLines("input.txt")

	tiles := make([]Tile, len(input)-1)
	for i, line := range input {
		if line == "" {
			continue
		}

		coords := strings.Split(line, ",")
		tiles[i] = Tile{helpers.MustAtoi(coords[0]), helpers.MustAtoi(coords[1])}
	}

	fmt.Printf("Part 1: %v\n", part1(tiles))
	fmt.Printf("Part 2: %v\n", part2(tiles))
}

func part1(tiles []Tile) int {

	greatestArea := 0
	for f := 0; f < len(tiles); f++ {
		first := tiles[f]
		for s := f + 1; s < len(tiles); s++ {
			second := tiles[s]

			a := area(first, second)
			greatestArea = max(greatestArea, a)
		}
	}
	return greatestArea
}

func part2(tiles []Tile) int {

	rects := []Rect{}
	for f := 0; f < len(tiles); f++ {
		first := tiles[f]
		for s := f + 1; s < len(tiles); s++ {
			second := tiles[s]

			rects = append(rects, NewRect(first, second))
		}
	}

	sort.Slice(rects, func(i, j int) bool {
		return rects[i].area > rects[j].area
	})

	lines := createLines(tiles)

	greatestArea := 0
outer:
	for _, r := range rects {
		for _, l := range lines {
			if !r.Touches(l) {
				continue
			}
			if r.Contains(l) {
				continue outer
			}
			if r.Touches(l.left()) {
				continue outer
			}
		}
		greatestArea = r.area
		break
	}

	return greatestArea
}

func area(first, second Tile) int {
	w := helpers.IntAbs(second.X-first.X) + 1
	h := helpers.IntAbs(second.Y-first.Y) + 1
	return w * h
}

type Rect struct {
	Start, End Tile
	area       int
}

func NewRect(t1, t2 Tile) Rect {
	r := Rect{}
	r.Start = Tile{min(t1.X, t2.X), min(t1.Y, t2.Y)}
	r.End = Tile{max(t1.X, t2.X), max(t1.Y, t2.Y)}
	r.area = area(r.Start, r.End)

	return r
}

func (r Rect) Touches(l Line) bool {
	if l.Start.X < r.Start.X && l.End.X < r.Start.X {
		return false
	}
	if l.Start.Y < r.Start.Y && l.End.Y < r.Start.Y {
		return false
	}
	if l.Start.X > r.End.X && l.End.X > r.End.X {
		return false
	}
	if l.Start.Y > r.End.Y && l.End.Y > r.End.Y {
		return false
	}
	return true
}

func (r Rect) Contains(l Line) bool {
	n := Rect{
		Start: Tile{r.Start.X + 1, r.Start.Y + 1},
		End:   Tile{r.End.X - 1, r.End.Y - 1},
	}

	if l.Start.X < n.Start.X && l.End.X < n.Start.X {
		return false
	}
	if l.Start.Y < n.Start.Y && l.End.Y < n.Start.Y {
		return false
	}
	if l.Start.X > n.End.X && l.End.X > n.End.X {
		return false
	}
	if l.Start.Y > n.End.Y && l.End.Y > n.End.Y {
		return false
	}
	return true
}

type Line struct {
	Start, End Tile
}

func (l Line) left() Line {
	newLine := Line{}

	if l.Start.Y == l.End.Y && l.Start.X < l.End.X {
		//North
		newLine.Start = Tile{newLine.Start.X, newLine.Start.Y - 1}
		newLine.End = Tile{newLine.End.X, newLine.End.Y - 1}

	} else if l.Start.Y == l.End.Y && l.Start.X > l.End.X {
		//south
		newLine.Start = Tile{newLine.Start.X, newLine.Start.Y + 1}
		newLine.End = Tile{newLine.End.X, newLine.End.Y + 1}
	} else if l.Start.X == l.End.X && l.Start.Y < l.End.Y {
		//east
		newLine.Start = Tile{newLine.Start.X + 1, newLine.Start.Y}
		newLine.End = Tile{newLine.End.X + 1, newLine.End.Y}
	} else if l.Start.X == l.End.X && l.Start.Y > l.End.Y {
		//west
		newLine.Start = Tile{newLine.Start.X - 1, newLine.Start.Y}
		newLine.End = Tile{newLine.End.X - 1, newLine.End.Y}
	}

	return newLine
}

func createLines(tiles []Tile) []Line {
	lines := make([]Line, len(tiles))
	for i := 0; i < len(tiles); i++ {
		start := tiles[i]
		var end Tile
		if i != len(tiles)-1 {
			end = tiles[i+1]
		} else {
			end = tiles[0]
		}
		lines[i] = Line{start, end}
	}
	return lines
}
