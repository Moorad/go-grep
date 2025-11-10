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
	var fileCounter = make(map[int]int)
	var output strings.Builder

	for _, result := range results {
		_, ok := fileCounter[result.FileId]

		if !ok {
			fileCounter[result.FileId] = 0
		}

		if result.Line != "" {
			fileCounter[result.FileId] += len(strings.Split(result.Line, "\n"))
		}
	}

	// Looping through args.Files instead of fileCounter because fileCounter is not sorted by insertion
	for i, file := range args.Files {
		count := fileCounter[i]

		if len(args.Files) > 1 {
			formatter.PrintMatchedFileName(&output, file)
		}
		output.WriteString(fmt.Sprint(count))
		output.WriteString("\n")

	}

	return output.String()
}
