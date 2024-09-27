package collector

import (
	"fmt"
	"strings"

	argparser "github.com/Moorad/go-grep/internal/arg_parser"
	textmatcher "github.com/Moorad/go-grep/internal/text_matcher"
)

func CollectMatches(args argparser.ParsedArguments, results []textmatcher.MatchResult) string {
	var output strings.Builder

	for i := 0; i < len(args.Files); i++ {
		for j := 0; j < len(results); j++ {
			if results[j].File == args.Files[i] {
				if len(args.Files) > 1 {
					output.WriteString(fmt.Sprintf("%v:", results[j].File))
				}
				output.WriteString(results[j].Line)
				output.WriteString("\n")
			}
		}
	}

	return output.String()
}
