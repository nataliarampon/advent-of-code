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

	totalPart2 := calculateResultForPart2BruteForce(data)
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
						destValue := calculateDisplacement(current, stepRange[i].destinationStart, stepRange[i].start)
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

func calculateResultForPart2BruteForce(data []string) int {
	var stepRange []SeedMapping
	var newValuesList []int

	seeds := strings.Split(data[0], " ")[1:]
	currentValueList := initCurrentValueArrayBruteForce(seeds)

	processingMap := false
	for _, line := range data[2:] {
		fmt.Println("Line: " + line)
		if processingMap && line != "" {
			stepRange = calculateStepRange(line, stepRange)
		} else {
			for i := 0; i < len(stepRange); i++ {
				for j, current := range currentValueList {
					if current >= stepRange[i].start && current <= stepRange[i].end {
						destValue := calculateDisplacement(current, stepRange[i].destinationStart, stepRange[i].start)
						newValuesList = append(newValuesList, destValue)
						removeArrayIndex(currentValueList, j)
					}
				}
			}
			if line == "" {
				stepRange = nil
				currentValueList = append(currentValueList, newValuesList...)
				newValuesList = nil
			}
		}
		processingMap = getProcessingStatus(line)
	}

	var minLocation int
	for i, location := range currentValueList {
		if i == 0 {
			minLocation = location
		}
		if location < minLocation {
			minLocation = location
		}
	}

	return minLocation
}

// TODO: fix this
func calculateResultForPart2RangeMapping(data []string) int {
	var stepRanges []SeedMapping
	var newRangeMappings []SeedRange

	seeds := strings.Split(data[0], " ")[1:]
	nextValueList := initSeedMapPart2(seeds)

	processingMap := false
	for _, line := range data[1:] {
		if processingMap && line != "" {
			stepRanges = calculateStepRange(line, stepRanges)
		} else {
			for _, seedRange := range nextValueList {
				for _, destRange := range stepRanges {
					switch {
					case seedRange.start >= destRange.start && seedRange.end <= destRange.end:
						newRange := newSeedRange(calculateDisplacement(seedRange.start, destRange.destinationStart, destRange.start), seedRange.rangeValue)
						newRangeMappings = append(newRangeMappings, *newRange)
					case seedRange.start < destRange.start && seedRange.end > destRange.end:
						newRange := newSeedRange(calculateDisplacement(destRange.start, destRange.destinationStart, destRange.start), destRange.end-destRange.start)
						newRangeMappings = append(newRangeMappings, *newRange)
					case seedRange.start < destRange.start && seedRange.end > destRange.start && seedRange.end <= destRange.end:
						newRange := newSeedRange(calculateDisplacement(destRange.start, destRange.destinationStart, destRange.start), seedRange.end-destRange.start)
						newRangeMappings = append(newRangeMappings, *newRange)
					case seedRange.start >= destRange.start && seedRange.start < destRange.end && seedRange.end > destRange.end:
						newRange := newSeedRange(calculateDisplacement(seedRange.start, destRange.destinationStart, destRange.start), destRange.end-seedRange.start)
						newRangeMappings = append(newRangeMappings, *newRange)
					}
					// i know the problem here. part of seed ranges that don't fall into any destination ranges :(
					// but i'm too lazy to fix it rn, so I'll brute force it
				}
			}
			stepRanges = nil
			if newRangeMappings != nil {
				nextValueList = newRangeMappings
			}
			newRangeMappings = nil
		}
		processingMap = getProcessingStatus(line)
	}

	var minLocation int
	for i, locationRange := range nextValueList {
		if i == 0 {
			minLocation = locationRange.start
		}
		if locationRange.start < minLocation {
			minLocation = locationRange.start
		}
	}
	return minLocation
}

func calculateDisplacement(seedStart int, destinationStart int, rangeStart int) int {
	return seedStart + (destinationStart - rangeStart)
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

func initCurrentValueArrayBruteForce(seeds []string) (currentValueList []int) {
	seedRanges := initSeedMapPart2(seeds)
	for i, seedRange := range seedRanges {
		fmt.Println("Generating seed range " + strconv.Itoa(i))
		for seed := seedRange.start; seed <= seedRange.end; seed++ {
			currentValueList = append(currentValueList, seed)
		}
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

func removeArrayIndex(array []int, i int) []int {
	array[i] = array[len(array)-1]
	return array[:len(array)-1]
}
