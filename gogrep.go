package main

import (
	"fmt"

	argparser "github.com/Moorad/go-grep/internal/arg_parser"
	fileparser "github.com/Moorad/go-grep/internal/file_parser"
	textmatcher "github.com/Moorad/go-grep/internal/text_matcher"
)

func main() {
	args, err := argparser.Parse()

	if err != nil {
		panic(err)
	}

	data, err := fileparser.Parse(args.FilePath)

	if err != nil {
		panic(err)
	}

	hasMatch, err := textmatcher.MatchSingleLine(data, args.Pattern)

	if err != nil {
		panic(err)
	}

	if hasMatch {
		fmt.Println(data)
	}

}
