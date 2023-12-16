package main

import (
	"common"
	"fmt"
	"log/slog"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func main() {
	common.InitSlogLogger(slog.LevelWarn)
	data := common.ReadTestFileContentInLines("input.txt")

	total := calculateResultForPart1(data)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	total = calculateResultForPart2BruteForce(data)
	fmt.Println("Result Part 2: " + fmt.Sprint(total))

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
	var stepRange [7][]SeedMapping

	seeds := strings.Split(data[0], " ")[1:]
	seedsList := initSeedMapPart2(seeds)

	processingMap := true
	j := 0
	for i := len(data) - 2; i > 1; i-- {
		if processingMap && !strings.Contains(data[i], "map:") {
			stepRange[j] = calculateStepRange(data[i], stepRange[j])
		} else if !processingMap {
			j++
		}
		processingMap = getProcessingStatusBottomDown(data[i])
	}

	sort.Slice(stepRange[0], func(i, j int) bool { return stepRange[0][i].destinationStart > stepRange[0][j].destinationStart })
	maxLocation := stepRange[0][0].destinationStart + stepRange[0][0].rangeValue - 1

	for location := 0; location <= maxLocation; location++ {
		if location%1000000 == 0 {
			slog.Info("Testing location: %d\n", location)
		}
		next := location
		for _, currentStep := range stepRange {
			for _, currentStepMapping := range currentStep {
				if next >= currentStepMapping.destinationStart && next < currentStepMapping.destinationStart+currentStepMapping.rangeValue {
					next = calculateDisplacement(next, currentStepMapping.start, currentStepMapping.destinationStart)
					break
				}
			}
		}

		if hasSeed(next, seedsList) {
			return location
		}
	}

	return -1
}

func calculateDisplacement(seedStart int, destinationStart int, rangeStart int) int {
	return seedStart + (destinationStart - rangeStart)
}

func hasSeed(seed int, seedsList []SeedRange) bool {
	for _, seedRange := range seedsList {
		if seed >= seedRange.start && seed <= seedRange.end {
			return true
		}
	}
	return false
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

func getProcessingStatusBottomDown(line string) (processingMap bool) {
	if strings.Contains(line, "map:") {
		processingMap = false
	} else {
		processingMap = true
	}
	return
}
