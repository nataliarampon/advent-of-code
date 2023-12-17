package main

import (
	"common"
	"fmt"
	"log/slog"
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

func calculateResultPart2(instructions string, nodes map[string]Node) (total int) {
	return -1
}
