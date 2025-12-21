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

	paths1 := Search1(network["you"], make([]*Node, 0))
	fmt.Printf("Part 1: %v\n", paths1)

	paths2 := Search2(network["svr"], make([]*Node, 0), network)
	fmt.Printf("Part 2: %v\n", paths2)
}

func Search1(currentNode *Node, alreadyVisited []*Node) int {
	if currentNode.label == "out" {
		return 1
	}

	total := 0

	newVisited := append(append(make([]*Node, 0), alreadyVisited...), currentNode)
	for _, child := range currentNode.children {
		if slices.Contains(alreadyVisited, child) {
			continue
		}

		paths := Search1(child, newVisited)
		total += paths
	}
	return total
}

func Search2(currentNode *Node, alreadyVisited []*Node, network map[string]*Node) int {
	PrintVisited(alreadyVisited)

	if currentNode.label == "out" {
		if slices.Contains(alreadyVisited, network["dac"]) && slices.Contains(alreadyVisited, network["fft"]) {
			return 1
		}
		return 0
	}

	total := 0

	newVisited := append(append(make([]*Node, 0), alreadyVisited...), currentNode)
	for _, child := range currentNode.children {
		if slices.Contains(alreadyVisited, child) {
			continue
		}

		paths := Search2(child, newVisited, network)
		total += paths
	}
	return total
}

func PrintVisited(visited []*Node) {
	str := ""

	for _, v := range visited {
		str += v.label
		str += " "
	}

	fmt.Println(str)
}
