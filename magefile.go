//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Build compiles the Go project into the `bin` directory.
func Build() error {
	// Create the `bin` directory if it doesn't exist
	err := os.MkdirAll("bin", os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create bin directory: %v", err)
	}

	// Compile the binary
	cmd := exec.Command("go", "build", "-o", "./bin/app.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Building the binary...")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("build failed: %v", err)
	}

	fmt.Println("Build complete!")
	return nil
}

// Run executes the binary in the `bin` directory.
func Run() error {
	fmt.Println("Running the binary...")

	cmd := exec.Command("./bin/app.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run the binary: %v", err)
	}

	return nil
}

// Default runs the Build and Run tasks sequentially.
func Default() error {
	if err := Build(); err != nil {
		return err
	}
	return Run()
}

// Clean removes the `bin` directory and its contents.
func Clean() error {
	fmt.Println("Cleaning up...")
	return os.RemoveAll("bin")
}

// Test runs all Go tests in the project.
func Test() error {
	fmt.Println("Running tests...")
	cmd := exec.Command("go", "test", "./...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
