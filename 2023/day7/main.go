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

	for _, line := range input {
		if line == "" {
			continue
		}

		inputParts := strings.Split(line, " ")
		rawHand := inputParts[0]
		handType := calcHand(rawHand)

		bid := helpers.MustAtoi(inputParts[1])

		hands = append(hands, hand{
			hType:  handType,
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
	fmt.Println(total)
}

var runeStrength string = "AKQJT98765432"

// is i stronger than j
func isStronger(i, j rune) int {
	iIndex := strings.IndexRune(runeStrength, i)
	jIndex := strings.IndexRune(runeStrength, j)

	return iIndex - jIndex
}

func calcHand(hand string) handType {
	handData := map[rune]int{}

	for _, c := range hand {
		amount, ok := handData[c]
		if !ok {
			handData[c] = 1
			continue
		}
		handData[c] = amount + 1
	}

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
