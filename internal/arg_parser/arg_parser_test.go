package argparser

import "testing"

func TestValidArgs(t *testing.T) {
	parsedArgs, err := Parse([]string{"test", "file.txt"})

	if err != nil || parsedArgs.Pattern != "test" || len(parsedArgs.Files) != 1 || parsedArgs.Files[0] != "file.txt" {
		t.Error(err)
	}
}

func TestMissingArgs(t *testing.T) {
	missingFileArgs, err := Parse([]string{"test"})

	if err == nil || missingFileArgs.Pattern != "" || len(missingFileArgs.Files) != 0 {
		t.Error(err)
	}

	missingPatternArgs, err := Parse([]string{"file.txt"})

	if err == nil || missingPatternArgs.Pattern != "" || len(missingFileArgs.Files) != 0 {
		t.Error(err)
	}

	noArgs, err := Parse(nil)

	if err == nil || noArgs.Pattern != "" || len(noArgs.Files) != 0 {
		t.Error(err)
	}
}

func TestMultipleFiles(t *testing.T) {
	multipleFilesArgs, err := Parse([]string{"world", "file1.txt", "file2.txt"})

	if err != nil || len(multipleFilesArgs.Files) != 2 || multipleFilesArgs.Files[0] != "file1.txt" || multipleFilesArgs.Files[1] != "file2.txt" {
		t.Error(err)
	}
}
