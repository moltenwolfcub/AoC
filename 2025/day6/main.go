package main

import (
	"fmt"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

type Operator int

const (
	add Operator = iota
	mul
)

func (o Operator) do(a, b int) int {
	switch o {
	case add:
		return a + b
	case mul:
		return a * b
	default:
		panic("Unknown operator")
	}
}

type Equation struct {
	operands []int
	operator Operator
}

func main() {
	input := helpers.ReadLines("input.txt")

	equations := make([]Equation, len(strings.Fields(input[0])))

	for _, line := range input {
		if line == "" {
			continue
		}
		ops := strings.Fields(line)

		for i, op := range ops {
			if op == "+" {
				equations[i].operator = add
				continue
			}
			if op == "*" {
				equations[i].operator = mul
				continue
			}
			operand := helpers.MustAtoi(op)
			if equations[i].operands == nil {
				equations[i].operands = make([]int, 0)
			}
			equations[i].operands = append(equations[i].operands, operand)
		}
	}

	total := 0
	for _, eq := range equations {
		result := int(eq.operator)

		for _, op := range eq.operands {
			result = eq.operator.do(result, op)
		}

		total += result
	}

	fmt.Println(total)
}
