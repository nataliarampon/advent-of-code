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

	dataInt := make([][]int, len(data))
	for i, line := range data {
		dataInt[i] = common.ConvertStringArrayToInt(strings.Fields(line))
	}

	total := calculateResult(dataInt, func(current, next int) int { return next + current })
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	for i, array := range dataInt {
		dataInt[i] = common.ReverseArray[int](array)
	}
	total = -calculateResult(dataInt, func(current, next int) int { return next - current })
	fmt.Println("Result Part 2: " + fmt.Sprint(total))
}

func calculateResult(data [][]int, getNext func(int, int) int) (total int) {

	var differenceArrays [][]int

	for _, valueHistory := range data {
		differenceArrays = [][]int{valueHistory}
		diff := valueHistory
		for isNonZero(diff) {
			diff = calculateDiffArray(diff)
			differenceArrays = append(differenceArrays, diff)
		}

		next := 0
		for i := len(differenceArrays) - 1; i >= 1; i-- {
			next = getNext(differenceArrays[i-1][len(differenceArrays[i-1])-1], next)
		}
		total += next
	}

	return
}

func calculateDiffArray(history []int) (diff []int) {
	diff = make([]int, len(history)-1)
	for i := 1; i < len(history); i++ {
		diff[i-1] = history[i] - history[i-1]
	}
	return diff
}

func isNonZero(array []int) bool {
	for _, v := range array {
		if v != 0 {
			return true
		}
	}
	return false
}
