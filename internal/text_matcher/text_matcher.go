package textmatcher

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Match(scanner *bufio.Scanner, file *os.File, pattern string) (string, error) {
	var matches = []string{}

	for scanner.Scan() {
		line := scanner.Text()
		hasMatch, err := MatchSingleLine(line, pattern)

		if err != nil {
			return "", err
		}

		if hasMatch {
			matches = append(matches, line)
		}
	}

	return strings.Join(matches, "\n"), nil
}

func MatchSingleLine(line string, pattern string) (bool, error) {
	regex, err := regexp.Compile(pattern)

	if err != nil {
		return false, fmt.Errorf("failed to parse regex pattern\n%v", err)
	}

	return regex.Match([]byte(line)), nil
}
