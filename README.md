# ffmpeg-go

This is a Go implementation of an ffmpeg CLI entry point.

Currently, it functions as a robust wrapper around the system's `ffmpeg` binary, ensuring compatibility and reliability while providing a Go-based structure for future pure-Go implementations or custom logic extensions.

## Building

```bash
go build -o ffmpeg-go cmd/ffmpeg/main.go
```

## Usage

You can use `ffmpeg-go` exactly like `ffmpeg`. It forwards all arguments to the underlying ffmpeg installation.

```bash
./ffmpeg-go -i input.mp4 output.mkv
./ffmpeg-go -version
```

## Structure

- `cmd/ffmpeg`: Main entry point.
- `internal/transcoder`: Interface and implementation for media processing.
- `internal/cli`: CLI argument handling and execution logic.
