package collector

import (
	"fmt"
	"strings"

	argparser "github.com/Moorad/go-grep/internal/arg_parser"
	formatter "github.com/Moorad/go-grep/internal/formatter"
	textmatcher "github.com/Moorad/go-grep/internal/text_matcher"
)

func CollectMatches(args argparser.Arguments, results []textmatcher.MatchResult) string {
	var output strings.Builder

	for i := 0; i < len(args.Files); i++ {
		for j := 0; j < len(results); j++ {
			if results[j].File == args.Files[i] && results[j].Line != "" {
				if len(args.Files) > 1 {
					output.WriteString(
						fmt.Sprintf("%v%v", formatter.ApplyANSI(results[j].File, formatter.Magenta), formatter.ApplyANSI(":", formatter.Cyan)))
				}
				output.WriteString(results[j].Line)
				output.WriteString("\n")
				break
			}
		}
	}

	return output.String()
}
