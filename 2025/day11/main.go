package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

type Node struct {
	label    string
	children []*Node
}

func main() {
	input := helpers.ReadLines("input.txt")

	network := make(map[string]*Node, 0)

	for _, line := range input {
		if line == "" {
			continue
		}
		parts := strings.Split(line, ":")
		thisLabel := parts[0]
		children := strings.Split(strings.Trim(parts[1], " "), " ")

		item, ok := network[thisLabel]
		if !ok {
			item = &Node{
				label:    thisLabel,
				children: make([]*Node, 0),
			}
			network[thisLabel] = item
		}
		for _, child := range children {
			childNode, ok := network[child]
			if !ok {
				childNode = &Node{
					label:    child,
					children: make([]*Node, 0),
				}
				network[child] = childNode
			}

			item.children = append(item.children, childNode)
		}
	}

	paths1 := Search1(network["you"], make([]*Node, 0), "out")
	fmt.Printf("Part 1: %v\n", paths1)

	fft_out := Search1(network["fft"], make([]*Node, 0), "out")
	dac_out := Search1(network["dac"], make([]*Node, 0), "out")
	dac_fft := Search1(network["dac"], make([]*Node, 0), "fft")
	fft_dac := Search1(network["fft"], make([]*Node, 0), "dac")

	fromKeyNodes := -1
	toKey := -1

	if dac_fft != 0 && fft_dac != 0 {
		panic("cyclic graph")
	}
	if dac_fft > 0 {
		fromKeyNodes = dac_fft * fft_out
		toKey = Search1(network["svr"], make([]*Node, 0), "dac")
	}
	if fft_dac > 0 {
		fromKeyNodes = fft_dac * dac_out
		toKey = Search1(network["svr"], make([]*Node, 0), "fft")
	}
	paths2 := toKey * fromKeyNodes

	fmt.Printf("Part 2: %v\n", paths2)
}

var search1Cache map[string]int = make(map[string]int)

func Search1(currentNode *Node, alreadyVisited []*Node, target string) int {
	if currentNode.label == target {
		return 1
	}

	total := 0

	newVisited := append(append(make([]*Node, 0), alreadyVisited...), currentNode)
	for _, child := range currentNode.children {
		if slices.Contains(alreadyVisited, child) {
			continue
		}

		key := child.label + target
		paths, ok := search1Cache[key]
		if !ok {
			paths = Search1(child, newVisited, target)
			search1Cache[key] = paths
		}

		total += paths
	}
	return total
}
