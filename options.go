package zlog

import (
	"os"
	"strconv"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// LoggerFormatMode defines the mode for log formatting.
type LoggerFormatMode string

const (
	// FormatModeJSON outputs logs in JSON format.
	FormatModeJSON LoggerFormatMode = "json"
	// FormatModeText outputs logs in a human-readable text format.
	FormatModeText LoggerFormatMode = "text"
)

// loggerOptions contains the configuration for the logger.
type loggerOptions struct {
	format     LoggerFormatMode
	level      zerolog.Level
	version    any
	withCaller bool
	withColor  bool
}

// defaultLoggerOptions returns the default logger settings.
// Default format is JSON, the log level is set to Info,
// and both withCaller and withColor are disabled.
// These settings can be overridden programmatically or via environment variables.
func defaultLoggerOptions() *loggerOptions {
	return &loggerOptions{
		withColor:  false,
		withCaller: false,
		format:     FormatModeJSON,
		level:      zerolog.InfoLevel,
		version:    nil,
	}
}

// envLoggerOptions overrides logger settings based on environment variables.
// Environment variables have precedence over programmatic options.
func envLoggerOptions(loggerOpts *loggerOptions) {
	// Process boolean environment variables.
	setBoolEnv("LOG_CALLER", &loggerOpts.withCaller)
	setBoolEnv("LOG_COLOR", &loggerOpts.withColor)

	// LOG_FORMAT: override the log output format.
	if val, ok := os.LookupEnv("LOG_FORMAT"); ok {
		switch LoggerFormatMode(val) {
		case FormatModeJSON, FormatModeText:
			loggerOpts.format = LoggerFormatMode(val)
		default:
			log.Warn().Msgf("Invalid '%s' as format mode, using default", val)
		}
	}

	// LOG_LEVEL: set the log level.
	if val, ok := os.LookupEnv("LOG_LEVEL"); ok {
		if lvl, err := zerolog.ParseLevel(val); err != nil {
			log.Error().Err(err).Msg("Invalid LOG_LEVEL value, using default")
		} else {
			loggerOpts.level = lvl
		}
	}
}

// setBoolEnv checks for the given environment variable and, if found,
// attempts to parse its value as a boolean. On success, the parsed value
// is stored in the provided target. In case of an error, the default is kept.
func setBoolEnv(envVar string, target *bool) {
	if val, ok := os.LookupEnv(envVar); ok {
		if b, err := strconv.ParseBool(val); err != nil {
			log.Error().Err(err).Msgf("Invalid %s value, using default", envVar)
		} else {
			*target = b
		}
	}
}

// LoggerOption defines an interface for modifying logger settings.
type LoggerOption interface {
	apply(*loggerOptions) //nolint: inamedparam
}

// funcLoggerOption wraps a function that modifies loggerOptions into an implementation of LoggerOption.
type funcLoggerOption struct {
	f func(*loggerOptions)
}

func (fdo *funcLoggerOption) apply(do *loggerOptions) {
	fdo.f(do)
}

func newFuncLoggerOption(f func(*loggerOptions)) *funcLoggerOption {
	return &funcLoggerOption{f: f}
}

// WithCaller enables or disables logging of the caller's file and line number.
// When called without parameters, it defaults to enabling (true).
func WithCaller(caller ...bool) LoggerOption { //nolint: ireturn
	enabled := true
	if len(caller) > 0 {
		enabled = caller[0]
	}

	return newFuncLoggerOption(func(o *loggerOptions) {
		o.withCaller = enabled
	})
}

// WithColor enables or disables colored output for text logs.
// When called without parameters, it defaults to enabling (true).
func WithColor(color ...bool) LoggerOption { //nolint: ireturn
	enabled := true
	if len(color) > 0 {
		enabled = color[0]
	}

	return newFuncLoggerOption(func(o *loggerOptions) {
		o.withColor = enabled
	})
}

// WithFormat sets the log output format.
func WithFormat(mode LoggerFormatMode) LoggerOption { //nolint: ireturn
	return newFuncLoggerOption(func(o *loggerOptions) {
		o.format = mode
	})
}

// WithLevel sets the log level.
func WithLevel(level zerolog.Level) LoggerOption { //nolint: ireturn
	return newFuncLoggerOption(func(o *loggerOptions) {
		o.level = level
	})
}

// WithVersion sets the application version that will be added to the logs.
func WithVersion(version any) LoggerOption { //nolint: ireturn
	return newFuncLoggerOption(func(o *loggerOptions) {
		o.version = version
	})
}
