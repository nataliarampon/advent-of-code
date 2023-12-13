package common

import (
	"log"
	"log/slog"
	"os"
	"strings"
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
