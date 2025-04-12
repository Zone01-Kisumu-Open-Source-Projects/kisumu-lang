package logger

import (
	"bytes"
	"log/slog"
	"sync"
	"testing"
)

func TestLogger(t *testing.T) {
	t.Run("Init", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Error("Initialization panicked unexpectedly")
			}
		}()

		var buf bytes.Buffer
		Init(Config{
			Level:  slog.LevelDebug,
			Output: &buf,
		})

		Info("test message")
		if buf.Len() == 0 {
			t.Error("No log output generated")
		}
	})

	t.Run("Uninitialized", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("Expected panic when using uninitialized logger")
			}
		}()

		// Reset for test
		instance = nil
		initOnce = sync.Once{}
		Info("should panic")
	})
}
