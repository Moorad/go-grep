package textmatcher

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	argparser "github.com/Moorad/go-grep/internal/argparser"
	"github.com/Moorad/go-grep/internal/formatter"
)

type MatchResult struct {
	FileId int
	File   string
	Line   string
}

func Match(scanner *bufio.Scanner, fileId int, file string, pattern string, options *argparser.Options) (MatchResult, error) {
	var matches = []string{}

	for scanner.Scan() {
		line := scanner.Text()
		matchIndices, err := findLineMatches(line, pattern, options)

		if err != nil {
			return MatchResult{}, err
		}

		if matchIndices != nil {
			matches = append(matches, formatter.ColorRanges(line, matchIndices))
		}
	}

	return MatchResult{
		FileId: fileId,
		File:   file,
		Line:   strings.Join(matches, "\n"),
	}, nil
}

func findLineMatches(line string, pattern string, options *argparser.Options) ([][]int, error) {
	flags := ""

	if options.IgnoreCase {
		flags += "(?i)"
	}

	regex, err := regexp.Compile(flags + pattern)

	if err != nil {
		return nil, fmt.Errorf("failed to parse regex pattern\n%v", err)
	}

	matchedText := regex.FindAllIndex([]byte(line), -1)

	return matchedText, nil
}
