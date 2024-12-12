package main

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/safeblock-dev/zlog"
)

// Example of usage
func main() {
	// Set environment variables to simulate configuration
	_ = os.Setenv("LOG_LEVEL", "info")
	_ = os.Setenv("LOG_PRETTY", "true")
	_ = os.Setenv("LOG_CALLER", "true")

	// Initialize logger with version information
	zlog.Init("1.0.3")

	// Log an informational message
	log.Info().Str("foo", "bar").Msg("hello world")
}
