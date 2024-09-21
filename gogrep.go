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

	scanner, file, err := fileparser.Parse(parsedArgs.FilePath)

	if err != nil {
		return "", err
	}

	defer file.Close()

	matchedLine, err := textmatcher.Match(scanner, file, parsedArgs.Pattern)

	if err != nil {
		return "", err
	}

	return matchedLine, nil
}
