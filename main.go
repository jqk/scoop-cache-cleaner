package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

func main() {
	showVersion()

	if shouldShowHelp() {
		showHelp()
		return
	}

	scoopPath, err := getScoopPath(os.Args[1])
	if err != nil {
		showError(err)
		return
	}

	showCleanStart(scoopPath)
	r, err := CleanScoopCache(scoopPath, showCleaningItem)
	showCleanResult(r, err)
}

func shouldShowHelp() bool {
	return len(os.Args) != 2
}

func showVersion() {
	fmt.Println()
	fmt.Println("Copyright (c) 1999-2023 Not a dream Co., Ltd.")
	fmt.Println("scoop cache cleaner (scc) 1.0.0, 2023-01-25")
	fmt.Println()
}

func showHelp() {
	color.Set(color.LightYellow)
	fmt.Println("Usage: ")
	fmt.Println("  scc [path/to/scoop/cache]")
	color.Reset()
	fmt.Println("      clean up the specified scoop cache directory.")
	color.Set(color.LightYellow)
	fmt.Println("  scc -e")
	color.Reset()
	fmt.Println("      clean up scoop cache directory defined in the environment.")
	fmt.Println()
	color.Set(color.LightYellow)
	fmt.Println("  all other parameters will display the above information.")
	fmt.Println()
	color.Reset()
}

func showCleanStart(scoopPath string) {
	fmt.Print("Cleaning ")
	color.Set(color.LightGreen)
	fmt.Println(scoopPath)
	fmt.Println()
}

var count = 1

func showCleaningItem(pack *PackageInfo) {
	fmt.Printf("%4d ", count)
	count++
	fmt.Println(pack.Name, pack.Version)
}

func showCleanResult(result *CleanResult, err error) {
	if err != nil {
		showError(err)
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

	if result.CleanCount > 0 {
		color.Reset()

		if result.CleanCount == 1 {
			fmt.Print("Cleaned file has been moved to ")
		} else {
			fmt.Print("Cleaned files have been moved to ")
		}

		color.Set(color.LightGreen)
		fmt.Println(result.BackupPath)
		fmt.Println()
	}

	color.Reset()
}

func showError(err error) {
	color.Set(color.Red)
	fmt.Println("---------- Error! ----------")
	fmt.Println(err)
	color.Reset()
}

// getScoopPath gets the formal path string from the command parameter
// or environment variable.
func getScoopPath(param string) (string, error) {
	if param == "-e" {
		scoop := os.Getenv("SCOOP")
		if scoop == "" {
			return "", fmt.Errorf("environment variable SCOOP not found")
		}

		return JoinFileName(scoop, "cache")
	}

	return FormatFileName(param)
}
