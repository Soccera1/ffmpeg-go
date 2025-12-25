package main

import (
	"flag"
	"fmt"
	"os"

	"ffmpeg-go/internal/isobmff"
)

func main() {
	inputFile := flag.String("i", "", "input file (MP4/ISO Base Media Format)")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Usage: ffmpeg-go -i <file.mp4>")
		os.Exit(1)
	}

	fmt.Printf("ffmpeg-go (Pure Go Implementation)\n")
	fmt.Printf("Input #0, isobmff, from '%s':\n", *inputFile)

	reader, err := isobmff.NewReader(*inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer reader.Close()

	boxes, err := reader.ReadBoxes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing boxes: %v\n", err)
		os.Exit(1)
	}

	isobmff.PrintStructure(boxes, 1)
}

