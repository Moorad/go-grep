package argparser

import "testing"

func TestValidArgs(t *testing.T) {
	parsedArgs, err := Parse([]string{"test", "file.txt"})

	if err != nil || parsedArgs.Pattern != "test" || parsedArgs.FilePath != "file.txt" {
		t.Error(err)
	}
}

func TestMissingArgs(t *testing.T) {
	missingFileArgs, err := Parse([]string{"test"})

	if err == nil || missingFileArgs.Pattern != "" || missingFileArgs.FilePath != "" {
		t.Error(err)
	}

	missingPatternArgs, err := Parse([]string{"file.txt"})

	if err == nil || missingPatternArgs.Pattern != "" || missingPatternArgs.FilePath != "" {
		t.Error(err)
	}

	noArgs, err := Parse(nil)

	if err == nil || noArgs.Pattern != "" || noArgs.FilePath != "" {
		t.Error(err)
	}
}
