package fileparser

import (
	"bufio"
	"os"
)

func Parse(file_path string) (*bufio.Scanner, *os.File, error) {
	file, err := os.Open(file_path)

	if err != nil {
		return nil, nil, err
	}

	scanner := bufio.NewScanner(file)

	return scanner, file, nil
}
