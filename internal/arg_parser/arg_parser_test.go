package argparser

import (
	"os"
	"testing"
)

func TestValidArgs(t *testing.T) {
	os.Args = []string{"gogrep", "test", "file.txt"}
	parsedArgs, err := Parse()

	if err != nil || parsedArgs.Pattern != "test" || len(parsedArgs.Files) != 1 || parsedArgs.Files[0] != "file.txt" {
		t.Error(err)
	}
}

func TestMissingArgs(t *testing.T) {
	os.Args = []string{"gogrep", "test"}
	missingFileArgs, err := Parse()

	if err == nil || missingFileArgs.Pattern != "" || len(missingFileArgs.Files) != 0 {
		t.Error(err)
	}

	os.Args = []string{"gogrep", "file.txt"}
	missingPatternArgs, err := Parse()

	if err == nil || missingPatternArgs.Pattern != "" || len(missingFileArgs.Files) != 0 {
		t.Error(err)
	}

	os.Args = []string{"gogrep"}
	noArgs, err := Parse()

	if err == nil || noArgs.Pattern != "" || len(noArgs.Files) != 0 {
		t.Error(err)
	}
}

func TestMultipleFiles(t *testing.T) {
	os.Args = []string{"gogrep", "world", "file1.txt", "file2.txt"}
	multipleFilesArgs, err := Parse()

	if err != nil || len(multipleFilesArgs.Files) != 2 || multipleFilesArgs.Files[0] != "file1.txt" || multipleFilesArgs.Files[1] != "file2.txt" {
		t.Error(err)
	}
}

func TestCaseInsensitiveArgument(t *testing.T) {
	os.Args = []string{"gogrep", "-i", "twinkle", "test_files/twinkle.txt"}
	parseArgs, err := Parse()

	if err != nil || parseArgs.CaseInsensitive == false {
		t.Error(err)
	}
}
