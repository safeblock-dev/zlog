package zlog

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Init initializes the logger with optional configurations based on environment variables.
func Init(opts ...LoggerOption) {
	loggerOpts := defaultLoggerOptions()
	for _, opt := range opts {
		opt.apply(loggerOpts)
	}
	envLoggerOptions(loggerOpts)

	zerolog.SetGlobalLevel(loggerOpts.level)

	if loggerOpts.format == FormatModeText {
		log.Logger = log.Logger.Output(zerolog.ConsoleWriter{ //nolint: exhaustruct
			NoColor:    !loggerOpts.withColor,
			Out:        os.Stderr,    // Output to stderr.
			TimeFormat: time.RFC3339, // Time format for logging.
		})
	}

	if loggerOpts.version != nil {
		log.Logger = log.With().Any("version", loggerOpts.version).Logger()
	}

	if loggerOpts.withCaller {
		log.Logger = log.With().Caller().Logger()
	}
}
