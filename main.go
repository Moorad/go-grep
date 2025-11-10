package main

import (
	"fmt"
	"os"

	"github.com/Moorad/go-grep/cmd"
	argparser "github.com/Moorad/go-grep/internal/argparser"
)

var rootCmd = cmd.NewRootCmd()

func init() {
	argparser.ParsedOptions = argparser.ParseOptions(rootCmd)
}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
