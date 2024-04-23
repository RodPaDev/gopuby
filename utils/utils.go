package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func EnsureDir(dirName string) error {
	// Check if the directory exists
	_, err := os.Stat(dirName)
	if err == nil {
		// Directory already exists
		return nil
	}

	// Check if the error is because the directory does not exist
	if os.IsNotExist(err) {
		// Directory does not exist, so try creating it
		err := os.MkdirAll(dirName, 0755) // You can adjust permissions as needed
		if err != nil {
			// Failed to create directory
			return err
		}
		log.Printf("Directory created: %s", dirName)
		return nil
	}

	// Some other error occurred when trying to access the directory
	return err
}

func GetAbsPath(path string) string {
	absolutePath, err := filepath.Abs(path)
	if err != nil {
		log.Fatalf("error getting absolute path: %s", err)
	}
	return absolutePath
}

func CopyExec(src string, dst string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// Windows command
		cmd = exec.Command("cmd", "/C", "copy", "/Y", src, dst)
	} else {
		// Unix-like OS command (Linux, MacOS)
		cmd = exec.Command("cp", src, dst)
	}

	// Run the command and capture any errors.
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to copy file: %v", err)
	}
	return nil
}

func MoveExec(src string, dst string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		// Windows command
		cmd = exec.Command("cmd", "/C", "move", "/Y", src, dst)
	} else {
		// Unix-like OS command (Linux, MacOS)
		cmd = exec.Command("mv", src, dst)
	}

	// Run the command and capture any errors.
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to move file: %v", err)
	}
	return nil
}
