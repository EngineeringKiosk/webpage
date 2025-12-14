package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func GetLogger(disableLogging, debugMode bool) zerolog.Logger {
	// Return a No-operation logger if logging is disabled
	if disableLogging {
		// Disabled logger
		logger := zerolog.Nop()
		zerolog.DefaultContextLogger = &logger
		return logger
	}

	// Human readable logger (with color)
	loggerConsoleOutput := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := zerolog.New(loggerConsoleOutput).With().Timestamp().Logger()
	zerolog.DefaultContextLogger = &logger

	// Set minimum log level
	minimumLogLevel := zerolog.InfoLevel
	if debugMode {
		minimumLogLevel = zerolog.DebugLevel
	}
	logger = logger.Level(minimumLogLevel)

	return logger
}
