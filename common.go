package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// FormatPath returns the formal path string.
func FormatPath(path string) (string, error) {
	if p, err := filepath.Abs(path); err != nil {
		return "", err
	} else {
		return p, nil
	}
}

// CheckPathExists checks whether the provided path exists.
func CheckPathExists(path string) error {
	file, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("the path does not exist")
	}

	if !file.IsDir() {
		return fmt.Errorf("the path is a file")
	}

	return nil
}
