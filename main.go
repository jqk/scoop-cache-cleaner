package main

import (
	"fmt"
	"os"
	"path"
	"strings"
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
