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

	// total = calculateResultForPart2(data)
	// fmt.Println("Result Part 2: " + fmt.Sprint(total))

}

func calculateResultForPart1(data []string) float64 {
	total := 0.0

	for _, line := range data {
		slog.Info("Processing line", slog.String("line", line))

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

		if cardTotal > 0 {
			total += math.Pow(2, float64(cardTotal-1))
		}
	}

	return total
}

func calculateResultForPart2(data []string) float64 {
	panic("unimplemented")
}
