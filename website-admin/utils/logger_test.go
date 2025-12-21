package utils

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/rs/zerolog"
)

func TestGetLogger_DisabledLogging(t *testing.T) {
	// Test that when disableLogging is true, we get a nop logger
	logger := GetLogger(true, false)

	// Verify it's a nop logger by checking if it outputs anything
	var buf bytes.Buffer
	testLogger := logger.Output(&buf)

	testLogger.Info().Msg("test message")
	testLogger.Debug().Msg("debug message")
	testLogger.Error().Msg("error message")

	// Nop logger should produce no output
	if buf.Len() > 0 {
		t.Errorf("Expected no output from disabled logger, got: %s", buf.String())
	}

	// Verify DefaultContextLogger is set
	if zerolog.DefaultContextLogger == nil {
		t.Error("Expected DefaultContextLogger to be set")
	}
}

func TestGetLogger_EnabledLogging_InfoLevel(t *testing.T) {
	// Save original stdout to restore later
	originalStdout := os.Stdout

	// Create a pipe to capture output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test that when disableLogging is false and debugMode is false, we get info level
	logger := GetLogger(false, false)

	// Restore stdout
	if err := w.Close(); err != nil {
		t.Fatalf("Failed to close pipe writer: %v", err)
	}
	os.Stdout = originalStdout

	// Read captured output
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	// Verify logger level is set to Info
	if logger.GetLevel() != zerolog.InfoLevel {
		t.Errorf("Expected log level to be Info (%d), got %d", zerolog.InfoLevel, logger.GetLevel())
	}

	// Verify DefaultContextLogger is set
	if zerolog.DefaultContextLogger == nil {
		t.Error("Expected DefaultContextLogger to be set")
	}
}

func TestGetLogger_EnabledLogging_DebugLevel(t *testing.T) {
	// Save original stdout to restore later
	originalStdout := os.Stdout

	// Create a pipe to capture output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Test that when disableLogging is false and debugMode is true, we get debug level
	logger := GetLogger(false, true)

	// Restore stdout
	if err := w.Close(); err != nil {
		t.Fatalf("Failed to close pipe writer: %v", err)
	}
	os.Stdout = originalStdout

	// Read captured output
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}

	// Verify logger level is set to Debug
	if logger.GetLevel() != zerolog.DebugLevel {
		t.Errorf("Expected log level to be Debug (%d), got %d", zerolog.DebugLevel, logger.GetLevel())
	}

	// Verify DefaultContextLogger is set
	if zerolog.DefaultContextLogger == nil {
		t.Error("Expected DefaultContextLogger to be set")
	}
}

func TestGetLogger_LogOutput_InfoLevel(t *testing.T) {
	// Save original stdout to restore later
	originalStdout := os.Stdout

	// Create a pipe to capture output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Get logger with info level
	logger := GetLogger(false, false)

	// Test different log levels
	logger.Debug().Msg("debug message")  // Should not appear
	logger.Info().Msg("info message")    // Should appear
	logger.Warn().Msg("warning message") // Should appear
	logger.Error().Msg("error message")  // Should appear

	// Close writer and restore stdout
	if err := w.Close(); err != nil {
		t.Fatalf("Failed to close pipe writer: %v", err)
	}
	os.Stdout = originalStdout

	// Read captured output
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}
	output := buf.String()

	// Debug message should not appear (below info level)
	if strings.Contains(output, "debug message") {
		t.Error("Debug message should not appear at Info level")
	}

	// Info, warn, and error messages should appear
	if !strings.Contains(output, "info message") {
		t.Error("Info message should appear at Info level")
	}
	if !strings.Contains(output, "warning message") {
		t.Error("Warning message should appear at Info level")
	}
	if !strings.Contains(output, "error message") {
		t.Error("Error message should appear at Info level")
	}
}

func TestGetLogger_LogOutput_DebugLevel(t *testing.T) {
	// Save original stdout to restore later
	originalStdout := os.Stdout

	// Create a pipe to capture output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Get logger with debug level
	logger := GetLogger(false, true)

	// Test different log levels
	logger.Debug().Msg("debug message")  // Should appear
	logger.Info().Msg("info message")    // Should appear
	logger.Warn().Msg("warning message") // Should appear
	logger.Error().Msg("error message")  // Should appear

	// Close writer and restore stdout
	if err := w.Close(); err != nil {
		t.Fatalf("Failed to close pipe writer: %v", err)
	}
	os.Stdout = originalStdout

	// Read captured output
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}
	output := buf.String()

	// All messages should appear at debug level
	if !strings.Contains(output, "debug message") {
		t.Error("Debug message should appear at Debug level")
	}
	if !strings.Contains(output, "info message") {
		t.Error("Info message should appear at Debug level")
	}
	if !strings.Contains(output, "warning message") {
		t.Error("Warning message should appear at Debug level")
	}
	if !strings.Contains(output, "error message") {
		t.Error("Error message should appear at Debug level")
	}
}

func TestGetLogger_ConsoleOutput_Format(t *testing.T) {
	// Save original stdout to restore later
	originalStdout := os.Stdout

	// Create a pipe to capture output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Get logger
	logger := GetLogger(false, false)

	// Log a message with timestamp
	logger.Info().Msg("test message")

	// Close writer and restore stdout
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	os.Stdout = originalStdout

	// Read captured output
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}
	output := buf.String()

	// Verify the output contains expected elements from ConsoleWriter
	if !strings.Contains(output, "test message") {
		t.Error("Output should contain the log message")
	}

	// Console writer should format the output in a human-readable way
	// It should contain timestamp and log level
	if !strings.Contains(output, "INF") && !strings.Contains(output, "INFO") {
		t.Error("Output should contain log level indicator")
	}
}

func TestGetLogger_TimestampPresent(t *testing.T) {
	// Save original stdout to restore later
	originalStdout := os.Stdout

	// Create a pipe to capture output
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Get logger
	logger := GetLogger(false, false)

	// Log a message
	beforeLog := time.Now()
	logger.Info().Msg("timestamp test")
	_ = time.Now() // afterLog not needed, just ensuring we have a reference time

	// Close writer and restore stdout
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	os.Stdout = originalStdout

	// Read captured output
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("Failed to read from pipe: %v", err)
	}
	output := buf.String()

	// The console writer should include a timestamp
	// Look for RFC3339 format patterns (YYYY-MM-DD)
	currentYear := beforeLog.Format("2006")
	if !strings.Contains(output, currentYear) {
		t.Error("Output should contain a timestamp with current year")
	}
}

func TestGetLogger_AllCombinations(t *testing.T) {
	testCases := []struct {
		name           string
		disableLogging bool
		debugMode      bool
		expectedLevel  zerolog.Level
		shouldLog      bool
	}{
		{
			name:           "Disabled logging, debug off",
			disableLogging: true,
			debugMode:      false,
			expectedLevel:  zerolog.Disabled,
			shouldLog:      false,
		},
		{
			name:           "Disabled logging, debug on",
			disableLogging: true,
			debugMode:      true,
			expectedLevel:  zerolog.Disabled,
			shouldLog:      false,
		},
		{
			name:           "Enabled logging, debug off",
			disableLogging: false,
			debugMode:      false,
			expectedLevel:  zerolog.InfoLevel,
			shouldLog:      true,
		},
		{
			name:           "Enabled logging, debug on",
			disableLogging: false,
			debugMode:      true,
			expectedLevel:  zerolog.DebugLevel,
			shouldLog:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Save original stdout
			originalStdout := os.Stdout

			// Create pipe for capturing output
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Get logger with test parameters
			logger := GetLogger(tc.disableLogging, tc.debugMode)

			// Test logging
			logger.Info().Msg("test message")

			// Restore stdout
			if err := w.Close(); err != nil {
				t.Fatal(err)
			}
			os.Stdout = originalStdout

			// Read output
			var buf bytes.Buffer
			if _, err := buf.ReadFrom(r); err != nil {
				t.Fatalf("Failed to read from pipe: %v", err)
			}
			output := buf.String()

			// Verify log level (for enabled loggers)
			if !tc.disableLogging && logger.GetLevel() != tc.expectedLevel {
				t.Errorf("Expected log level %d, got %d", tc.expectedLevel, logger.GetLevel())
			}

			// Verify output behavior
			hasOutput := len(output) > 0
			if tc.shouldLog && !hasOutput {
				t.Error("Expected log output but got none")
			}
			if !tc.shouldLog && hasOutput {
				t.Errorf("Expected no log output but got: %s", output)
			}

			// Verify DefaultContextLogger is always set
			if zerolog.DefaultContextLogger == nil {
				t.Error("Expected DefaultContextLogger to be set")
			}
		})
	}
}
