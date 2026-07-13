package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/repl"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/pkg/logger"
)

// Declared only once so that flag.String is called once
var file = flag.String("file", "", "path to source file")

func main() {
	// Initialize logger with debug level in dev, info in production
	logLevel := slog.LevelInfo
	if os.Getenv("KSM_DEBUG") != "" {
		logLevel = slog.LevelDebug
	}

	logger.Init(logger.Config{
		Level:     logLevel,
		AddSource: true,
	})

	flag.Parse()

	logger.Info("Starting Kisumu REPL",
		slog.String("version", "0.4.0"),
		slog.Bool("file_mode", *file != ""),
	)

	if *file != "" {
		if err := repl.ReadFile(*file); err != nil {
			logger.Error("File execution failed",
				slog.String("path", *file),
				slog.String("error", err.Error()),
			)
			cleanup()
			os.Exit(1)
		}
		cleanup()
		return
	}

	logger.Warn("No input file provided - entering interactive mode")
	// fmt.Println("Kisumu REPL (type :exit to quit)")
	if err := repl.Start(); err != nil {
		logger.Error("Failed to start REPL",
			slog.String("error", err.Error()),
			slog.String("action", "exiting REPL"),
		)
		panic("REPL startup failed")
	}
}

func cleanup() {
	logger.Info("Cleaning up after file execution")
}
