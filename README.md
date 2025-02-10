# ZLog

ZLog is a lightweight wrapper around the popular [`zerolog`](https://github.com/rs/zerolog) logging library. It allows you to configure logging easily through environment variables, enabling you to set log levels, use pretty logging, and include caller information with minimal setup.

## Features

- Configurable log levels via environment variables
- Human-readable (pretty) logging output
- Optional inclusion of caller information (file and line number)
- Simple one-time initialization

## Installation

Install ZLog using the following command:

```bash
go get github.com/safeblock-dev/zlog
```

## Usage

Initialize the logger once at the start of your application (for example, in your `main` function):

```go
package main

import (
    "github.com/safeblock-dev/zlog"
    "github.com/rs/zerolog/log"
)

func main() {
    // Initialize the logger
    zlog.Init()

    // Log messages
    log.Info().Msg("This is an informational log")
    log.Error().Msg("This is an error log")
}
```

## Environment Variables

ZLog can be configured using the following environment variables:

| Environment Variable | Description                                                              | Default |
|----------------------|--------------------------------------------------------------------------|---------|
| `LOG_CALLER`         | Include caller information (file and line number) in log entries         | `false` |
| `LOG_COLOR`          | Enable pretty logging output (works with `LOG_FORMAT=text`)              | `false` |
| `LOG_FORMAT`         | Set the log output format (`json` or `text`)                             | `json`  |
| `LOG_LEVEL`          | Set the log level (e.g., `debug`, `info`, `warn`, `error`)                | `info`  |

Example configuration:

```bash
export LOG_CALLER=true
export LOG_COLOR=true
export LOG_FORMAT=text
export LOG_LEVEL=debug
```

## License

ZLog is open-source software distributed under the MIT License.

## Contributions

Contributions are welcome! Feel free to fork the repository, make improvements, and submit pull requests.

## Acknowledgements

This package is a simple wrapper around the excellent [`zerolog`](https://github.com/rs/zerolog) library.
