package isobmff

import (
	"fmt"
	"strings"
)

// PrintStructure prints the box hierarchy to stdout.
func PrintStructure(boxes []*Box, indentLevel int) {
	indent := strings.Repeat("  ", indentLevel)
	for _, box := range boxes {
		name := string(box.Header.Type[:])
		size := box.Header.Size
		// If size was 1, we don't have the real size in header struct in this simple model,
		// but for display purposes the 32-bit field is sufficient or we could store real size.
		
		fmt.Printf("%s[%s] @ %d (size: %d)\n", indent, name, box.Offset, size)
		
		if len(box.Children) > 0 {
			PrintStructure(box.Children, indentLevel+1)
		}
	}
}

