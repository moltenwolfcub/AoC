package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/moltenwolfcub/AoC/helpers"
)

func main() {
	input := helpers.ReadLines("input.txt")

	hands := []hand{}
	hands2 := []hand{}

	for _, line := range input {
		if line == "" {
			continue
		}

		inputParts := strings.Split(line, " ")
		rawHand := inputParts[0]
		handType := calcHand(rawHand, false)
		handType2 := calcHand(rawHand, true)

		bid := helpers.MustAtoi(inputParts[1])

		hands = append(hands, hand{
			hType:  handType,
			bid:    bid,
			rawVal: rawHand,
		})
		hands2 = append(hands2, hand{
			hType:  handType2,
			bid:    bid,
			rawVal: rawHand,
		})
	}

	//sort from highest-lowest
	sort.Slice(hands, func(i, j int) bool {
		iHand := hands[i]
		jHand := hands[j]

		if iHand.hType < jHand.hType {
			return false
		}
		if iHand.hType > jHand.hType {
			return true
		}

		for i := 0; i < len(iHand.rawVal); i++ {
			ci := rune(iHand.rawVal[i])
			cj := rune(jHand.rawVal[i])
			comp := isStronger(ci, cj)

			if comp < 0 {
				return false
			}
			if comp > 0 {
				return true
			}
		}
		panic("Same string")
	})

	total := 0
	for i, h := range hands {
		total += (i + 1) * h.bid
	}
	fmt.Printf("Part 1: %d\n", total)

	//Part 2

	//sort from highest-lowest
	sort.Slice(hands2, func(i, j int) bool {
		iHand := hands2[i]
		jHand := hands2[j]

		if iHand.hType < jHand.hType {
			return false
		}
		if iHand.hType > jHand.hType {
			return true
		}

		for i := 0; i < len(iHand.rawVal); i++ {
			ci := rune(iHand.rawVal[i])
			cj := rune(jHand.rawVal[i])
			comp := isStronger2(ci, cj)

			if comp < 0 {
				return false
			}
			if comp > 0 {
				return true
			}
		}
		panic("Same string")
	})

	total = 0
	for i, h := range hands2 {
		total += (i + 1) * h.bid
	}
	fmt.Printf("Part 2: %d\n", total)
}

var runeStrength string = "AKQJT98765432"
var runeStrength2 string = "AKQT98765432J"

// is i stronger than j
func isStronger(i, j rune) int {
	iIndex := strings.IndexRune(runeStrength, i)
	jIndex := strings.IndexRune(runeStrength, j)

	return iIndex - jIndex
}

// is i stronger than j
func isStronger2(i, j rune) int {
	iIndex := strings.IndexRune(runeStrength2, i)
	jIndex := strings.IndexRune(runeStrength2, j)

	return iIndex - jIndex
}

func calcHand(hand string, jokers bool) handType {
	handData := map[rune]int{}
	jokerCount := 0

	for _, c := range hand {
		if c == 'J' && jokers {
			jokerCount++
			continue
		}

		amount, ok := handData[c]
		if !ok {
			handData[c] = 1
			continue
		}
		handData[c] = amount + 1
	}

	maximumNumber := 0
	maxKey := rune(0)
	for k, v := range handData {
		if v > maximumNumber {
			maxKey = k
			maximumNumber = v
		}
	}

	handData[maxKey] += jokerCount

	found3 := false
	found2 := false
	for _, i := range handData {
		switch i {
		case 5:
			return fiveOfAKind
		case 4:
			return fourOfAKind
		case 3:
			found3 = true
			if found2 {
				return fullHouse
			}
		case 2:
			if found2 {
				return TwoPair
			}

			found2 = true
			if found3 {
				return fullHouse
			}
		}
	}
	if found3 {
		return threeOfAKind
	}
	if found2 {
		return OnePair
	}
	return HighCard
}

type handType int

const (
	fiveOfAKind handType = iota
	fourOfAKind
	fullHouse
	threeOfAKind
	TwoPair
	OnePair
	HighCard
)

type hand struct {
	hType  handType
	bid    int
	rawVal string
}
