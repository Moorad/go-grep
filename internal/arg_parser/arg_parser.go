package argparser

import (
	"fmt"
)

type ParsedArguments struct {
	Pattern  string
	FilePath string
}

func Parse(args []string) (ParsedArguments, error) {
	if len(args) < 2 {
		return ParsedArguments{}, fmt.Errorf("expected at least 2 arguments but received %v\ngo-grep [pattern] [file_path]", len(args)-1)
	}

	pattern := args[len(args)-2]
	filename := args[len(args)-1]

	return ParsedArguments{
		FilePath: filename,
		Pattern:  pattern,
	}, nil
}
