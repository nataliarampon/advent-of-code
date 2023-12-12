package main

import (
	"common"
	"fmt"
	"log/slog"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var total int
	common.InitSlogLogger(slog.LevelWarn)

	testData := common.ReadTestFileContent("input.txt")
	splitTestData := strings.Split(strings.ReplaceAll(testData, "\r\n", "\n"), "\n")

	for _, line := range splitTestData {
		slog.Info("Processing line", slog.String("line", line))
		total += countLineValue(line)
	}

	fmt.Println("Result: " + strconv.Itoa(total))
}

func countLineValue(line string) int {
	numbers := regexp.MustCompile(`^\D*(\d)?\w*(\d){1}\D*$`).FindAllStringSubmatch(line, -1)[0]
	if numbers[1] == "" {
		slog.Debug("Line with one single value detected", slog.String("value", numbers[1]))
		numbers[1] = numbers[2]
	}
	value, _ := strconv.Atoi(numbers[1] + numbers[2])
	slog.Debug("Calculated line value", slog.Int("value", value))
	return value
}
