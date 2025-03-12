package main

import (
	"flag"
	"log"
	"log/slog"
	"os"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/repl"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/pkg/logger"
)

func main() {
	_ = os.Mkdir("data", 0o744)

	// Default log level assumed to be info
	level := slog.LevelInfo
	err := logger.InitLogger("data/app.log", level)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	file := flag.String("file", "", "File to read input from")
	flag.Parse()

	if *file != "" {
		repl.ReadFile(*file)
	} else {
		logger.Warn("No file provided. Starting REPL...")
		repl.Start()
	}
}
