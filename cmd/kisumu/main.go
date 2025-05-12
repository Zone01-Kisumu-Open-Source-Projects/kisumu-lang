package main

import (
	"log/slog"
	"os"

	"github.com/Zone01-Kisumu-Open-Source-Projects/kisumu-lang/pkg/logger"
)

func main() {
	// Initialize with debug level and source locations
	logger.Init(logger.Config{
		Level:     slog.LevelDebug,
		AddSource: true,
	})

	logger.Info("Starting Kisumu Lang",
		slog.String("version", "0.4.0"),
		slog.Int("pid", os.Getpid()),
	)

	// Standard usage
	logger.Debug("Debugging information")
	logger.Warn("Potential issue detected")
}
