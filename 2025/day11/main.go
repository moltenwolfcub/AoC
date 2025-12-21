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

	var start *Node
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

		if thisLabel == "you" {
			if start != nil {
				panic("Found multiple yous")
			}
			start = item
		}
	}

	paths := Search(start, make([]*Node, 0))

	fmt.Println(paths)
}

func Search(currentNode *Node, alreadyVisited []*Node) int {
	if currentNode.label == "out" {
		return 1
	}

	total := 0

	newVisited := append(append(make([]*Node, len(alreadyVisited)), alreadyVisited...), currentNode)
	for _, child := range currentNode.children {
		if slices.Contains(alreadyVisited, child) {
			continue
		}

		paths := Search(child, newVisited)
		total += paths
	}
	return total
}
