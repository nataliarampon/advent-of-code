package main

import (
	"common"
	"fmt"
	"log/slog"
)

func main() {
	common.InitSlogLogger(slog.LevelInfo)
	data := common.ReadTestFileContentInLines("input.txt")
	data2dArray, startPostion := convertDataTo2dArray(data)

	total := calculateResultPart1(data2dArray, startPostion)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	total = calculateResultPart2(data2dArray, startPostion)
	fmt.Println("Result Part 2: " + fmt.Sprint(total))
}

func calculateResultPart1(data [][]rune, startPosition [2]int) int {
	currentPos := [2]int{startPosition[0] + 1, startPosition[1]} //start by going down
	lastPos := startPosition
	currentChar := '.'
	steps := 1

	for currentChar != 'S' {
		nextPos := getNextPosition(data, lastPos, currentPos)
		currentChar = data[nextPos[0]][nextPos[1]]
		lastPos, currentPos = currentPos, nextPos
		steps++
	}

	return steps / 2
}

func getNextPosition(data [][]rune, lastPos [2]int, currentPos [2]int) (nextPos [2]int) {
	i, j := currentPos[0], currentPos[1]
	iPast, jPast := lastPos[0], lastPos[1]
	currentChar := data[i][j]

	isGoingDown := iPast < i
	isGoingUp := iPast > i
	isGoingLeft := jPast > j

	switch currentChar {
	case '|':
		nextPos = common.TernaryOp(isGoingDown, [2]int{i + 1, j}, [2]int{i - 1, j})
	case '-':
		nextPos = common.TernaryOp(isGoingLeft, [2]int{i, j - 1}, [2]int{i, j + 1})
	case 'L':
		nextPos = common.TernaryOp(isGoingDown, [2]int{i, j + 1}, [2]int{i - 1, j})
	case 'J':
		nextPos = common.TernaryOp(isGoingDown, [2]int{i, j - 1}, [2]int{i - 1, j})
	case '7':
		nextPos = common.TernaryOp(isGoingUp, [2]int{i, j - 1}, [2]int{i + 1, j})
	case 'F':
		nextPos = common.TernaryOp(isGoingUp, [2]int{i, j + 1}, [2]int{i + 1, j})
	}
	return
}

func convertDataTo2dArray(data []string) (array2D [][]rune, startPosition [2]int) {
	LINES := len(data) + 2
	COLUMNS := len(data[0]) + 2

	array2D = common.Make2DArray[rune](LINES, COLUMNS)

	for i, line := range data {
		for j, char := range line {
			array2D[i+1][j+1] = char
			if char == 'S' {
				startPosition = [2]int{i + 1, j + 1}
			}
		}
	}
	return
}

func calculateResultPart2(data [][]rune, startPosition [2]int) int {
	return -1
}
