package main_test

import (
	"os/exec"
	"testing"
)

// grep must be installed for tests to work
func TestCheckGrepInstalled(t *testing.T) {
	err := exec.Command("grep", "--version").Run()

	if err != nil {
		t.Fatalf("The 'grep' command is not installed")
	}
}
