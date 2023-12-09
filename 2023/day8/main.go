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

	current := nodes["AAA"]

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
	fmt.Println(steps)
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
