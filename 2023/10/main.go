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
	data2dArray, startPostion := convertDataTo2dArray(data)

	total, _ := calculateResult(data2dArray, startPostion)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	_, total = calculateResult(data2dArray, startPostion)
	fmt.Println("Result Part 2: " + fmt.Sprint(total))
}

func calculateResult(data [][]rune, startPosition [2]int) (int, int) {
	const POLYGON_CORNERS_CHARS = `LJ7FS`

	var polygonCorners [][2]int
	var currentChar rune

	currentPos := [2]int{startPosition[0] + 1, startPosition[1]} //start by going down
	lastPos := startPosition
	steps := 1
	polygonCorners = append(polygonCorners, [2]int{startPosition[0], startPosition[1]})

	for currentChar != 'S' {
		nextPos := getNextPosition(data, lastPos, currentPos)
		currentChar = data[nextPos[0]][nextPos[1]]
		lastPos, currentPos = currentPos, nextPos
		steps++
		if strings.Contains(POLYGON_CORNERS_CHARS, string(currentChar)) {
			polygonCorners = append(polygonCorners, [2]int{currentPos[0], currentPos[1]})
		}
	}

	area := shoelaceFormula(polygonCorners)

	return steps / 2, picksTheorem(area, steps)
}

func shoelaceFormula(polygonCorners [][2]int) (area int) {
	for i := 0; i < len(polygonCorners); i++ {
		if i > 0 {
			area -= polygonCorners[i][0] * polygonCorners[i-1][1]
		}
		if i < len(polygonCorners)-1 {
			area += polygonCorners[i][0] * polygonCorners[i+1][1]
		}

	}
	area = common.Abs(area) / 2
	return
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

func picksTheorem(area int, perimeterPoints int) int {
	return area + 1 - perimeterPoints/2
}
