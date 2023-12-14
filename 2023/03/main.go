package main

import (
	"common"
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
)

func main() {
	common.InitSlogLogger(slog.LevelWarn)
	data := common.ReadTestFileContentInLines("input.txt")

	total := calculateResultForPart1(data)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	// total = calculateResultForPart2(data)
	// fmt.Println("Result Part 2: " + fmt.Sprint(total))
}

func calculateResultForPart1(data []string) int {
	total := 0

	for i := 0; i < len(data); i++ {
		slog.Info("Processing line", slog.String("line", data[i]))

		regexNumbers := regexp.MustCompile(`\d+`)

		numbers := regexNumbers.FindAllStringSubmatch(data[i], -1)
		numbersIndex := regexNumbers.FindAllStringIndex(data[i], -1)

		for j, nb := range numbers {
			if isPartNumber(numbersIndex[j], data, i) {
				nbInt, _ := strconv.Atoi(nb[0])
				slog.Info("Part number found", slog.Int("number", nbInt))
				total += nbInt
			}
		}
	}

	return total
}

func isPartNumber(numberIndexRange []int, data []string, i int) bool {
	specialCharFound := false

	regexSpecialChars := regexp.MustCompile(`[^\d|\.]`)
	upperBound := i - 1
	lowerBound := i + 1
	leftBound := numberIndexRange[0] - 1
	rightBound := numberIndexRange[1] + 1

	if leftBound >= 0 {
		specialCharFound = regexSpecialChars.Match([]byte{data[i][leftBound]})
	} else {
		leftBound++
	}

	if rightBound <= len(data[i]) - 1 {
		specialCharFound = regexSpecialChars.Match([]byte{data[i][rightBound-1]}) || specialCharFound
	} else {
		rightBound--
	}

	if !specialCharFound && upperBound >= 0 {
		specialCharFound = regexSpecialChars.Match([]byte(data[upperBound][leftBound:rightBound]))
	}
	if !specialCharFound && lowerBound <= len(data) - 1 {
		specialCharFound = regexSpecialChars.Match([]byte(data[lowerBound][leftBound:rightBound]))
	}

	return specialCharFound
}

func calculateResultForPart2(data []string) int {
	panic("unimplemented")
}
