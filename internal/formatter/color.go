package formatter

import (
	"fmt"
	"strings"
)

const (
	Red     = "31"
	Magenta = "35"
	Cyan    = "36"
	Bold    = "01"
)

func ApplyANSI(text string, ANSICodes ...string) string {
	return fmt.Sprintf("\x1b[%vm\x1b[K%s\x1b[m\x1b[K", strings.Join(ANSICodes, ";"), text)
}

func ColorRanges(line string, ranges [][]int) string {
	// (no color) From start to before first match
	formattedLine := line[0:ranges[0][0]]

	for i := 0; i < len(ranges); i++ {
		// (blue) From first match char to last char
		formattedLine += ApplyANSI(line[ranges[i][0]:ranges[i][1]], Bold, Red)

		// (no color) If there is more matches: from after last char of prev match to first char of next match
		if i < len(ranges)-1 {
			formattedLine += line[ranges[i][1]:ranges[i+1][0]]
		}
	}

	// (no color) From last char of last match to the end of the line
	formattedLine += line[ranges[len(ranges)-1][1]:]

	return formattedLine
}
