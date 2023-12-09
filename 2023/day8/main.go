package main

import (
	"fmt"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")
	instructions := input[0]

	nodes := buildGraph(input[2:])

	fmt.Printf("Part 1: %d\n", part1(nodes, instructions))
	fmt.Printf("Part 2: %d\n", part2(nodes, instructions))
}

func part1(nodes map[string]*node, instructions string) int {
	current, ok := nodes["AAA"]
	if !ok {
		fmt.Print("Couldn't find node AAA! ")
		return 0
	}

	steps := 0

	for {
		for _, c := range instructions {
			current = current.travel(string(c))
			steps++
		}
		if current.label == "ZZZ" {
			break
		}
	}
	return steps
}

func part2(nodes map[string]*node, instructions string) int {
	startingNodes := []*node{}

	for _, n := range nodes {
		if n.label[2] == 'A' {
			startingNodes = append(startingNodes, n)
		}
	}

	allSteps := []int{}

	for _, n := range startingNodes {
		current := n
		steps := 0
		for {
			shouldBreak := false
			for _, c := range instructions {
				current = current.travel(string(c))
				steps++
				if current.label[2] == 'Z' {
					shouldBreak = true
					break
				}
			}
			if shouldBreak {
				break
			}
		}
		allSteps = append(allSteps, steps)
	}

	return helpers.LCM(allSteps...)
}

func buildGraph(input []string) map[string]*node {
	nodes := make(map[string]*node, len(input))
	getOrCreateNode := func(label string) *node {
		found, ok := nodes[label]
		if ok {
			return found
		}
		nodes[label] = &node{
			label: label,
		}
		return nodes[label]
	}

	for _, nodeString := range input {
		if nodeString == "" {
			continue
		}

		nodeDataString := strings.Split(nodeString, " = ")
		nodeLabel := nodeDataString[0]

		node := getOrCreateNode(nodeLabel)

		nodeChilderen := strings.Split(strings.TrimSuffix(strings.TrimPrefix(nodeDataString[1], "("), ")"), ", ")

		left := getOrCreateNode(nodeChilderen[0])
		right := getOrCreateNode(nodeChilderen[1])

		node.left = left
		node.right = right
	}

	return nodes
}

type node struct {
	label string
	left  *node
	right *node
}

func (n *node) String() string {
	return fmt.Sprintf("%s-> %s, %s", n.label, n.left.label, n.right.label)
}

func (n *node) travel(dir string) *node {
	switch dir {
	case "L":
		return n.left
	case "R":
		return n.right
	default:
		panic(fmt.Errorf("unknown direction of travel: %s", dir))
	}
}
