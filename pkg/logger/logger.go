package logger

import (
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"os"
)

var Logger *slog.Logger

// InitLogger relies on the slog package to provide a custom logger with more information on encountered errors logged to both standard output and log file.
func InitLogger(path string, level slog.Level) error {
	if path == "" {
		return errors.New("path to log cannot be empty")
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o744)
	if err != nil {
		return fmt.Errorf("failed openning log file: %w", err)
	}

	// Duplicate logs encountered to both standard output and a file
	multiWriter := io.MultiWriter(os.Stdout, file)
	Logger = slog.New(slog.NewTextHandler(multiWriter, &slog.HandlerOptions{
		Level: level,
	}))

	if Logger == nil {
		return errors.New("failed creating logger")
	}
	return nil
}

// Info attributes its logs with the INFO tag.
func Info(msg string, args ...slog.Attr) {
	if Logger == nil {
		log.Printf("Logger not initialized: %s\n", msg)
	}

	argSlice := make([]any, 0)

	for _, arg := range args {
		argSlice = append(argSlice, arg.Key, arg.Value)
	}
	Logger.Info(msg, argSlice...)
}

// Info attributes its logs with the ERROR tag.
func Error(msg string, args ...slog.Attr) {
	if Logger == nil {
		log.Printf("Logger not initialized: %s\n", msg)
	}

	argSlice := make([]any, 0)

	for _, arg := range args {
		argSlice = append(argSlice, arg.Key, arg.Value)
	}
	Logger.Error(msg, argSlice...)
}

// Info attributes its logs with the WARN tag.
func Warn(msg string, args ...slog.Attr) {
	if Logger == nil {
		log.Printf("Logger not initialized: %s\n", msg)
	}

	argSlice := make([]any, 0)

	for _, arg := range args {
		argSlice = append(argSlice, arg.Key, arg.Value)
	}
	Logger.Warn(msg, argSlice...)
}