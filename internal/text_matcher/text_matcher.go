package textmatcher

import (
	"fmt"
	"regexp"
)

func MatchSingleLine(line *[]byte, pattern string) (bool, error) {
	regex, err := regexp.Compile(pattern)

	if err != nil {
		return false, fmt.Errorf("failed to parse regex pattern\n%v", err)
	}

	return regex.Match(*line), nil
}
