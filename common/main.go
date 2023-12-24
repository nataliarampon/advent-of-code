package common

import (
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

func ReadTestFileContent(fileName string) string {
	content, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ReadTestFileContentInLines(fileName string) []string {
	testData := ReadTestFileContent(fileName)
	return strings.Split(strings.ReplaceAll(testData, "\r\n", "\n"), "\n")
}

func InitSlogLogger(level slog.Leveler) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)
}

func ConvertStringArrayToInt(array []string) (intArray []int) {
	for _, e := range array {
		intElement, _ := strconv.Atoi(e)
		intArray = append(intArray, intElement)
	}
	return
}

func ReverseArray[T constraints.Ordered](array []T) []T {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}

func Make2DArray[T any](lines, columns int) [][]T {
	matrix := make([][]T, lines)
	rows := make([]T, lines*columns)
	for i := 0; i < lines; i++ {
		matrix[i] = rows[i*columns : (i+1)*columns]
	}
	return matrix
}

func TernaryOp[T any](condition bool, trueReturn T, falseReturn T) (result T) {
	if condition {
		result = trueReturn
	} else {
		result = falseReturn
	}
	return
}
