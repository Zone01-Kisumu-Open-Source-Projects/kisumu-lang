package main

import (
	"log"
	"log/slog"
	"os"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/pkg/logger"
)

func main() {
	_ = os.Mkdir("data", 0744)

	// Default log level assumed to be info
	level := slog.LevelInfo
	err := logger.InitLogger("data/app.log", level)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	logger.Info("Kisumu Lang CLI")
}
