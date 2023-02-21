package main

import (
	"os"
	"path"
	"path/filepath"
)

// FormatFileName returns the formal file name string.
//
// the fileName could be a path.
func FormatFileName(fileName string) (string, error) {
	if p, err := filepath.Abs(fileName); err != nil {
		return "", err
	} else {
		return p, nil
	}
}

// FileExists checks whether the specified file exists.
//
// the fileName could be a path.
func FileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

// JoinFileName returns formal file name constructed by parameters.
func JoinFileName(elem ...string) (string, error) {
	s := path.Join(elem...)
	return FormatFileName(s)
}
