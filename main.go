package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	showVersion()

	if shouldShowHelp() {
		showHelp()
		return
	}

	scoopPath, err := getScoopPath(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	backupPath, err := prepareBackupPath(scoopPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	showCleanStarting(scoopPath, backupPath)

	r, err := cleanScoopCache(scoopPath, backupPath)
	showCleanResult(r, err)
}

func shouldShowHelp() bool {
	return len(os.Args) != 2
}

func showVersion() {
	fmt.Println()
	fmt.Println("Copyright (c) 1999-2023 Not a dream Co., Ltd.")
	fmt.Println("scoop cache cleaner 1.0.0, 2023-01-25")
	fmt.Println()
}

func showHelp() {
	fmt.Println("Usage: ")
	fmt.Println("  scc [path/to/scoop/cache]")
	fmt.Println("      clean specified directory.")
	fmt.Println("  scc -e")
	fmt.Println("      clean scoop cache directory defined in environment.")
	fmt.Println()
	fmt.Println("  all other parameters will show this screen.")
	fmt.Println()
}

func showCleanStarting(scoopPath string, backupPath string) {
	fmt.Println("Cleaning", scoopPath)
	fmt.Println("Outdated setup files will be moved into", backupPath)
	fmt.Println()
}

func showCleanResult(result *CleanResult, err error) {
	if err != nil {
		fmt.Println(err)
		return
	}

	if result.CleanCount != 0 {
		fmt.Println()
	}

	fmt.Println("-------------------")
	fmt.Println("File checked      :", result.FileCount)
	fmt.Println("Setup file found  :", result.SoftwareCount)
	fmt.Println("Setup file cleaned:", result.CleanCount)
	fmt.Println("-------------------")
	fmt.Println()
}

func getScoopPath(param string) (string, error) {
	if strings.HasPrefix(param, "-e") {
		scoop := os.Getenv("SCOOP")
		if scoop == "" {
			return "", fmt.Errorf("environment variable SCOOP is not found")
		}

		param = path.Join(scoop, "cache")
	}

	return formatPath(param)
}

func formatPath(path string) (string, error) {
	if p, err := filepath.Abs(path); err != nil {
		return "", err
	} else {
		return p, nil
	}
}

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
