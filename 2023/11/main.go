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

	total := calculateResult(data, 1)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	total = calculateResult(data, 1000000-1)
	fmt.Println("Result Part 2: " + fmt.Sprint(total))
}

func calculateResult(data []string, expansionFactor int) (total int) {

	data = expandUniverse(data)
	galaxies := getGalaxies(data)

	for i, currGalaxy := range galaxies {
		for _, pairGalaxy := range galaxies[i+1:] {
			total += getShortestPathBetweenGalaxies(data, currGalaxy, pairGalaxy, expansionFactor)
		}
	}

	return total
}

func expandUniverse(data []string) (universe []string) {

	universe = expandVertically(data)

	expandHorizontally(universe)

	return
}

func expandVertically(data []string) (universe []string) {
	const EXPANSION = "X"
	for _, line := range data {
		universe = append(universe, line)
		if !strings.ContainsRune(line, '#') {
			universe = append(universe, strings.Repeat(EXPANSION, len(line)))
		}
	}
	return universe
}

func expandHorizontally(universe []string) {
	const EXPANSION = "X"
	for i := 0; i < len(universe[0]); i++ {
		hasGalaxy := false
		for _, line := range universe {
			if line[i] == '#' {
				hasGalaxy = true
				break
			}
		}
		if !hasGalaxy {
			for j, line := range universe {
				universe[j] = line[:i] + EXPANSION + line[i:]
			}
			i++
		}
	}
}

func getGalaxies(data []string) [][2]int {
	var galaxies [][2]int
	for i, line := range data {
		for j, char := range line {
			if char == '#' {
				galaxies = append(galaxies, [2]int{i, j})
			}
		}
	}
	return galaxies
}

func getShortestPathBetweenGalaxies(universe []string, currGalaxy [2]int, pairGalaxy [2]int, expansionFactor int) (path int) {
	startY, endY := min(currGalaxy[0], pairGalaxy[0]), max(currGalaxy[0], pairGalaxy[0])
	startX, endX := min(currGalaxy[1], pairGalaxy[1]), max(currGalaxy[1], pairGalaxy[1])

	for i := startY; i < endY; i++ {
		if universe[i][startX] == 'X' {
			path += expansionFactor
		} else {
			path++
		}
	}

	for j := startX; j < endX; j++ {
		if universe[startY][j] == 'X' {
			path += expansionFactor
		} else {
			path++
		}
	}
	return
}
