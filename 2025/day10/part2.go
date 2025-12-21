package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

type Equation []int

func (e Equation) Scale(sf int) {
	for i, v := range e {
		e[i] = v * sf
	}
}

func (e Equation) GetScaled(sf int) Equation {
	eq := make(Equation, len(e))
	for i, v := range e {
		eq[i] = v * sf
	}
	return eq
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

func (es *EquationSystem) AddEquationRows(a, b Equation, dest int) {
	sum := make(Equation, len(a))
	for i := range sum {
		sum[i] = a[i] + b[i]
	}

	es.equations[dest] = sum
}

var joltageRe = regexp.MustCompile(`{((?:\d+,?)+)}`)

func part2(input []string) int {
	runningTotal := 0

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
		ReduceSystem(es)

		constraints := make([]int, es.numButtons)
		for i := range constraints {
			constraints[i] = math.MaxInt
		}

		for button := 0; button < es.numButtons; button++ {
			for _, connection := range buttons[button] {
				constraints[button] = min(constraints[button], joltages[connection])
			}
		}

		sol := make([]int, es.numButtons)
		for i := range sol {
			sol[i] = -1
		}
		minPresses := math.MaxInt
		SolveSystemRecursive(buttons, joltages, *es, min(len(sol)-1, len(es.equations)-1), len(sol)-1, constraints, sol, &minPresses)

		if minPresses >= math.MaxInt {
			panic("AHHHHHHH")
		}

		runningTotal += minPresses
	}

	return runningTotal
}

func ReduceSystem(es *EquationSystem) {
	diagonalLen := min(es.numJoltages, es.numButtons)

	for diagonal := 0; diagonal < diagonalLen; diagonal++ {
		rowWithValue := -1
		// make diagonal non-zero
		if es.equations[diagonal][diagonal] == 0 {
			for i := diagonal + 1; i < es.numJoltages; i++ {
				if es.equations[i][diagonal] != 0 {
					rowWithValue = i
					break
				}
			}
			if rowWithValue == -1 {
				continue
				// if it and every value below is zero then this column is already reduced
			}
			es.SwapRows(diagonal, rowWithValue)
		} else {
			rowWithValue = diagonal
		}
		if es.equations[diagonal][diagonal] < 0 {
			es.ScaleRow(diagonal, -1)
		}

		// make every other value zero
		for i := rowWithValue + 1; i < es.numJoltages; i++ {
			currentVal := es.equations[i][diagonal]
			if currentVal == 0 {
				continue
			}

			lcm := helpers.LCM(currentVal, es.equations[diagonal][diagonal])
			es.ScaleRow(i, lcm/currentVal)

			sf := -1 * (lcm / es.equations[diagonal][diagonal])
			scaledEq := es.equations[diagonal].GetScaled(sf)
			es.AddEquationRows(es.equations[i], scaledEq, i)
		}
	}
}

func SolveSystem(es *EquationSystem) []int {
	solution := make([]int, es.numButtons)

	diagonalLen := min(es.numJoltages, es.numButtons)

	for diagonal := diagonalLen - 1; diagonal >= 0; diagonal-- {
		dVal := es.equations[diagonal][diagonal]
		if dVal == 0 {
			solution[diagonal] = 0
			continue
		}

		c := es.equations[diagonal][es.numButtons]

		if float64(c)/float64(dVal) != float64(c/dVal) {
			panic("Doesn't yield integer solution")
		}

		presses := c / dVal
		solution[diagonal] = presses

		for row := diagonal - 1; row >= 0; row-- {
			coefficient := es.equations[row][diagonal]

			v := coefficient * presses
			es.equations[row][es.numButtons] -= v
		}
	}

	return solution
}

func SolveSystemRecursive(buttons [][]int, joltages []int, es EquationSystem, rowToSolve int, nextUnknown int, constraints []int, partialSolution []int, minPresses *int) {
	if rowToSolve == -1 {
		thisJolts := make([]int, len(joltages))
		for button := range buttons {
			for _, connection := range buttons[button] {
				thisJolts[connection] += partialSolution[button]
			}
		}

		if slices.Equal(thisJolts, joltages) {
			sum := 0
			for _, v := range partialSolution {
				sum += v
			}

			*minPresses = min(*minPresses, sum)
		}
		return
	}

	if nextUnknown > rowToSolve {
		for guess := 0; guess <= constraints[nextUnknown]; guess++ {
			partialSolution[nextUnknown] = guess
			SolveSystemRecursive(buttons, joltages, es, rowToSolve, nextUnknown-1, constraints, partialSolution, minPresses)
		}
		return
	}

	if rowToSolve != nextUnknown {
		panic("assert false")
	}
	if es.equations[rowToSolve][nextUnknown] == 0 {
		for guess := 0; guess <= constraints[nextUnknown]; guess++ {
			partialSolution[nextUnknown] = guess
			SolveSystemRecursive(buttons, joltages, es, rowToSolve-1, nextUnknown-1, constraints, partialSolution, minPresses)
		}
		return
	}

	rowTargetSum := es.equations[rowToSolve][es.numButtons]
	for known := nextUnknown + 1; known < len(partialSolution); known++ {
		rowTargetSum -= es.equations[rowToSolve][known] * partialSolution[known]
	}

	if rowTargetSum%es.equations[rowToSolve][rowToSolve] != 0 {
		return //non integer solution
	}
	possibleSol := rowTargetSum / es.equations[rowToSolve][rowToSolve]
	if possibleSol < 0 {
		return
	}
	partialSolution[nextUnknown] = possibleSol
	SolveSystemRecursive(buttons, joltages, es, rowToSolve-1, nextUnknown-1, constraints, partialSolution, minPresses)
}
