package main

import (
	"common"
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

var MAX_QUANTITIES = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

const NUMBER_INDEX = 1
const COLOR_INDEX = 2

func main() {
	var gameLost bool

	common.InitSlogLogger(slog.LevelWarn)
	testData := common.ReadTestFileContentInLines("input.txt")
	total := 0

	for _, line := range testData {
		slog.Info("Processing line", slog.String("line", line))

		splitLine := parseLine(line)

		gameNumber, _ := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(splitLine[0]))
		gameLost = false

		for _, round := range splitLine[1:] {
			for _, coloredCubes := range strings.Split(round, ",") {
				matches := regexp.MustCompile(`^ (\d+) (\w*)$`).FindAllStringSubmatch(coloredCubes, -1)

				nbDrawnCubes, _ := strconv.Atoi(matches[0][NUMBER_INDEX])
				colorDrawnCubes := matches[0][COLOR_INDEX]

				if nbDrawnCubes > MAX_QUANTITIES[colorDrawnCubes] {
					gameLost = true
				}

				slog.Debug("Processing drawn colored cubes",
					slog.Int("number of cubes", nbDrawnCubes),
					slog.String("drawn cube color", colorDrawnCubes))
			}
		}

		if gameLost == false {
			slog.Info("Game was won ", slog.Int("game number", gameNumber))
			total += gameNumber
		}
	}

	fmt.Println("Result: " + fmt.Sprint(total))
}

func parseLine(line string) []string {
	return strings.FieldsFunc(line, func(delimiter rune) bool {
		return delimiter == ':' || delimiter == ';'
	})
}
