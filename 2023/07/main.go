package main

import (
	"common"
	"fmt"
	"log/slog"
	"sort"
	"strings"
)

func main() {
	common.InitSlogLogger(slog.LevelWarn)
	data := common.ReadTestFileContentInLines("input.txt")

	hands := readHands(data, 1)

	total := calculateResult(hands)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	hands = readHands(data, 2)

	total = calculateResult(hands)
	fmt.Println("Result Part 2: " + fmt.Sprint(total))
}

func readHands(data []string, version int) (hands []Hand) {
	for _, line := range data {
		parsedLine := strings.Fields(line)
		if version == 1 {
			hands = append(hands, *NewHand(parsedLine[0], parsedLine[1]))
		} else {
			hands = append(hands, *NewHandPart2(parsedLine[0], parsedLine[1]))
		}

	}
	return hands
}

func calculateResult(hands []Hand) (total int) {
	sort.Slice(hands, func(i, j int) bool { return hands[i].compare(hands[j]) })

	for i, h := range hands {
		total += (i + 1) * h.bid
	}
	return
}
