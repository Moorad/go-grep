package fileparser

import (
	"fmt"
	"os"
)

func Parse(file_path string) ([]byte, error) {
	data, err := os.ReadFile(file_path)

	if err != nil {
		return nil, fmt.Errorf("failed to read file\n%v", err)
	}

	return data, nil
}
