package workers

import (
	argparser "github.com/Moorad/go-grep/internal/argparser"
	"github.com/Moorad/go-grep/internal/fileparser"
	textmatcher "github.com/Moorad/go-grep/internal/textmatcher"
	"golang.org/x/sync/errgroup"
)

func Dispatch(workerPoolSize int, pattern string, files chan string, results chan textmatcher.MatchResult) *errgroup.Group {
	eg := new(errgroup.Group)

	eg.Go(func() error {
		return worker(pattern, files, results)
	})

	return eg

}

func Wait(eg *errgroup.Group) error {

	if err := eg.Wait(); err != nil {
		return err
	}

	return nil
}

func worker(pattern string, files chan string, results chan textmatcher.MatchResult) error {
	fileId := 0
	for filePath := range files {
		scanner, file, err := fileparser.Parse(filePath)

		if err != nil {
			return err
		}

		defer file.Close()

		matchedLines, err := textmatcher.Match(scanner, fileId, filePath, pattern, argparser.ParsedOptions)

		if err != nil {
			return err
		}
		fileId++
		results <- matchedLines
	}

	return nil
}
