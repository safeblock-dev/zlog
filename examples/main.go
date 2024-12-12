package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/safeblock-dev/zlog"
)

// Example of usage
func main() {
	// Set environment variables to simulate configuration
	os.Setenv("LOG_LEVEL", "info")
	os.Setenv("LOG_FORMAT", "text")
	os.Setenv("LOG_COLOR", "true")
	os.Setenv("LOG_CALLER", "true")

	// Initialize logger with version information
	zlog.Init()

	// Log an informational message
	log.Info().Str("foo", "bar").Msg("hello world")
}
