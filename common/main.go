package common

import (
	"log"
	"log/slog"
	"os"
)

func ReadTestFileContent(fileName string) string {
	content, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func InitSlogLogger(level slog.Leveler) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logger)
}
