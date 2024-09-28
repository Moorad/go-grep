package main

import (
	"fmt"

	argparser "github.com/Moorad/go-grep/internal/arg_parser"
	"github.com/Moorad/go-grep/internal/collector"
	fileparser "github.com/Moorad/go-grep/internal/file_parser"
	textmatcher "github.com/Moorad/go-grep/internal/text_matcher"
	"golang.org/x/sync/errgroup"
)

const numOfWorkers = 5

func main() {
	results, err := Run()

	if err != nil {
		panic(err)
	}

	fmt.Print(results)
}

func Run() (string, error) {
	parsedArgs, err := argparser.Parse()
	eg := new(errgroup.Group)

	if err != nil {
		return "", err
	}

	numOfFiles := len(parsedArgs.Files)

	files := make(chan string, numOfFiles)
	results := make(chan textmatcher.MatchResult, numOfFiles)

	for i := 0; i < numOfWorkers; i++ {
		eg.Go(func() error {
			return worker(parsedArgs.Pattern, files, results)
		})
	}

	for i := 0; i < numOfFiles; i++ {
		files <- parsedArgs.Files[i]
	}

	close(files)

	if err := eg.Wait(); err != nil {
		return "", nil
	}

	matchResults := make([]textmatcher.MatchResult, numOfFiles)
	for i := 0; i < numOfFiles; i++ {
		matchResults = append(matchResults, <-results)

	}

	output := collector.CollectMatches(parsedArgs, matchResults)

	return output, nil
}

func worker(pattern string, files chan string, results chan textmatcher.MatchResult) error {
	for filePath := range files {
		scanner, file, err := fileparser.Parse(filePath)

		if err != nil {
			return err
		}

		defer file.Close()

		matchedLines, err := textmatcher.Match(scanner, filePath, pattern)

		if err != nil {
			return err
		}

		results <- matchedLines
	}

	return nil
}
