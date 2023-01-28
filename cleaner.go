package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

// CleanResult stores various count during processing.
type CleanResult struct {
	FileCount     int
	CleanCount    int
	SoftwareCount int
}

// checkPath make sure the provided scoop path is exist.
func checkPath(path string) error {
	file, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("scoop cache path is not exist")
	}

	if !file.IsDir() {
		return fmt.Errorf("scoop cache path should not be a file")
	}

	return nil
}

func prepareBackupPath(scoopPath string) (string, error) {
	s := path.Join(scoopPath, time.Now().Format("bak_2006-01-02T15-04-05"))

	if err := os.Mkdir(s, 0777); err != nil {
		return "", err
	}

	return formatPath(s)
}

func cleanScoopCache(scoopPath string, backupPath string) (*CleanResult, error) {
	if err := checkPath(scoopPath); err != nil {
		return nil, err
	}

	f, err := os.Open(scoopPath)
	if err != nil {
		return nil, err
	}

	files, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	result := CleanResult{0, 0, 0}
	count := len(files)

	name := ""
	version := ""

	for i := count - 1; i >= 0; i-- {
		file := files[i]

		if !file.IsDir() {
			result.FileCount++
			software := strings.Split(file.Name(), "#")

			if len(software) != 3 {
				continue
			}

			n := software[0]
			v := software[1]

			if n != name {
				result.SoftwareCount++
				name = n
				version = v
			} else if v != version {
				result.CleanCount++

				old := path.Join(scoopPath, file.Name())
				new := path.Join(backupPath, file.Name())

				if err := os.Rename(old, new); err != nil {
					return &result, err
				}

				fmt.Println(name, version)
			}
		}
	}

	if result.CleanCount == 0 {
		os.Remove(backupPath)
	}

	return &result, nil
}
