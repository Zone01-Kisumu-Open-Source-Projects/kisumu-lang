package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"sync"
)

var (
	instance *slog.Logger
	initOnce sync.Once
	mu       sync.RWMutex
)

// Config holds logger configuration
type Config struct {
	Level     slog.Level
	Output    io.Writer // Default: os.Stdout
	AddSource bool      // Include call site information
}

// Init initializes the global logger with safe defaults
func Init(cfg Config) {
	initOnce.Do(func() {
		if cfg.Output == nil {
			cfg.Output = os.Stdout
		}

		mu.Lock()
		defer mu.Unlock()

		instance = slog.New(slog.NewTextHandler(cfg.Output, &slog.HandlerOptions{
			Level:     cfg.Level,
			AddSource: cfg.AddSource,
		}))
	})
}

// getLogger safely returns the logger instance
func getLogger() *slog.Logger {
	mu.RLock()
	defer mu.RUnlock()

	if instance == nil {
		panic("logger not initialized: call logger.Init() first")
	}
	return instance
}

// Log methods with attribute processing
func logAttrs(lvl slog.Level, msg string, attrs ...slog.Attr) {
	var args []any
	for _, attr := range attrs {
		args = append(args, attr.Key, attr.Value)
	}
	getLogger().Log(context.TODO(), lvl, msg, args...)
}

// Public API
func Debug(msg string, attrs ...slog.Attr) {
	logAttrs(slog.LevelDebug, msg, attrs...)
}

func Info(msg string, attrs ...slog.Attr) {
	logAttrs(slog.LevelInfo, msg, attrs...)
}

func Warn(msg string, attrs ...slog.Attr) {
	logAttrs(slog.LevelWarn, msg, attrs...)
}

func Error(msg string, attrs ...slog.Attr) {
	logAttrs(slog.LevelError, msg, attrs...)
}
