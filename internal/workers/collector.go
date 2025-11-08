package workers

import (
	"fmt"
	"strings"

	argparser "github.com/Moorad/go-grep/internal/argparser"
	formatter "github.com/Moorad/go-grep/internal/formatter"
	textmatcher "github.com/Moorad/go-grep/internal/textmatcher"
)

func CollectMatches(args argparser.Arguments, results []textmatcher.MatchResult) string {
	var output strings.Builder

	for _, currentFile := range args.Files {
		for _, result := range results {
			if result.File == currentFile && result.Line != "" {
				if len(args.Files) > 1 {
					formatter.PrintMatchedFileName(&output, result.File)
				}
				output.WriteString(result.Line)
				output.WriteString("\n")
				break
			}
		}
	}

	return output.String()
}

func CountMatches(args argparser.Arguments, results []textmatcher.MatchResult) string {
	var fileCounter = make(map[string]int)
	var output strings.Builder

	for i := 0; i < len(results); i++ {
		_, ok := fileCounter[results[i].File]

		if !ok {
			fileCounter[results[i].File] = 0
		}

		if results[i].Line != "" {
			fileCounter[results[i].File] += len(strings.Split(results[i].Line, "\n"))
		}
	}

	// Looping through args.Files instead of fileCounter because fileCounter is not sorted by insertion
	for _, currentFile := range args.Files {
		count := fileCounter[currentFile]

		if len(args.Files) > 1 {
			formatter.PrintMatchedFileName(&output, currentFile)
		}
		output.WriteString(fmt.Sprint(count))
		output.WriteString("\n")

	}

	return output.String()
}
