module example

go 1.21.6

require (
	github.com/rs/zerolog v1.33.0
	github.com/safeblock-dev/zlog v0.0.0-00010101000000-000000000000
)

replace github.com/safeblock-dev/zlog => ./..

require (
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.25.0 // indirect
)
