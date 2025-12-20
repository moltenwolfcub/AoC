package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

type Equation []int

func (e Equation) Scale(sf int) {
	for i, v := range e {
		e[i] = v * sf
	}
}

type EquationSystem struct {
	numJoltages int //how many equations we'll have
	numButtons  int //how many variables will be in the equation
	equations   []Equation
}

func NewEquationSystem(buttons [][]int, joltages []int) *EquationSystem {
	es := EquationSystem{}
	es.numButtons = len(buttons)
	es.numJoltages = len(joltages)
	es.equations = make([]Equation, es.numJoltages)

	for i := range es.equations {
		es.equations[i] = make(Equation, es.numButtons+1)
		es.equations[i][es.numButtons] = joltages[i]
	}

	for i, b := range buttons {
		// i is the button index (2,3):0 (1,3):1 etc its the index into the equation
		// the individual values are the index into the equation system
		for _, v := range b {
			es.equations[v][i] = 1
		}
	}

	return &es
}

func (es EquationSystem) String() string {
	str := ""

	str += fmt.Sprintf("#Equations (buttons): %d\n", es.numButtons)
	str += fmt.Sprintf("#Variables (joltages): %d\n", es.numJoltages)
	for _, eq := range es.equations {
		str += "| "
		for i := 0; i < len(eq)-1; i++ {
			str += fmt.Sprintf("%3d ", eq[i])
		}
		str += fmt.Sprintf(" | %3d |\n", eq[len(eq)-1])
	}

	return str
}

func (es *EquationSystem) SwapRows(first, second int) {
	es.equations[first], es.equations[second] = es.equations[second], es.equations[first]
}

func (es *EquationSystem) ScaleRow(row, sf int) {
	es.equations[row].Scale(sf)
}

func (es *EquationSystem) AddRows(a, b, dest int) {
	rowA := es.equations[a]
	rowB := es.equations[b]

	sum := make(Equation, len(rowA))
	for i := range sum {
		sum[i] = rowA[i] + rowB[i]
	}

	es.equations[dest] = sum
}

var joltageRe = regexp.MustCompile(`{((?:\d+,?)+)}`)

func part2(input []string) int {

	for _, line := range input {
		if line == "" {
			continue
		}

		inputButtons := buttonRe.FindAllStringSubmatch(line, -1)
		inputJoltage := joltageRe.FindStringSubmatch(line)

		buttons := make([][]int, len(inputButtons))
		for i, b := range inputButtons {
			strJoltageIndicies := strings.Split(b[1], ",")
			joltageIndicies := make([]int, len(strJoltageIndicies))
			for j, strJoltageIndex := range strJoltageIndicies {
				joltageIndicies[j] = helpers.MustAtoi(strJoltageIndex)
			}

			buttons[i] = joltageIndicies
		}

		strJoltages := strings.Split(inputJoltage[1], ",")
		joltages := make([]int, len(strJoltages))
		for i, j := range strJoltages {
			joltages[i] = helpers.MustAtoi(j)
		}

		es := NewEquationSystem(buttons, joltages)
		fmt.Println(es)
		es.SwapRows(0, 3)
		es.SwapRows(0, 2)
		es.ScaleRow(2, -1)
		es.AddRows(0, 2, 2)
		es.AddRows(1, 2, 2)
		fmt.Println(es)
	}

	return 0
}
