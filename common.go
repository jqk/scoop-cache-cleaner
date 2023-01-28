package main

import (
	"path/filepath"
)

func formatPath(path string) (string, error) {
	if p, err := filepath.Abs(path); err != nil {
		return "", err
	} else {
		return p, nil
	}
}
