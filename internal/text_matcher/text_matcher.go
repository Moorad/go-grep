package textmatcher

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func Match(scanner *bufio.Scanner, file *os.File, pattern string) (string, error) {
	for scanner.Scan() {
		line := scanner.Text()
		hasMatch, err := MatchSingleLine(line, pattern)

		if err != nil {
			return "", err
		}

		if hasMatch {
			return fmt.Sprintln(line), nil
		}
	}

	return "", nil
}

func MatchSingleLine(line string, pattern string) (bool, error) {
	regex, err := regexp.Compile(pattern)

	if err != nil {
		return false, fmt.Errorf("failed to parse regex pattern\n%v", err)
	}

	return regex.Match([]byte(line)), nil
}
