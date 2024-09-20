package main

import (
	"fmt"

	argparser "github.com/Moorad/go-grep/internal/arg_parser"
)

func main() {
	args, err := argparser.Parse()

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", args)
}
