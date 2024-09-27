package argparser

import (
	"fmt"
)

type ParsedArguments struct {
	Pattern string
	Files   []string
}

func Parse(args []string) (ParsedArguments, error) {
	if len(args) < 2 {
		return ParsedArguments{}, fmt.Errorf("expected at least 2 arguments but %v was recieved\ngo-grep [pattern] [file_path]", len(args))
	}

	pattern := args[0]
	files := args[1:]

	parseArgs := ParsedArguments{
		Files:   files,
		Pattern: pattern,
	}

	return parseArgs, nil
}
