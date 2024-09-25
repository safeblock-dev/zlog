package zlog

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogPretty = "LOG_PRETTY"
	envLogCaller = "LOG_CALLER"

	defaultLogLevel = zerolog.InfoLevel // Default log level if none is set
)

// Init initializes the logger with optional configurations based on environment variables.
func Init(version any) {
	// Set default log level and version
	WithLevel(defaultLogLevel)
	WithVersion(version)

	// Override log level if specified in environment variable
	if logLevel := os.Getenv(envLogLevel); logLevel != "" {
		if level, err := zerolog.ParseLevel(logLevel); err == nil {
			WithLevel(level)
		} else {
			log.Warn().Str("logLevel", logLevel).Msg("Invalid log level, using default")
		}
	}

	// Enable pretty logging if specified in environment variables
	if os.Getenv(envLogPretty) == "true" || os.Getenv(envLogPretty) == "1" {
		WithPretty()
	}

	// Enable caller information if specified in environment variables
	if os.Getenv(envLogCaller) == "true" || os.Getenv(envLogCaller) == "1" {
		WithCaller()
	}
}

// WithLevel sets the global log level.
func WithLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}

// WithVersion adds the version information to the global logger.
func WithVersion(version any) {
	log.Logger = log.Logger.With().Any("version", version).Logger()
}

// WithPretty configures the logger to use a human-readable console output format.
func WithPretty() {
	log.Logger = log.Logger.Output(zerolog.ConsoleWriter{ // nolint: exhaustruct
		Out:        os.Stdout,    // Output to standard output
		TimeFormat: time.RFC3339, // Time format for logging
	})
}

// WithCaller enables logging of the caller's file and line number.
func WithCaller() {
	log.Logger = log.With().Caller().Logger()
}
