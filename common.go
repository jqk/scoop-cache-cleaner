package main

import (
	"fmt"
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

type Size interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

var kb float64 = 1024
var mb = kb * kb
var gb = mb * kb
var tb = gb * kb

// FormatSize convert size to string with unit.
func FormatSize[T Size](size T) string {
	var value = float64(size)

	if value < kb {
		return fmt.Sprintf("%.0f bytes", value)
	} else if value < mb {
		return fmt.Sprintf("%.2f KB", value/kb)
	} else if value < gb {
		return fmt.Sprintf("%.2f MB", value/mb)
	} else if value < tb {
		return fmt.Sprintf("%.2f GB", value/gb)
	} else {
		return fmt.Sprintf("%.2f TB", value/tb)
	}
}
