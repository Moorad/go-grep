package argparser

import (
	"flag"
	"fmt"
)

type ParsedArguments struct {
	Pattern         string
	Files           []string
	CaseInsensitive bool
}

var iFlag = flag.Bool("i", false, "ignore case sensitivity in patterns and data")

func Parse() (ParsedArguments, error) {
	flag.Parse()

	// Ignore first argument (program file path)
	nonFlagArgs := flag.Args()

	if len(nonFlagArgs) < 2 {
		return ParsedArguments{}, fmt.Errorf("expected at least 2 arguments but %v was recieved\ngo-grep [pattern] [file_path]", len(nonFlagArgs))
	}

	pattern := nonFlagArgs[0]
	files := nonFlagArgs[1:]

	parseArgs := ParsedArguments{
		Files:           files,
		Pattern:         pattern,
		CaseInsensitive: *iFlag,
	}

	return parseArgs, nil
}
