package main

import (
	"common"
	"fmt"
	"log/slog"
	"math"
	"regexp"
	"strings"
)

func main() {
	common.InitSlogLogger(slog.LevelWarn)
	data := common.ReadTestFileContentInLines("input.txt")

	total := calculateResultForPart1(data)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	totalPart2 := calculateResultForPart2(data)
	fmt.Println("Result Part 2: " + fmt.Sprint(totalPart2))

}

func calculateResultForPart1(data []string) float64 {
	total := 0.0

	for _, line := range data {
		slog.Info("Processing line", slog.String("line", line))

		cardTotal := calculateCardPoints(line)

		if cardTotal > 0 {
			total += math.Pow(2, float64(cardTotal-1))
		}
	}

	return total
}

func calculateResultForPart2(data []string) int {
	scratchcardCopies := make([]int, len(data))
	total := 0
	const ORIGINAL_SCRATCHPAD_COPY = 1

	for i, line := range data {
		slog.Info("Processing line", slog.String("line", line))

		cardTotal := calculateCardPoints(line)

		for j := 1; j <= cardTotal && i+j < len(data); j++ {
			if scratchcardCopies[i] > 0 {
				scratchcardCopies[i+j] += scratchcardCopies[i]
			}

			scratchcardCopies[i+j] += ORIGINAL_SCRATCHPAD_COPY
		}
	}

	for _, copies := range scratchcardCopies {
		total += copies
	}

	return total + len(data)
}

func calculateCardPoints(line string) int {
	cardTotal := 0

	splitStr := strings.Split(strings.Split(line, ":")[1], "|")
	myCard := strings.Fields(splitStr[0])
	winningCard := splitStr[1]

	for _, number := range myCard {
		hasNumberOnWinningCard := regexp.MustCompile(`(^|\s+)` + number + `(\s+|$)`).Match([]byte(winningCard))
		if hasNumberOnWinningCard {
			slog.Debug("Winning number detected", slog.String("number", number))
			cardTotal++
		}
	}
	return cardTotal
}
