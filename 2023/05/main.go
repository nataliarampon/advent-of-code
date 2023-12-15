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
	data := common.ReadTestFileContentInLines("test.txt")

	total := calculateResultForPart1(data)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	totalPart2 := calculateResultForPart2(data)
	fmt.Println("Result Part 2: " + fmt.Sprint(totalPart2))

}

func calculateResultForPart1(data []string) int {
	var seedToNextValueMap map[int]int
	var stepRange []SeedMapping
	var nbLinesInMapping int

	seeds := strings.Split(data[0], " ")[1:]
	seedToNextValueMap = initSeedMap(seeds)

	processingMap := false
	for _, line := range data[1:] {
		if processingMap && line != "" {
			stepRange = calculateStepRange(line, stepRange)
			nbLinesInMapping++
		} else {
			for seed, current := range seedToNextValueMap {
				for i := 0; i < nbLinesInMapping; i++ {
					if current >= stepRange[i].start && current <= stepRange[i].end {
						destValue := current + (stepRange[i].destinationStart - stepRange[i].start)
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

func calculateResultForPart2(data []string) int {
	var stepRange []SeedMapping

	seeds := strings.Split(data[0], " ")[1:]
	nextValueList := initSeedMapPart2(seeds)

	processingMap := false
	for _, line := range data[1:] {
		if processingMap && line != "" {
			stepRange = calculateStepRange(line, stepRange)
		} else {
			for _, current := range nextValueList {
				for i := 0; i < len(stepRange); i++ {
					// three cases
					// dest range within seed range
					// 		seed range 1 - 4
					// 		dest range  2 - 3 

					// dest range within left part of seed range
					// 		seed range 6 - 10
					// 		dest range  2 - 8 

					// dest range within right part of seed range
					// 		seed range 6 - 10
					// 		dest range  8 - 12


					// if current >= stepRange[i].start && current <= stepRange[i].start+stepRange[i].rangeValue-1 {
					// 	destValue := current + (stepRange[i].destinationStart - stepRange[i].start)
					// 	nextValueList[seed] = destValue
					}
				}
			}
			stepRange = nil
		}

		processingMap = getProcessingStatus(line)
	}

	var minLocation int
	// for i, location := range maps.Values(seedToNextValueMap) {
	// 	if i == 0 {
	// 		minLocation = location
	// 	}
	// 	if location < minLocation {
	// 		minLocation = location
	// 	}
	// }
	fmt.Println(nextValueList)
	return minLocation
}

func calculateStepRange(line string, stepRange []SeedMapping) []SeedMapping {
	var seedMapping SeedMapping
	fmt.Sscanf(line, "%d %d %d", &seedMapping.destinationStart, &seedMapping.start, &seedMapping.rangeValue)
	seedMapping = *newSeedMapping(seedMapping.start, seedMapping.rangeValue, seedMapping.destinationStart)
	return append(stepRange, seedMapping)
}

func initSeedMapPart2(seeds []string) (seedRanges []SeedRange) {
	for i := 0; i < len(seeds); i += 2 {
		start, _ := strconv.Atoi(seeds[i])
		rangeValue, _ := strconv.Atoi(seeds[i+1])
		seedRanges = append(seedRanges, *newSeedRange(start, rangeValue))
	}
	return
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
