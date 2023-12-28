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

	data = expandUniverse(data)
	galaxies := getGalaxies(data)

	for i, currGalaxy := range galaxies {
		for _, pairGalaxy := range galaxies[i+1:] {
			total += getShortestPathBetweenGalaxies(currGalaxy, pairGalaxy)
		}
	}

	return total
}

func getShortestPathBetweenGalaxies(currGalaxy [2]int, pairGalaxy [2]int) (path int) {
	path += common.Abs(currGalaxy[0] - pairGalaxy[0])
	path += common.Abs(currGalaxy[1] - pairGalaxy[1])
	return
}

func expandUniverse(data []string) (universe []string) {

	universe = expandVertically(data)

	expandHorizontally(universe)

	return
}

func expandVertically(data []string) (universe []string) {
	for _, line := range data {
		universe = append(universe, line)
		if !strings.ContainsRune(line, '#') {
			universe = append(universe, line)
		}
	}
	return universe
}

func expandHorizontally(universe []string) {
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
				universe[j] = line[:i] + "." + line[i:]
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
