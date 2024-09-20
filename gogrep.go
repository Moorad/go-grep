package main

import (
	"fmt"

	argparser "github.com/Moorad/go-grep/internal/arg_parser"
	fileparser "github.com/Moorad/go-grep/internal/file_parser"
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

	fmt.Println(data)
}
