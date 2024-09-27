package textmatcher

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"

	"github.com/Moorad/go-grep/internal/formatter"
)

type MatchResult struct {
	File string
	Line string
}

func Match(scanner *bufio.Scanner, file string, pattern string) (MatchResult, error) {
	var matches = []string{}

	for scanner.Scan() {
		line := scanner.Text()
		matchIndices, err := findLineMatches(line, pattern)

		if err != nil {
			return MatchResult{}, err
		}

		if matchIndices != nil {
			matches = append(matches, colorIndices(matchIndices, line))
		}
	}

	return MatchResult{
		File: file,
		Line: strings.Join(matches, "\n"),
	}, nil
}

func findLineMatches(line string, pattern string) ([][]int, error) {
	regex, err := regexp.Compile(pattern)

	if err != nil {
		return nil, fmt.Errorf("failed to parse regex pattern\n%v", err)
	}

	matchedText := regex.FindAllIndex([]byte(line), -1)

	return matchedText, nil
}

func colorIndices(indicies [][]int, line string) string {
	// (no color) From start to before first match
	formattedLine := line[0:indicies[0][0]]

	for i := 0; i < len(indicies); i++ {
		// (blue) From first match char to last char
		formattedLine += formatter.ApplyANSI(line[indicies[i][0]:indicies[i][1]], formatter.Bold, formatter.Blue)

		// (no color) If there is more matches: from after last char of prev match to first char of next match
		if i < len(indicies)-1 {
			formattedLine += line[indicies[i][1]:indicies[i+1][0]]
		}
	}

	// (no color) From last char of last match to the end of the line
	formattedLine += line[indicies[len(indicies)-1][1]:]

	return formattedLine
}
