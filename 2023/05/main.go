package main

import (
	"common"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	common.InitSlogLogger(slog.LevelWarn)
	data := common.ReadTestFileContentInLines("input.txt")

	total := calculateResultForPart1(data)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	// totalPart2 := calculateResultForPart2(data)
	// fmt.Println("Result Part 2: " + fmt.Sprint(totalPart2))

}

func calculateResultForPart1(data []string) int {
	const ORIGIN = 0
	const DESTINATION = 1
	const RANGE = 2

	var seedToNextValueMap map[int]int
	var stepRange [][3]int // [nb mappings][range][originStart,destinationStart]
	var destinationStart, originStart, rangeValue, nbLinesInMapping int

	seeds := strings.Split(data[0], " ")[1:]
	seedToNextValueMap = initSeedMap(seeds)

	processingMap := false
	for _, line := range data[1:] {
		if processingMap && line != "" {
			fmt.Sscanf(line, "%d %d %d", &destinationStart, &originStart, &rangeValue)
			stepRange = append(stepRange, [3]int{originStart, destinationStart, rangeValue})
			nbLinesInMapping++
		} else {
			for seed, current := range seedToNextValueMap {
				for i := 0; i < nbLinesInMapping; i++ {
					if current >= stepRange[i][ORIGIN] && current <= stepRange[i][ORIGIN]+stepRange[i][RANGE]-1 {
						destValue := current + (stepRange[i][DESTINATION] - stepRange[i][ORIGIN])
						seedToNextValueMap[seed] = destValue
					}
				}
			}
			nbLinesInMapping = 0
			stepRange = nil
		}

		processingMap = getProcessingStatus(line)
	}

	var minLocation int
	for i, location := range maps.Values(seedToNextValueMap) {
		if i == 0 {
			minLocation = location
		}
		if location < minLocation {
			minLocation = location
		}
	}

	return minLocation
}

func initSeedMap(seeds []string) map[int]int {
	seedToNextValueMap := make(map[int]int)
	for _, seed := range seeds {
		seedInt, _ := strconv.Atoi(seed)
		seedToNextValueMap[seedInt] = seedInt
	}
	return seedToNextValueMap
}

// using Go's naked return
func getProcessingStatus(line string) (processingMap bool) {
	if line == "" {
		processingMap = false
	} else {
		processingMap = true
	}
	return
}

func calculateResultForPart2(data []string) int {
	panic("unimplemented")
}
