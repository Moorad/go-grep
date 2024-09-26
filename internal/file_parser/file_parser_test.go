package fileparser

import (
	"testing"
)

func TestValidFile(t *testing.T) {
	scanner, file, err := Parse("../../test_files/twinkle.txt")

	if err != nil || scanner == nil || file == nil {
		t.Error(err)
	}
}

func TestNonexistantFile(t *testing.T) {
	scanner, file, err := Parse("../../test_files/unknown.txt")

	if err == nil || scanner != nil || file != nil {
		t.Error(err)
	}
}
