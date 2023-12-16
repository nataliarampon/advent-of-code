package main

import (
	"common"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

func main() {
	common.InitSlogLogger(slog.LevelWarn)
	data := common.ReadTestFileContentInLines("input.txt")

	races := readRaces(data)

	total := calculateResultForPart1(races)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	race := readRacePart2(data)

	total = calculateResultForPart2(race)
	fmt.Println("Result Part 2: " + fmt.Sprint(total))
}

func readRaces(data []string) (races []Race) {
	raceTimes := common.ConvertStringArrayToInt(strings.Fields(strings.Split(data[0], ":")[1]))
	raceDistances := common.ConvertStringArrayToInt(strings.Fields(strings.Split(data[1], ":")[1]))

	for i := 0; i < len(raceTimes); i++ {
		races = append(races, Race{time: raceTimes[i], record: raceDistances[i]})
	}
	return
}

func readRacePart2(data []string) Race {
	raceTime, _ := strconv.Atoi(strings.Join(strings.Fields(strings.Split(data[0], ":")[1]), ""))
	raceDistance, _ := strconv.Atoi(strings.Join(strings.Fields(strings.Split(data[1], ":")[1]), ""))

	return Race{time: raceTime, record: raceDistance}
}

func calculateResultForPart1(races []Race) int {
	for i := 0; i < len(races); i++ {
		for j := 1; j <= races[i].time; j++ {
			distance := j * (races[i].time - j)
			if distance > races[i].record {
				races[i].waysToWin++
			}
		}
	}

	total := 1
	for _, race := range races {
		total *= race.waysToWin
	}
	return total
}

func calculateResultForPart2(race Race) int {
	for j := 1; j <= race.time; j++ {
		distance := j * (race.time - j)
		if distance > race.record {
			race.waysToWin++
		}
	}

	return race.waysToWin
}
