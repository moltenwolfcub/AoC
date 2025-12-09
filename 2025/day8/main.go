package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	_ "net/http/pprof"
	"os"
	"runtime/pprof"
	"sort"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

type JunctionBox struct {
	X, Y, Z int
	Circuit *Circuit
}

func (v JunctionBox) Dist(o JunctionBox) float64 {
	return math.Sqrt(float64((v.X-o.X)*(v.X-o.X) + (v.Y-o.Y)*(v.Y-o.Y) + (v.Z-o.Z)*(v.Z-o.Z)))
}

type Circuit struct {
	Boxes []*JunctionBox
}

func (c Circuit) Merge(o Circuit) *Circuit {
	newCircuit := &Circuit{
		Boxes: make([]*JunctionBox, len(c.Boxes)+len(o.Boxes)),
	}
	for i, box := range c.Boxes {
		box.Circuit = newCircuit
		newCircuit.Boxes[i] = box
	}
	for i, box := range o.Boxes {
		box.Circuit = newCircuit
		newCircuit.Boxes[len(c.Boxes)+i] = box
	}
	return newCircuit
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	input := helpers.ReadLines("input.txt")

	boxes := make([]*JunctionBox, len(input)-1)
	circuits := make([]*Circuit, len(input)-1)
	for i, line := range input {
		if line == "" {
			continue
		}

		c := &Circuit{}
		circuits[i] = c

		coords := strings.Split(line, ",")
		box := &JunctionBox{
			helpers.MustAtoi(coords[0]),
			helpers.MustAtoi(coords[1]),
			helpers.MustAtoi(coords[2]),
			c,
		}
		c.Boxes = append(c.Boxes, box)
		boxes[i] = box
	}

	fmt.Printf("Part 1: %v\n", part1(boxes, circuits))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(boxes []*JunctionBox, circuits []*Circuit) int {
	const maxConnection int = 1000
	var completedConnections map[[2]*JunctionBox]bool = make(map[[2]*JunctionBox]bool, 0)

	for z := 0; z < maxConnection; z++ {
		closest, empties, newConns := ShortestDist(boxes, completedConnections)
		completedConnections = newConns

		z += empties
		if z >= maxConnection {
			break
		}

		// remove previous 2 circuits
		toRemove := -1
		for i, c := range circuits {
			if c == closest[0].Circuit {
				toRemove = i
				break
			}
		}
		circuits = append(circuits[:toRemove], circuits[toRemove+1:]...)
		toRemove = -1
		for i, c := range circuits {
			if c == closest[1].Circuit {
				toRemove = i
				break
			}
		}
		circuits = append(circuits[:toRemove], circuits[toRemove+1:]...)

		// merge new circuits
		newCircuit := closest[0].Circuit.Merge(*closest[1].Circuit)
		circuits = append(circuits, newCircuit)
	}

	sort.SliceStable(circuits, func(i, j int) bool {
		return len(circuits[i].Boxes) > len(circuits[j].Boxes)
	})

	product := len(circuits[0].Boxes) * len(circuits[1].Boxes) * len(circuits[2].Boxes)

	return product
}

func part2(input []string) int {
	return 0
}

func ShortestDist(boxes []*JunctionBox, completedConnections map[[2]*JunctionBox]bool) ([2]*JunctionBox, int, map[[2]*JunctionBox]bool) {
	emptyTests := 0

loop:
	least := math.Inf(1)
	leastPair := [2]*JunctionBox{}
	for i := 0; i < len(boxes); i++ {
		first := boxes[i]

		for j := i + 1; j < len(boxes); j++ {
			second := boxes[j]

			if first.Circuit == second.Circuit && first.Circuit != nil {
				completed := [2]*JunctionBox{first, second}
				if _, ok := completedConnections[completed]; ok {
					continue
				}
			}

			d := first.Dist(*second)
			if d < least {
				least = d
				leastPair[0] = first
				leastPair[1] = second
			}
		}
	}
	if leastPair[0].Circuit == leastPair[1].Circuit && leastPair[0].Circuit != nil {
		completedConnections[leastPair] = true
		emptyTests++
		goto loop
	}
	completedConnections[leastPair] = true
	return leastPair, emptyTests, completedConnections
}
