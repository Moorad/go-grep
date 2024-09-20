package argparser

import (
	"fmt"
	"os"
)

type ParsedArguments struct {
	pattern  string
	filePath string
}

func Parse() (ParsedArguments, error) {
	if len(os.Args) < 3 {
		return ParsedArguments{}, fmt.Errorf("expected at least 2 arguments but received %v\ngo-grep [pattern] [file_path]", len(os.Args)-1)
	}

	pattern := os.Args[len(os.Args)-2]
	filename := os.Args[len(os.Args)-1]

	return ParsedArguments{
		filePath: filename,
		pattern:  pattern,
	}, nil
}
