# GEMINI Project Context

This project was bootstrapped by a Gemini CLI agent on December 25, 2025.

## Goal
To rewrite FFmpeg entirely in Go, without relying on existing C libraries or wrappers.

## Evolution
1. **Wrapper**: Initially created as a wrapper around the system `ffmpeg` binary.
2. **CGO**: Briefly evolved into a CGO-based application linking against `libavcodec`/`libavformat`.
3. **Pure Go**: Pivot to a from-scratch rewrite.

## Architecture
The current implementation focuses on the **ISO Base Media File Format (ISOBMFF)** parser.
- **Entry Point**: `cmd/ffmpeg/main.go`
- **Core Logic**: `internal/isobmff/` containing `reader.go` (binary parsing) and `printer.go` (structure visualization).
