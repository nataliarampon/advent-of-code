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

	hands := readHands(data)

	total := calculateResultForPart1(hands)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	total = calculateResultForPart2(hands)
	fmt.Println("Result Part 2: " + fmt.Sprint(total))
}

func readHands(data []string) (hands []Hand) {
	for _, line := range data {
		parsedLine := strings.Fields(line)
		hands = append(hands, *NewHand(parsedLine[0], parsedLine[1]))
	}
	return hands
}

func calculateResultForPart1(hands []Hand) (total int) {
	sort.Slice(hands, func(i, j int) bool { return hands[i].compare(hands[j]) })

	for i, h := range hands {
		total += (i + 1) * h.bid
	}
	return
}

func calculateResultForPart2(hands []Hand) int {
	return 0
}
