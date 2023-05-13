package cmd_test

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/taylormonacelli/limewing/cmd"
)

func TestArgTest1Command(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("../limewing", "test1")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed with error: %v", err)
	}

	if stdout.String() != "Hello, test1!\n" {
		t.Errorf("Unexpected output. Expected 'You entered: hello\\n', but got '%s'", stdout.String())
	}

	if stderr.String() != "" {
		t.Errorf("Unexpected error output: %s", stderr.String())
	}
}

func TestArgTestCommand(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("../limewing", "test")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed with error: %v", err)
	}

	if stdout.String() != "test called\n" {
		t.Errorf("Unexpected output. Expected 'You entered: hello\\n', but got '%s'", stdout.String())
	}

	if stderr.String() != "" {
		t.Errorf("Unexpected error output: %s", stderr.String())
	}
}

func TestNoArgCommand(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command("../limewing")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		t.Fatalf("Command failed with error: %v", err)
	}

	if stdout.String() != "" {
		t.Errorf("Unexpected output. Expected 'You entered: hello\\n', but got '%s'", stdout.String())
	}

	if stderr.String() != "" {
		t.Errorf("Unexpected error output: %s", stderr.String())
	}
}

func _TestHelpCommand(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// Capture the output of `cmd.RootCmd.Help()` to a string
	helpBuf := new(bytes.Buffer)
	cmd.RootCmd.SetOut(helpBuf)
	cmd.RootCmd.Help()
	helpOutput := strings.TrimSpace(helpBuf.String())

	// Write both output strings to files for inspection
	helpFile, err := os.Create("f1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer helpFile.Close()
	_, err = helpFile.WriteString(helpOutput)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Run `../limewing --help` and capture its output to a string
	cmd1 := exec.Command("../limewing", "--help")
	cmd1.Stdout = &stdout
	cmd1.Stderr = &stderr
	err = cmd1.Run()
	if err != nil {
		t.Fatalf("Command failed with error: %v", err)
	}
	out := strings.TrimSpace(stdout.String())

	outFile, err := os.Create("f2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outFile.Close()
	_, err = outFile.WriteString(out)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Compare the output strings and fail the test if they don't match
	if out != helpOutput {
		t.Errorf("Unexpected output. Expected: '%s', but got '%s'",
			helpOutput, out)
	}

	if stderr.String() != "" {
		t.Errorf("Unexpected error output: %s", stderr.String())
	}
}
