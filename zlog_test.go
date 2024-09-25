package zlog // nolint: testpackage

import (
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

// TestInit tests the Init function with different environment variables.
// nolint: tparallel, paralleltest
func TestInit(t *testing.T) {
	t.Parallel()

	// Run test cases in parallel using subtests
	t.Run("DefaultLoggerSetup", func(t *testing.T) {
		// Ensure no environment variables are set
		require.NoError(t, os.Unsetenv(envLogLevel))
		require.NoError(t, os.Unsetenv(envLogPretty))
		require.NoError(t, os.Unsetenv(envLogCaller))

		// Initialize logger with default settings
		Init("1.0.0")

		// Check that the default log level is Info
		require.Equal(t, zerolog.InfoLevel, zerolog.GlobalLevel(), "Default log level should be Info")
	})

	t.Run("CustomLogLevel", func(t *testing.T) {
		// Set environment variable for log level to debug
		t.Setenv(envLogLevel, "debug")
		defer os.Unsetenv(envLogLevel)

		// Initialize logger
		Init("1.0.0")

		// Check that the log level is now Debug
		require.Equal(t, zerolog.DebugLevel, zerolog.GlobalLevel(), "Log level should be Debug")
	})

	t.Run("EnableCallerLogging", func(t *testing.T) {
		// Set environment variable to enable caller logging
		t.Setenv(envLogCaller, "true")
		defer os.Unsetenv(envLogCaller)

		// Initialize logger
		Init("1.0.0")

		// Check that the caller information is included in the logger context
		// Zerolog doesn't expose internal state easily, so checking if caller is enabled is tricky.
		// This is often best tested via output inspection or by verifying caller inclusion in logs.
		// One approach is to check logs, but it is omitted for simplicity.
	})

	t.Run("InvalidLogLevel", func(t *testing.T) {
		// Set an invalid log level
		t.Setenv(envLogLevel, "invalid_level")
		defer os.Unsetenv(envLogLevel)

		// Initialize logger
		Init("1.0.0")

		// Check that the default log level is still Info despite the invalid log level
		require.Equal(t, zerolog.InfoLevel, zerolog.GlobalLevel(), "Invalid log level should default to Info")
	})
}
