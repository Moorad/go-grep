package main

import (
	"os/exec"
	"testing"
)

func grepRun(args ...string) (string, error) {
	cmd := exec.Command("grep", args...)

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

// grep must be installed for tests to work
func TestCheckGrepInstalled(t *testing.T) {
	_, err := grepRun("--version")

	if err != nil {
		t.Fatalf("The 'grep' command is not installed")
	}
}

func TestSingleLine(t *testing.T) {
	var args = []string{"Hello", "./test_files/one-line.txt"}

	stdout, err := grepRun(args...)

	if err != nil {
		t.Error(err)
	}

	results, err := Run(args)

	if err != nil {
		t.Error(err)
	}

	if stdout != results {
		t.Errorf("Output mismatch, exepected %v but found %v", stdout, results)
	}

}
