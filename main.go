package main

import (
	"fmt"
	"os"

	argparser "github.com/Moorad/go-grep/internal/argparser"

	textmatcher "github.com/Moorad/go-grep/internal/textmatcher"
	"github.com/Moorad/go-grep/internal/workers"
	"github.com/spf13/cobra"
)

const workerPoolSize = 5

var rootCmd = &cobra.Command{
	Use:           "gogrep [pattern] [...files]",
	Short:         "A simple clone of grep written in go",
	Args:          cobra.MinimumNArgs(2),
	SilenceErrors: true,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		return nil
	},
	RunE: runGrep,
}

func init() {
	argparser.ParsedOptions = argparser.ParseOptions(rootCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func runGrep(cmd *cobra.Command, args []string) error {
	parsedArguments := argparser.ParseArguments(args)

	numOfFiles := len(parsedArguments.Files)

	files := make(chan string, numOfFiles)
	results := make(chan textmatcher.MatchResult, numOfFiles)

	errorGroup := workers.Dispatch(workerPoolSize, parsedArguments.Pattern, files, results)

	for i := 0; i < numOfFiles; i++ {
		files <- parsedArguments.Files[i]
	}

	close(files)

	err := workers.Wait(errorGroup)

	if err != nil {
		return err
	}

	matchResults := []textmatcher.MatchResult{}
	for i := 0; i < numOfFiles; i++ {
		matchResults = append(matchResults, <-results)
	}

	noMatches := true
	for _, result := range matchResults {
		if result.Line != "" {
			noMatches = false
			break
		}
	}

	if noMatches {
		return fmt.Errorf("")

	}

	var output string

	if argparser.ParsedOptions.Count {
		output = workers.CountMatches(parsedArguments, matchResults)
	} else {
		output = workers.CollectMatches(parsedArguments, matchResults)
	}

	fmt.Fprintf(cmd.OutOrStdout(), output)
	return nil
}
