package main

import (
	"flag"
	"log/slog"
	"os"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/internal/repl"
	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/pkg/logger"
)

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

	defer func() {
		if r := recover(); r != nil {
			logger.Error("REPL panic recovered",
				slog.Any("error", r),
				slog.String("action", "restarting REPL"),
			)
			main() // Restart on panic
		}
	}()

	file := flag.String("file", "", "Source file to execute")
	flag.Parse()

	logger.Info("Starting Kisumu REPL",
		slog.String("version", "0.4.0"),
		slog.Bool("file_mode", *file != ""),
	)

	if *file != "" {
		if err := repl.RunFile(*file); err != nil {
			logger.Error("File execution failed",
				slog.String("path", *file),
				slog.String("error", err.Error()),
			)
			os.Exit(1)
		}
		return
	}

	logger.Warn("No input file provided - entering interactive mode")
	// fmt.Println("Kisumu REPL (type :exit to quit)")
	repl.Start()
}
