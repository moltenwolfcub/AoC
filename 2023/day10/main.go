package main

import (
	"fmt"
	"image"
	"slices"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	grid := [][]tile{}
	startLoc := image.Pt(-1, -1)

	for y, line := range input {
		if line == "" {
			continue
		}
		gridLine := []tile{}

		for x, c := range line {
			tile := tile{
				pos:   image.Pt(x, y),
				tType: typeFromRune(c),
			}
			gridLine = append(gridLine, tile)
			if tile.tType == startPos {
				if startLoc != image.Pt(-1, -1) {
					panic("more than one startpos found")
				}

				startLoc = tile.pos
			}
		}
		grid = append(grid, gridLine)
	}
	if startLoc == image.Pt(-1, -1) {
		panic("couldn't find startpos")
	}

	originalStart := grid[startLoc.Y][startLoc.X]
	startNeighbours := []image.Point{}

	checkDirForStart := func(offset image.Point) {
		pos := startLoc.Add(offset)
		if pos.X < 0 || pos.Y < 0 {
			return
		}

		neighbour := grid[pos.Y][pos.X]
		if neighbour.tType == ground {
			return
		}

		if slices.Contains(neighbour.getNeighbours(), startLoc) {
			startNeighbours = append(startNeighbours, offset)
		}
	}

	checkDirForStart(image.Pt(0, 1))
	checkDirForStart(image.Pt(0, -1))
	checkDirForStart(image.Pt(1, 0))
	checkDirForStart(image.Pt(-1, 0))

	start := tile{
		pos:   originalStart.pos,
		tType: typeFromNeighbours(startNeighbours),
	}
	checked := []image.Point{}
	{
		current := []tile{start}
		next := []tile{}

		stepsAway := 0

		for len(current) > 0 {
			for _, n := range current[0].getNeighbours() {
				if slices.Contains(checked, n) {
					continue
				}
				next = append(next, grid[n.Y][n.X])
			}

			checked = append(checked, current[0].pos)
			current = current[1:]
			if len(current) == 0 {
				if len(next) == 0 {
					break
				}
				stepsAway++
				current = next
				next = []tile{}
			}
		}

		fmt.Printf("Part 1: %d\n", stepsAway)
	}

	notLoop := []tile{}

	for y, l := range grid {
		for x, t := range l {
			if !slices.Contains(checked, t.pos) {
				grid[y][x].tType = ground
				notLoop = append(notLoop, grid[y][x])
			}
		}
	}

	// for each ground count how many | J L there are to the left
	// odd = in; even = out
	total := 0
	for _, t := range notLoop {
		toLeft := 0
		currentPos := t.pos

		for {
			currentPos = currentPos.Sub(image.Point{1, 0})
			if currentPos.X < 0 {
				break
			}
			currentTileType := grid[currentPos.Y][currentPos.X].tType
			if currentTileType == vertical || currentTileType == neBend || currentTileType == nwBend {
				toLeft++
			}
		}
		if toLeft%2 == 1 {
			total++
		}
	}
	fmt.Printf("Part 2: %d\n", total)
}

type tile struct {
	pos   image.Point
	tType tileType
}

func (t tile) getNeighbours() []image.Point {
	ret := []image.Point{}

	if t.tType == ground {
		fmt.Println("tried to get neighbours of a ground tile")
		return ret
	}
	if t.tType == startPos {
		panic("can't infer neighbours of start position")
	}

	if t.tType == vertical || t.tType == neBend || t.tType == nwBend {
		ret = append(ret, t.pos.Add(image.Pt(0, -1)))
	}
	if t.tType == vertical || t.tType == seBend || t.tType == swBend {
		ret = append(ret, t.pos.Add(image.Pt(0, 1)))
	}
	if t.tType == horizontal || t.tType == neBend || t.tType == seBend {
		ret = append(ret, t.pos.Add(image.Pt(1, 0)))
	}
	if t.tType == horizontal || t.tType == nwBend || t.tType == swBend {
		ret = append(ret, t.pos.Add(image.Pt(-1, 0)))
	}

	if len(ret) != 2 {
		fmt.Printf("not 2 neighbours: instead %d\n", len(ret))
	}

	return ret
}

func (t tile) getAdjacent() []image.Point {
	return []image.Point{
		t.pos.Add(image.Pt(0, -1)),
		t.pos.Add(image.Pt(0, 1)),
		t.pos.Add(image.Pt(-1, 0)),
		t.pos.Add(image.Pt(1, 0)),
	}
}

type tileType int

const (
	vertical tileType = iota
	horizontal
	neBend
	nwBend
	swBend
	seBend
	startPos
	ground
)

func typeFromRune(r rune) tileType {
	switch r {
	case '|':
		return vertical
	case '-':
		return horizontal
	case 'L':
		return neBend
	case 'J':
		return nwBend
	case '7':
		return swBend
	case 'F':
		return seBend
	case 'S':
		return startPos
	case '.':
		return ground

	default:
		panic(fmt.Errorf("unknown tiletype: %s", string(r)))
	}
}

func typeFromNeighbours(neighbours []image.Point) tileType {
	if len(neighbours) == 0 {
		return ground
	}
	if len(neighbours) != 2 {
		panic(fmt.Errorf("wrong number of neighbours: %d", len(neighbours)))
	}

	if slices.Contains(neighbours, image.Point{Y: -1}) {
		if slices.Contains(neighbours, image.Point{X: 1}) {
			return neBend
		}
		if slices.Contains(neighbours, image.Point{X: -1}) {
			return nwBend
		}
		return vertical
	}
	if slices.Contains(neighbours, image.Point{Y: 1}) {
		if slices.Contains(neighbours, image.Point{X: 1}) {
			return seBend
		}
		if slices.Contains(neighbours, image.Point{X: -1}) {
			return swBend
		}
		return horizontal
	}
	panic("couldn't determine tile from neighbours")
}
