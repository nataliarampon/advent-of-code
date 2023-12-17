package main

import (
	"common"
	"fmt"
	"log/slog"

	"golang.org/x/exp/maps"
)

func main() {
	common.InitSlogLogger(slog.LevelInfo)
	data := common.ReadTestFileContentInLines("input.txt")

	instructions := readInstructions(data)
	nodes := readNodes(data[2:])

	total := calculateResultPart1(instructions, nodes)
	fmt.Println("Result Part 1: " + fmt.Sprint(total))

	total = calculateResultPart2(instructions, nodes)
	fmt.Println("Result Part 2: " + fmt.Sprint(total))
}

func readInstructions(data []string) string {
	return data[0]
}

func readNodes(data []string) map[string]Node {
	var key string
	nodeMap := make(map[string]Node)
	for _, line := range data {
		node := new(Node)
		fmt.Sscanf(line, "%3s = (%3s, %3s)", &key, &node.left, &node.right)
		nodeMap[key] = *node
	}
	return nodeMap
}

func calculateResultPart1(instructions string, nodes map[string]Node) (stepCount int) {
	const END = "ZZZ"
	const START = "AAA"

	step := START
	i := 0

	for step != END {
		cyclingIndex := i % len(instructions)
		instr := instructions[cyclingIndex]

		if instr == 'L' {
			step = nodes[step].left
		} else {
			step = nodes[step].right
		}
		i++
		stepCount++
	}

	return
}

func calculateResultPart2(instructions string, nodes map[string]Node) (stepCount int) {
	steps := getStartingPoints(maps.Keys(nodes))
	loops := make([]int, len(steps))

	i := 0
	for j := 0; j < len(steps); j++ {
		stepCount := 0
		for steps[j][2] != 'Z' {
			cyclingIndex := i % len(instructions)
			instr := instructions[cyclingIndex]
			if instr == 'L' {
				steps[j] = nodes[steps[j]].left
			} else {
				steps[j] = nodes[steps[j]].right
			}
			stepCount++
			i++
		}
		loops[j] = stepCount
	}

	return loopsLCM(loops...)
}

func getStartingPoints(nodes []string) (starts []string) {
	for _, node := range nodes {
		if node[2] == 'A' {
			starts = append(starts, node)
		}
	}
	return
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func loopsLCM(numbers ...int) int {
	result := 1
	for i := 0; i < len(numbers); i++ {
		result *= numbers[i] / gcd(result, numbers[i])
	}

	return result
}
