package argparser

import (
	"testing"
)

func TestValidArguments(t *testing.T) {
	args := []string{"test", "file.txt"}
	parsedArguments := ParseArguments(args)

	if parsedArguments.Pattern != "test" || len(parsedArguments.Files) != 1 || parsedArguments.Files[0] != "file.txt" {
		t.Fail()
	}
}

func TestMissingSingleArgument(t *testing.T) {
	args := []string{"test"}
	parsedArguments := ParseArguments(args)

	if parsedArguments.Pattern != "test" || len(parsedArguments.Files) != 0 {
		t.Fail()
	}

}

func TestMissingAllArguments(t *testing.T) {
	args := []string{}
	parsedArguments := ParseArguments(args)

	if parsedArguments.Pattern != "" || len(parsedArguments.Files) != 0 {
		t.Fail()
	}
}

func TestMultipleFiles(t *testing.T) {
	args := []string{"world", "file1.txt", "file2.txt"}
	parsedArguments := ParseArguments(args)

	if len(parsedArguments.Files) != 2 || parsedArguments.Files[0] != "file1.txt" || parsedArguments.Files[1] != "file2.txt" {
		t.Fail()
	}
}
