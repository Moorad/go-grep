package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func grepRun(args ...string) (string, error) {
	cmd := exec.Command("grep", append([]string{"--color=always"}, args...)...)

	cmd.Env = append(cmd.Env, "GREP_COLOR=01;34")

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func compareGrepAndMain(args []string) error {
	stdout, err := grepRun(args...)

	if err != nil {
		return err
	}

	results, err := Run(args)

	if err != nil {
		return err
	}

	if stdout != results {
		return fmt.Errorf("Output mismatch\nExpected:\n%v\nFound:\n%v", stdout, results)
	}

	return nil
}

// grep must be installed for tests to work
func TestCheckGrepInstalled(t *testing.T) {
	_, err := grepRun("--version")

	if err != nil {
		t.Fatalf("The 'grep' command is not installed")
	}
}

func TestSingleLineFile(t *testing.T) {
	var args = []string{"Hello", "./test_files/one-line.txt"}

	err := compareGrepAndMain(args)

	if err != nil {
		t.Error(err)
	}
}

func TestMultiLineFile(t *testing.T) {
	var args = []string{"blazing", "./test_files/twinkle.txt"}

	err := compareGrepAndMain(args)

	if err != nil {
		t.Error(err)
	}
}

func TestMultipleMatches(t *testing.T) {
	var args = []string{"twinkle", "./test_files/twinkle.txt"}

	err := compareGrepAndMain(args)

	if err != nil {
		t.Error(err)
	}
}

func TestMultipleFiles(t *testing.T) {
	var args = []string{"world", "./test_files/one-line.txt", "./test_files/twinkle.txt"}

	err := compareGrepAndMain(args)

	if err != nil {
		t.Error(err)
	}
}

func TestDuplicateFiles(t *testing.T) {
	var args = []string{"Hello", "./test_files/one-line.txt", "./test_files/one-line.txt"}

	err := compareGrepAndMain(args)

	if err != nil {
		t.Error(err)
	}
}
