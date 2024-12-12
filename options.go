package zlog

import (
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

type LoggerFormatMode string

const (
	FormatModeJSON LoggerFormatMode = "json"
	FormatModeText LoggerFormatMode = "text"
)

type loggerOptions struct {
	format     LoggerFormatMode
	level      zerolog.Level
	version    any
	withCaller bool
	withColor  bool
}

func defaultLoggerOptions() *loggerOptions {
	return &loggerOptions{
		withColor:  false,
		withCaller: false,
		format:     FormatModeJSON,
		level:      zerolog.InfoLevel,
		version:    nil,
	}
}
func envLoggerOptions(loggerOpts *loggerOptions) {
	// Enable caller information if specified in environment variables
	if val, ok := os.LookupEnv("LOG_CALLER"); ok {
		v, err := strconv.ParseBool(val)
		if err != nil {
			panic(err)
		}
		loggerOpts.withCaller = v
	}

	// Enable pretty logging if specified in environment variables
	if val, ok := os.LookupEnv("LOG_COLOR"); ok {
		v, err := strconv.ParseBool(val)
		if err != nil {
			panic(err)
		}
		loggerOpts.withColor = v
	}

	// Override log format if specified in environment variable
	if val, ok := os.LookupEnv("LOG_FORMAT"); ok {
		v, err := strconv.ParseBool(val)
		if err != nil {
			panic(err)
		}
		loggerOpts.withColor = v
	}

	// Override log level if specified in environment variable
	if val, ok := os.LookupEnv("LOG_LEVEL"); ok {
		v, err := zerolog.ParseLevel(val)
		if err != nil {
			panic(err)
		}
		loggerOpts.level = v
	}
}

// LoggerOption configures the logger.
type LoggerOption interface {
	apply(*loggerOptions)
}

// funcLoggerOption wraps a function that modifies loggerOptions into an
// implementation of the LoggerOption interface.
type funcLoggerOption struct {
	f func(*loggerOptions)
}

func (fdo *funcLoggerOption) apply(do *loggerOptions) {
	fdo.f(do)
}

func newFuncLoggerOption(f func(*loggerOptions)) *funcLoggerOption {
	return &funcLoggerOption{
		f: f,
	}
}

// WithCaller sets logging of the caller's file and line number mode (default: true).
func WithCaller(caller ...bool) LoggerOption {
	ok := true
	if len(caller) > 0 {
		ok = caller[0]
	}

	return newFuncLoggerOption(func(o *loggerOptions) {
		o.withCaller = ok
	})
}

// WithColor sets the color mode (default: true).
func WithColor(color ...bool) LoggerOption {
	ok := true
	if len(color) > 0 {
		ok = color[0]
	}

	return newFuncLoggerOption(func(o *loggerOptions) {
		o.withColor = ok
	})
}

// WithFormat sets the logger format.
func WithFormat(mode LoggerFormatMode) LoggerOption {
	return newFuncLoggerOption(func(o *loggerOptions) {
		o.format = mode
	})
}

// WithLevel sets the logger level.
func WithLevel(level zerolog.Level) LoggerOption {
	return newFuncLoggerOption(func(o *loggerOptions) {
		o.level = level
	})
}

// WithVersion sets the app version.
func WithVersion(version any) LoggerOption {
	return newFuncLoggerOption(func(o *loggerOptions) {
		o.version = version
	})
}
