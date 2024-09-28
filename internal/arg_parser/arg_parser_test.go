package argparser

import (
	"os"
	"testing"
)

func TestValidArgs(t *testing.T) {
	os.Args = []string{"gogrep", "test", "file.txt"}
	parsedArgs, err := Parse()

	if err != nil || parsedArgs.Arguments.Pattern != "test" || len(parsedArgs.Arguments.Files) != 1 || parsedArgs.Arguments.Files[0] != "file.txt" {
		t.Error(err)
	}
}

func TestMissingArgs(t *testing.T) {
	os.Args = []string{"gogrep", "test"}
	missingFileArgs, err := Parse()

	if err == nil || missingFileArgs.Arguments.Pattern != "" || len(missingFileArgs.Arguments.Files) != 0 {
		t.Error(err)
	}

	os.Args = []string{"gogrep", "file.txt"}
	missingPatternArgs, err := Parse()

	if err == nil || missingPatternArgs.Arguments.Pattern != "" || len(missingFileArgs.Arguments.Files) != 0 {
		t.Error(err)
	}

	os.Args = []string{"gogrep"}
	noArgs, err := Parse()

	if err == nil || noArgs.Arguments.Pattern != "" || len(noArgs.Arguments.Files) != 0 {
		t.Error(err)
	}
}

func TestMultipleFiles(t *testing.T) {
	os.Args = []string{"gogrep", "world", "file1.txt", "file2.txt"}
	multipleFilesArgs, err := Parse()

	if err != nil || len(multipleFilesArgs.Arguments.Files) != 2 || multipleFilesArgs.Arguments.Files[0] != "file1.txt" || multipleFilesArgs.Arguments.Files[1] != "file2.txt" {
		t.Error(err)
	}
}

func TestCaseInsensitiveArgument(t *testing.T) {
	os.Args = []string{"gogrep", "-i", "twinkle", "test_files/twinkle.txt"}
	parseArgs, err := Parse()

	if err != nil || parseArgs.Options.IgnoreCase == false {
		t.Error(err)
	}
}
