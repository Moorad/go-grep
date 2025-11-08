package formatter

import (
	"fmt"
	"strings"
)

func PrintMatchedFileName(strBuilder *strings.Builder, filename string) {
	strBuilder.WriteString(
		fmt.Sprintf("%v%v", ApplyANSI(filename, Magenta), ApplyANSI(":", Cyan)))
}
