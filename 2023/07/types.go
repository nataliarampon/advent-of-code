package main

import (
	"slices"
	"sort"
	"strconv"

	"golang.org/x/exp/maps"
)

const (
	FiveOfAKind  = 6
	FourOfAKind  = 5
	FullHouse    = 4
	ThreeOfAKind = 3
	TwoPair      = 2
	OnePair      = 1
	None         = 0
)

var CHAR_TO_VALUE_MAP = map[byte]int{
	'A': 0, 'K': 1, 'Q': 2, 'T': 4, '9': 5, '8': 6,
	'7': 7, '6': 8, '5': 9, '4': 10, '3': 11, '2': 12, 'J': 13,
}

type Hand struct {
	content string
	typeOf  int
	bid     int
}

func NewHand(content string, bid string) *Hand {
	h := new(Hand)
	h.content = content
	h.bid, _ = strconv.Atoi(bid)
	h.typeOf = h.getType()
	return h
}

func (h Hand) getType() int {
	charactersMap := make(map[rune]int)
	for _, c := range h.content {
		charactersMap[c]++
	}

	repetitions := maps.Values(charactersMap)
	tryFullHouse := false
	tryTwoPair := false
	sort.Slice(repetitions, func(i, j int) bool { return repetitions[i] > repetitions[j] })
	for _, v := range repetitions {
		switch v {
		case 5:
			return FiveOfAKind
		case 4:
			return FourOfAKind
		case 3:
			tryFullHouse = true
		case 2:
			if tryFullHouse {
				return FullHouse
			}
			if tryTwoPair {
				return TwoPair
			}
			tryTwoPair = true
		case 1:
			if tryTwoPair {
				return OnePair
			}
			if tryFullHouse {
				return ThreeOfAKind
			}
		}
	}
	return None
}

func NewHandPart2(content string, bid string) *Hand {
	h := new(Hand)
	h.content = content
	h.bid, _ = strconv.Atoi(bid)
	h.typeOf = h.getTypePart2()
	return h
}

func (h Hand) getTypePart2() int {
	charactersMap := make(map[rune]int)
	for _, c := range h.content {
		charactersMap[c]++
	}

	repetitions := maps.Values(charactersMap)
	tryFullHouse := false
	tryTwoPair := false
	sort.Slice(repetitions, func(i, j int) bool { return repetitions[i] > repetitions[j] })
	for _, v := range repetitions {
		switch v {
		case 5:
			return FiveOfAKind
		case 4:
			if charactersMap['J'] == 1 || charactersMap['J'] == 4 {
				return FiveOfAKind
			}
			return FourOfAKind
		case 3:
			if charactersMap['J'] == 1 {
				return FourOfAKind
			}
			if charactersMap['J'] == 2 {
				return FiveOfAKind
			}
			tryFullHouse = true
		case 2:
			switch {
			case tryTwoPair && charactersMap['J'] == 2:
				if slices.Contains(repetitions[1:], 2) {
					return FourOfAKind
				}
			case tryFullHouse:
				if charactersMap['J'] == 3 {
					return FiveOfAKind
				}
				return FullHouse
			case tryTwoPair && charactersMap['J'] == 1:
				return FullHouse
			case tryTwoPair:
				return TwoPair
			}
			tryTwoPair = true
		case 1:
			switch {
			case tryTwoPair && charactersMap['J'] == 1:
				return ThreeOfAKind
			case tryTwoPair && charactersMap['J'] == 2:
				return ThreeOfAKind
			case charactersMap['J'] == 3:
				return FourOfAKind
			case charactersMap['J'] == 2:
				return FourOfAKind
			case charactersMap['J'] == 1:
				return OnePair
			case tryTwoPair:
				return OnePair
			case tryFullHouse:
				return ThreeOfAKind
			}
		}
	}
	return None
}

func (h Hand) compare(other Hand) bool {
	if h.typeOf == other.typeOf {
		for i := 0; i < len(h.content); i++ {
			if CHAR_TO_VALUE_MAP[h.content[i]] > CHAR_TO_VALUE_MAP[other.content[i]] {
				return true
			} else if CHAR_TO_VALUE_MAP[h.content[i]] < CHAR_TO_VALUE_MAP[other.content[i]] {
				return false
			}
		}
	}
	return h.typeOf < other.typeOf
}
