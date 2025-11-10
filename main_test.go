package main

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"testing"
)

func grepRun(args ...string) (string, error) {
	cmd := exec.Command("grep", append([]string{"--color=always"}, args...)...)

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func compareParityWithGrep(args []string) error {
	grepOutput, grepErr := grepRun(args...)

	outputBuff := bytes.NewBufferString("")

	rootCmd.SetOut(outputBuff)
	rootCmd.SetErr(outputBuff)
	rootCmd.SetArgs(args)
	gogrepErr := rootCmd.Execute()

	output, _ := io.ReadAll(outputBuff)

	gogrepOutput := string(output)

	if grepErr != nil && gogrepErr == nil {
		return fmt.Errorf("Exit code mismatch\nExpected exit code = 1 but found exit code = 0 with output:\n%v", gogrepOutput)
	}

	if grepOutput != gogrepOutput {
		return fmt.Errorf("Output mismatch\nExpected:\n%v\nFound:\n%v", grepOutput, gogrepOutput)
	}

	return nil
}

// grep must be installed for tests to work
func TestCheckGrepInstalled(t *testing.T) {
	out, err := grepRun("--version")

	if strings.Split(out, "\n")[0] != "grep (GNU grep) 3.11" {
		fmt.Println("WARNING: Current version of grep is NOT 3.11, tests may fail.")
	}

	if err != nil {
		t.Fatalf("The 'grep' command is not installed")
	}
}

func TestSingleLineFile(t *testing.T) {
	var args = []string{"Hello", "./test_files/one-line.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestNoMatches(t *testing.T) {
	var args = []string{"Hello", "./test_files/twinkle.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestMultiLineFile(t *testing.T) {
	var args = []string{"blazing", "./test_files/twinkle.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestMultipleMatches(t *testing.T) {
	var args = []string{"twinkle", "./test_files/twinkle.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestMultipleFiles(t *testing.T) {
	var args = []string{"world", "./test_files/one-line.txt", "./test_files/twinkle.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestMultipleFilesReverse(t *testing.T) {
	var args = []string{"world", "./test_files/one-line.txt", "./test_files/twinkle.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestDuplicateFiles(t *testing.T) {
	var args = []string{"Hello", "./test_files/one-line.txt", "./test_files/one-line.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestSomeFilesHaveMatches(t *testing.T) {
	var args = []string{"Hello", "./test_files/one-line.txt", "./test_files/twinkle.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestCaseInsensitive(t *testing.T) {
	var args = []string{"-i", "twinkle", "./test_files/twinkle.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestCountSingleFileMatches(t *testing.T) {
	var args = []string{"-c", "twinkle", "./test_files/twinkle.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestCountMultiFileMatches(t *testing.T) {
	var args = []string{"-c", "world", "./test_files/one-line.txt", "./test_files/twinkle.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestCountMultiFileWithSingleFileMatch(t *testing.T) {
	var args = []string{"-c", "twinkle", "./test_files/one-line.txt", "./test_files/twinkle.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestCountMultiFileNoMatch(t *testing.T) {
	var args = []string{"-c", "x", "./test_files/one-line.txt", "./test_files/twinkle.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}

func TestCountDuplicateFiles(t *testing.T) {
	var args = []string{"-c", "twinkle", "./test_files/twinkle.txt", "./test_files/twinkle.txt"}

	err := compareParityWithGrep(args)

	if err != nil {
		t.Error(err)
	}
}
