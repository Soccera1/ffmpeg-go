package isobmff

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// BoxHeader represents the standard 8-byte header of an MP4 box (atom).
type BoxHeader struct {
	Size uint32
	Type [4]byte
}

// Box represents a parsed MP4 box.
type Box struct {
	Header   BoxHeader
	Offset   int64
	Payload  []byte
	Children []*Box
}

// Reader handles the parsing of the MP4 file structure.
type Reader struct {
	file *os.File
}

// NewReader creates a new Reader for the given file path.
func NewReader(path string) (*Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &Reader{file: f}, nil
}

// Close closes the underlying file.
func (r *Reader) Close() error {
	return r.file.Close()
}

// ReadBoxes reads all top-level boxes and their children (recursively for containers).
func (r *Reader) ReadBoxes() ([]*Box, error) {
	info, err := r.file.Stat()
	if err != nil {
		return nil, err
	}
	return r.readBoxesRange(0, info.Size())
}

// readBoxesRange reads boxes within a specific file range.
func (r *Reader) readBoxesRange(start, end int64) ([]*Box, error) {
	var boxes []*Box
	offset := start

	for offset < end {
		if _, err := r.file.Seek(offset, io.SeekStart); err != nil {
			return nil, err
		}

		header, err := r.readHeader()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		// Basic validation: box size cannot be smaller than header (8 bytes)
		// Usually size=0 means "rest of file", and size=1 means "extended size" (64-bit).
		// For this MVP, we'll handle standard 32-bit sizes > 8.
		boxSize := int64(header.Size)
		if boxSize == 0 {
			boxSize = end - offset
		} else if boxSize == 1 {
			// Handle 64-bit size (largesize)
			var largeSize uint64
			err = binary.Read(r.file, binary.BigEndian, &largeSize)
			if err != nil {
				return nil, err
			}
			boxSize = int64(largeSize)
			// Header is effectively 16 bytes now (8 regular + 8 extended)
		}

		box := &Box{
			Header: header,
			Offset: offset,
		}

		// Identify container boxes to recurse into
		// Common container atoms: moov, trak, mdia, minf, dinf, stbl
		boxType := string(header.Type[:])
		if isContainer(boxType) {
			// Header size is usually 8, but 16 for large size.
			headerSize := int64(8)
			if header.Size == 1 {
				headerSize = 16
			}
			
			children, err := r.readBoxesRange(offset+headerSize, offset+boxSize)
			if err != nil {
				return nil, fmt.Errorf("failed to read children of %s: %w", boxType, err)
			}
			box.Children = children
		} 

		boxes = append(boxes, box)
		offset += boxSize
	}

	return boxes, nil
}

func (r *Reader) readHeader() (BoxHeader, error) {
	var h BoxHeader
	if err := binary.Read(r.file, binary.BigEndian, &h.Size); err != nil {
		return h, err
	}
	if err := binary.Read(r.file, binary.BigEndian, &h.Type); err != nil {
		return h, err
	}
	return h, nil
}

func isContainer(t string) bool {
	switch t {
	case "moov", "trak", "mdia", "minf", "dinf", "stbl", "udta":
		return true
	}
	return false
}
