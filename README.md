# ZLog - Simple Configurable Logging with `zerolog`

ZLog is a lightweight wrapper around the popular `zerolog` logging library, offering easy configuration through environment variables. The library allows you to set log levels, enable pretty logging, and include caller information without complex setup.

## Features

- Configurable log levels (via environment variables)
- Pretty logging for human-readable logs
- Optionally include caller information in logs
- Simple initialization with version tagging

## Installation

To install ZLog, use the following `go` command:

```bash
go get github.com/safeblock-dev/zlog
```

## Usage

### Initializing the Logger

ZLog is designed to be initialized once during your application's startup, typically in the `main` function. You can specify a version number that will be included in every log entry.

```go
package main

import (
    "github.com/safeblock-dev/zlog"
    "github.com/rs/zerolog/log"
)

func main() {
    // Initialize the logger with a version string
    zlog.Init()

    // Log messages
    log.Info().Str("key", "value").Msg("This is an informational log")
    log.Error().Msg("This is an error log")
}
```

### Environment Variables

ZLog behavior is configured through environment variables. These variables allow you to customize the log level, enable pretty logs, and include caller information.

| Environment Variable | Description                                                                   | Default |
|----------------------|-------------------------------------------------------------------------------|---------|
| `LOG_CALLER`         | Includes caller information in the log entries                                | `false` |
| `LOG_COLOR`          | Enables pretty logging for human-readable output (works with LOG_FORMAT=text) | `false` |
| `LOG_FORMAT`         | Sets the log format mode (`json`, `text`)                                     | `json`  |
| `LOG_LEVEL`          | Sets the log level (`debug`, `info`, `error`, etc.)                           | `info`  |

#### Example Configuration

```bash
export LOG_CALLER=true
export LOG_COLOR=true
export LOG_FORMAT=text
export LOG_LEVEL=debug
```

This setup configures ZLog to log messages at the `debug` level, output logs in a human-readable format, and include caller information.

### Logging Levels

ZLog uses `zerolog`'s built-in logging levels:

- `trace`
- `debug`
- `info`
- `warn`
- `error`
- `fatal`
- `panic`

You can configure the log level using the `LOG_LEVEL` environment variable:

```go
package main

import (
    "os"
    "github.com/safeblock-dev/zlog"
    "github.com/rs/zerolog/log"
)

func main() {
    // Set the log level to debug via environment variable
    os.Setenv("LOG_LEVEL", "debug")

    zlog.Init()

    log.Debug().Msg("This is a debug message")
    log.Info().Msg("This is an info message")
}
```

### Pretty Logging

You can enable pretty logging (human-readable format) by setting the `LOG_FORMAT` and `LOG_COLOR` environment variable:

```bash
export LOG_COLOR=true
export LOG_FORMAT=text
```

When enabled, the logs will be printed in a more readable format:

```bash
2024-09-20T13:50:59+00:00 INF This is an info message
```

### Caller Information

You can enable caller information logging by setting the `LOG_CALLER` environment variable. This will include the file and line number from which the log was generated.

```bash
export LOG_CALLER=true
```

When enabled, log entries will include caller information like this:

```bash
{"level":"info","time":"2024-12-12T23:53:15+05:00","caller":"main.go:42","message":"This is an info message"}
```

### Version Information

You can set version information logging by setting the `LOG_CALLER` environment variable. This will include the file and line number from which the log was generated.

```go
package main

import (
    "errors"
    "github.com/safeblock-dev/zlog"
    "github.com/rs/zerolog/log"
)

func main() {
    zlog.Init(zlog.WithVersion("1.0.0"))

	  log.Info().Msg("This is an info message")
}
```

```bash
{"level":"info","version":"1.0.0","time":"2024-12-12T23:51:22+05:00","message":"This is an info message"}
```

### Multiply options

You can set version information logging by setting the `LOG_CALLER` environment variable. This will include the file and line number from which the log was generated.

```go
package main

import (
    "errors"
    "github.com/safeblock-dev/zlog"
    "github.com/rs/zerolog/log"
)

func main() {
	zlog.Init(
		zlog.WithColor(),
		zlog.WithFormat(zlog.FormatModeText),
		zlog.WithVersion("1.0.0"),
	)

	  log.Info().Msg("This is an info message")
}
```

```bash
2024-09-20T13:50:59+00:00 INF This is an info message version=1.0.0
```

### Error Handling

ZLog uses `zerolog`'s standard error handling. Here's an example of logging an error message:

```go
package main

import (
    "errors"
    "github.com/safeblock-dev/zlog"
    "github.com/rs/zerolog/log"
)

func main() {
    zlog.Init()

    err := errors.New("something went wrong")
    log.Error().Err(err).Msg("An error occurred")
}
```

### Testing the Logger

You can write tests for ZLog to verify that logging behaves as expected. ZLog supports logging to a buffer, allowing you to capture and inspect log output during tests.

```go
package zlog_test

import (
    "bytes"
    "testing"
    "github.com/stretchr/testify/require"
    "github.com/rs/zerolog/log"
    "github.com/safeblock-dev/zlog"
)

func TestLogging(t *testing.T) {
    var buf bytes.Buffer
    log.Logger = log.Output(&buf)

    zlog.Init()
    log.Info().Msg("test log")

    require.Contains(t, buf.String(), "test log")
}
```

## License

ZLog is open-source and distributed under the MIT License. Feel free to use it in your projects!

---

## Contribution

We welcome contributions! If you'd like to improve ZLog, feel free to fork the repository, make your changes, and open a pull request.

---

## Acknowledgements

This package is a simple wrapper around the fantastic [zerolog](https://github.com/rs/zerolog) library.
