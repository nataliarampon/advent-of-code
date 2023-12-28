package main

import (
	"common"
	"fmt"
	"log/slog"
	"strings"
)

func main() {
	common.InitSlogLogger(slog.LevelInfo)
	data := common.ReadTestFileContentInLines("input.txt")

	total := calculateResult(data)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))
}

func calculateResult(data []string) (total int) {
	for _, line := range data {
		args := strings.Split(line, " ")
		nbSprings := common.ConvertStringArrayToInt(strings.Split(args[1], ","))
		total += getArrangements(args[0], nbSprings)
	}
	return
}

func getArrangements(condition string, springGroups []int) (total int) {
	if len(springGroups) == 0 || len(condition) == 0 {
		return 0
	}
	for i, char := range condition {
		if char == '#' {
			springGroups[0]--
		}
		if char == '?' {
			total += getArrangements(condition[i+1:], springGroups[1:])
			total += getArrangements(condition[i+1:], springGroups)
		}
	}

	return
}
