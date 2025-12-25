# ffmpeg-go

**A pure Go implementation of FFmpeg.**

This project aims to be a complete rewrite of FFmpeg in native Go, removing all dependencies on C libraries (no `cgo`, no `libav*`).

## Current Status

- **ISOBMFF (MP4) Parser**: Implemented. The tool can currently inspect the internal atom/box structure of MP4 files.

## Building

```bash
go build -o ffmpeg-go cmd/ffmpeg/main.go
```

## Usage

```bash
./ffmpeg-go -i input.mp4
```

This will output the tree structure of the MP4 boxes (atoms).

## Structure

- `cmd/ffmpeg`: Main CLI entry point.
- `internal/isobmff`: Pure Go package for parsing ISO Base Media File Format.