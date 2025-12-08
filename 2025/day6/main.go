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

	fmt.Printf("Part 1: %v\n", part1(input))
	fmt.Printf("Part 2: %v\n", part2(input))
}

func part1(input []string) int {

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

	ans := calcEquations(equations)
	return ans
}

func part2(input []string) int {

	equations := make([]Equation, 0)
	rowLen := len(input[0])

	eq := Equation{}
	eq.operands = make([]int, 0)
	for i := 0; i < rowLen; i++ {

		value := 0
		filled := false
		for _, line := range input {
			if line == "" {
				continue
			}

			c := line[i]
			if c != ' ' {
				filled = true
			} else {
				continue
			}

			if c == '+' {
				eq.operator = add
				continue
			}
			if c == '*' {
				eq.operator = mul
				continue
			}

			d := int(c) - 48 // convert rune -> int
			value *= 10
			value += d
		}
		if filled {
			eq.operands = append(eq.operands, value)
		} else {
			equations = append(equations, eq)

			eq = Equation{}
			eq.operands = make([]int, 0)
		}
	}
	equations = append(equations, eq)

	ans := calcEquations(equations)
	return ans
}

func calcEquations(equations []Equation) int {
	total := 0
	for _, eq := range equations {
		result := int(eq.operator)

		for _, op := range eq.operands {
			result = eq.operator.do(result, op)
		}

		total += result
	}

	return total
}
