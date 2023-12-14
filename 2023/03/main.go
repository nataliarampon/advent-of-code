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

	total = calculateResultForPart2(data)
	fmt.Println("Result Part 2: " + fmt.Sprint(total))
}

func calculateResultForPart1(data []string) int {
	total := 0

	for i := 0; i < len(data); i++ {
		slog.Info("Processing line", slog.String("line", data[i]))

		regexNumbers := regexp.MustCompile(`\d+`)

		numbers := regexNumbers.FindAllStringSubmatch(data[i], -1)
		numbersIndex := regexNumbers.FindAllStringIndex(data[i], -1)

		for j, nb := range numbers {
			isPartNumber, _ := findSurroundingMatches(numbersIndex[j], data, i, `[^\d|\.]`)
			if isPartNumber {
				nbInt, _ := strconv.Atoi(nb[0])
				slog.Info("Part number found", slog.Int("number", nbInt))
				total += nbInt
			}
		}
	}

	return total
}

func findSurroundingMatches(numberIndexRange []int, data []string, i int, regexString string) (bool, []string) {
	var specialCharIndexes []string
	const INDEX_FORMAT = "%d-%d"

	regexSpecialChars := regexp.MustCompile(regexString)
	upperBound := i - 1
	lowerBound := i + 1
	leftBound := numberIndexRange[0] - 1
	rightBound := numberIndexRange[1] + 1
	specialCharFound := false

	if leftBound >= 0 {
		specialCharFound = regexSpecialChars.Match([]byte{data[i][leftBound]})
		if specialCharFound { specialCharIndexes = append(specialCharIndexes, fmt.Sprintf(INDEX_FORMAT, i, leftBound)) }
	} else {
		leftBound++
	}

	if rightBound <= len(data[i]) - 1 {
		specialCharFound = regexSpecialChars.Match([]byte{data[i][rightBound-1]}) || specialCharFound
		if specialCharFound { specialCharIndexes = append(specialCharIndexes, fmt.Sprintf(INDEX_FORMAT, i, rightBound-1)) }
	} else {
		rightBound--
	}

	if !specialCharFound && upperBound >= 0 {
		specialCharFound = regexSpecialChars.Match([]byte(data[upperBound][leftBound:rightBound]))
		for _, index := range regexSpecialChars.FindAllIndex([]byte(data[upperBound][leftBound:rightBound]),-1) {
			specialCharIndexes = append(specialCharIndexes, fmt.Sprintf(INDEX_FORMAT, upperBound, leftBound + index[0]))
		}
	}
	if !specialCharFound && lowerBound <= len(data) - 1 {
		specialCharFound = regexSpecialChars.Match([]byte(data[lowerBound][leftBound:rightBound]))
		for _, index := range regexSpecialChars.FindAllIndex([]byte(data[lowerBound][leftBound:rightBound]),-1) {
			specialCharIndexes = append(specialCharIndexes, fmt.Sprintf(INDEX_FORMAT, lowerBound, leftBound + index[0]))
		}
	}

	return specialCharFound, specialCharIndexes
}

func calculateResultForPart2(data []string) int {
	total := 0
	gearIndexes := make(map[string][]int)

	for i := 0; i < len(data); i++ {
		slog.Info("Processing line", slog.String("line", data[i]))

		regexNumbers := regexp.MustCompile(`\d+`)

		numbers := regexNumbers.FindAllStringSubmatch(data[i], -1)
		numbersIndex := regexNumbers.FindAllStringIndex(data[i], -1)

		for j, nb := range numbers {
			 _, indexes := findSurroundingMatches(numbersIndex[j], data, i, `\*`)
			for _, gearIndex := range indexes {
				nbInt, _ := strconv.Atoi(nb[0])
				slog.Info("Gear part found", slog.Int("number", nbInt), slog.String("gear index", gearIndex))
				gearIndexes[gearIndex] = append(gearIndexes[gearIndex], nbInt)
			}
		}
	}

	for i, gearParts := range gearIndexes {
		if len(gearParts) == 2 {
			gearRatio := gearParts[0]*gearParts[1]
			total += gearRatio
			slog.Info("Gear found", slog.String("index", i), slog.Int("gear ratio", gearRatio))
		}
	}

	return total
}
