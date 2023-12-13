package main

import (
	"common"
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

const NUMBER_INDEX = 1
const COLOR_INDEX = 2

func main() {
	common.InitSlogLogger(slog.LevelWarn)
	data := common.ReadTestFileContentInLines("input.txt")

	total := calculateResultForPart1(data)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	total = calculateResultForPart2(data)
	fmt.Println("Result Part 2: " + fmt.Sprint(total))
}

func calculateResultForPart1(data []string) int {

	var MAX_QUANTITIES = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	var gameLost bool
	total := 0

	for _, line := range data {
		slog.Info("Processing line", slog.String("line", line))

		splitLine := parseLine(line)

		gameNumber, _ := strconv.Atoi(regexp.MustCompile(`\d+`).FindString(splitLine[0]))
		gameLost = false

		for _, round := range splitLine[1:] {
			for _, coloredCubesString := range strings.Split(round, ",") {
				nbDrawnCubes, colorDrawnCubes := getCubeNumberAndColor(coloredCubesString)

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
	return total
}

func calculateResultForPart2(data []string) int {

	total := 0
	for i, line := range data {
		slog.Info("Processing line", slog.String("line", line))
		minCubes := make(map[string]int)

		splitLine := parseLine(line)
		for _, round := range splitLine[1:] {
			for _, coloredCubesString := range strings.Split(round, ",") {
				nbDrawnCubes, colorDrawnCubes := getCubeNumberAndColor(coloredCubesString)

				if minCubes[colorDrawnCubes] == 0 || nbDrawnCubes > minCubes[colorDrawnCubes] {
					minCubes[colorDrawnCubes] = nbDrawnCubes
				}

				slog.Debug("Minimum cubes info for color "+colorDrawnCubes,
					slog.Int("min cubes", minCubes[colorDrawnCubes]),
					slog.Int("nb cubes drawn", nbDrawnCubes))
			}
		}

		total += calculateGamePower(maps.Values(minCubes))
		slog.Debug(fmt.Sprintf("Total after Game %d : %d", i+1, total))
	}

	return total
}

func calculateGamePower(minCubes []int) int {
	total := 1
	for _, minCubeAmount := range minCubes {
		total *= minCubeAmount
	}
	return total
}

func parseLine(line string) []string {
	return strings.FieldsFunc(line, func(delimiter rune) bool {
		return delimiter == ':' || delimiter == ';'
	})
}

func getCubeNumberAndColor(coloredCubes string) (int, string) {
	matches := regexp.MustCompile(`^ (\d+) (\w*)$`).FindAllStringSubmatch(coloredCubes, -1)

	nbDrawnCubes, _ := strconv.Atoi(matches[0][NUMBER_INDEX])
	colorDrawnCubes := matches[0][COLOR_INDEX]
	return nbDrawnCubes, colorDrawnCubes
}
