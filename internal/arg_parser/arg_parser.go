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

func Parse() (ParsedArguments, error) {
	iFlag := registerBoolFlag("i", false, "ignore case sensitivity in patterns and data")

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

func registerBoolFlag(name string, value bool, usage string) *bool {
	if flag.Lookup(name) == nil {
		return flag.Bool(name, value, usage)
	}

	val := flag.Lookup(name).Value.(flag.Getter).Get().(bool)
	return &val
}
