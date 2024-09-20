package fileparser

import (
	"fmt"
	"os"
)

func Parse(file_path string) (string, error) {
	data, err := os.ReadFile(file_path)

	if err != nil {
		return "", fmt.Errorf("failed to read file\n%v", err)
	}

	return string(data), nil
}
