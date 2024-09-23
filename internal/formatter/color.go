package formatter

import (
	"fmt"
	"strings"
)

const (
	Blue = "34"
	Bold = "01"
)

func ApplyANSI(text string, ANSICodes ...string) string {
	return fmt.Sprintf("\x1b[%vm\x1b[K%s\x1b[m\x1b[K", strings.Join(ANSICodes, ";"), text)
}
