package main

import (
	"fmt"
	"log"
	"log/slog"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/pkg/logger"
)

func main() {
	// Default log level assumed to be info
	level := slog.LevelInfo
	err := logger.InitLogger("app.log", level)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	fmt.Println("Kisumu Lang CLI")
}
