package main

import (
	"fmt"
	"os"

	argparser "github.com/Moorad/go-grep/internal/arg_parser"
	fileparser "github.com/Moorad/go-grep/internal/file_parser"
	textmatcher "github.com/Moorad/go-grep/internal/text_matcher"
)

func main() {
	results, err := Run(os.Args[1:])

	if err != nil {
		panic(err)
	}

	fmt.Println(results)
}

func Run(args []string) (string, error) {
	parsedArgs, err := argparser.Parse(args)

	if err != nil {
		return "", err
	}

	data, err := fileparser.Parse(parsedArgs.FilePath)

	if err != nil {
		return "", err
	}

	hasMatch, err := textmatcher.MatchSingleLine(&data, parsedArgs.Pattern)

	if err != nil {
		return "", err
	}

	if hasMatch {
		return fmt.Sprintln(string(data)), nil
	}

	return "", nil
}
