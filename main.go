package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

var packageInfoFormat string

func main() {
	showVersion()

	action := parseCmdParameter()
	if action == nil {
		showHelp()
		return
	}

	var err error
	action.ScoopPath, err = GetScoopPath(action.ScoopPath)
	if err != nil {
		showError(err)
		return
	}

	showCleanStart(action)

	result, err := FindObsoletePackages(action)
	if err != nil {
		showError(err)
		return
	}

	setPackageInfoFormat(result)

	if err := CleanScoopCache(result, showCleaningItem); err != nil {
		showError(err)
		return
	}

	showCleanResult(result)
}

func parseCmdParameter() *ActionInfo {
	n := len(os.Args)

	if n == 2 {
		return NewAction(os.Args[1], "")
	} else if n == 3 {
		return NewAction(os.Args[1], os.Args[2])
	} else {
		return nil
	}
}

func showVersion() {
	fmt.Println()
	fmt.Println("Copyright (c) 1999-2023 Not a dream Co., Ltd.")
	fmt.Println("scoop cache cleaner (scc) 2.1.2, 2023-02-28")
	fmt.Println()
}

func showHelp() {
	color.Set(color.LightYellow)
	fmt.Println("Usage:")
	fmt.Println("  scc <command> [path/to/scoop/cache]")

	color.Reset()
	fmt.Println("      clean up the specified scoop cache directory.")
	fmt.Println("      if the path is omitted, it will use the path defined in the environment variable %SCOOP%.")

	color.Set(color.LightYellow)
	fmt.Println("\nCommand:")
	fmt.Print("  -l: ")

	color.Reset()
	fmt.Println(" list the obsolete packages.")

	color.Set(color.LightYellow)
	fmt.Print("  -b: ")

	color.Reset()
	fmt.Println(" backup the obsolete packages.")

	color.Set(color.LightYellow)
	fmt.Print("  -d: ")

	color.Reset()
	fmt.Println(" delete the obsolete packages.")

	color.Set(color.LightYellow)
	fmt.Println("\nall other parameters will display the above information.")
	fmt.Println()

	color.Reset()
}

func showCleanStart(action *ActionInfo) {
	if action.Action == List {
		color.Set(color.LightYellow)
		fmt.Print("List")
	} else if action.Action == Backup {
		color.Set(color.LightYellow)
		fmt.Print("Backup")
	} else { // if action.Action == Delete
		color.Set(color.LightRed)
		fmt.Print("Delete")
	}

	color.Reset()
	fmt.Print(" obsolete packages in: ")

	color.Set(color.LightGreen)
	fmt.Println(action.ScoopPath)
	fmt.Println()
}

func setPackageInfoFormat(result *CleanResult) {
	if result.CleanCount == 0 {
		return
	}

	nameLength := 0
	versionLength := 0

	for _, p := range result.CleanPackages {
		nl := len(p.Name)
		vl := len(p.Version)

		if nl > nameLength {
			nameLength = nl
		}
		if vl > versionLength {
			versionLength = vl
		}
	}

	packageInfoFormat = fmt.Sprintf("%%4d %%-%ds  %%-%ds  %%s  ", nameLength, versionLength)

	t := fmt.Sprintf("     %%-%ds  %%-%ds  %%s  %%7s\n", nameLength, versionLength)
	s := fmt.Sprintf(t, "Name", "Version", "Extension", "Size")

	color.Set(color.LightYellow)
	fmt.Println(s)
	color.Reset()
}

var count = 1
var sizeColorLimit int64 = 1024 * 1024

func showCleaningItem(pack *PackageInfo) {
	color.Reset()
	// 9 is the length of 'Extension'.
	fmt.Printf(packageInfoFormat, count, pack.Name, pack.Version, pack.FileName[len(pack.FileName)-9:])
	count++

	if pack.Size < sizeColorLimit {
		color.Set(color.Magenta)
	} else {
		color.Set(color.LightRed)
	}

	fmt.Printf("%10s\n", FormatSize(pack.Size))
}

func showCleanResult(result *CleanResult) {
	if result.CleanCount != 0 {
		fmt.Println()
	}

	color.Set(color.LightGreen)
	fmt.Println("-----------------------")
	fmt.Println("File found            :", result.FileCount)
	fmt.Println("Software found        :", result.SoftwareCount)
	fmt.Println("Obsolete Package found:", result.CleanCount)
	fmt.Print("Obsolete Package Size : ")

	color.Set(color.LightRed)
	size := FormatSize(result.CleanSize)
	fmt.Println(size)

	color.Set(color.LightGreen)
	fmt.Println("-----------------------")

	color.Reset()
	fmt.Println()

	if result.CleanCount == 0 {
		fmt.Println("Obsolete package not found.")
	} else if result.Action == Backup {
		fmt.Print("Obsolete Packages have been moved to ")

		color.Set(color.LightGreen)
		fmt.Println(result.BackupPath)
	} else if result.Action == Delete {
		fmt.Print("Free up ")

		color.Set(color.LightRed)
		fmt.Print(size)

		color.Reset()
		fmt.Println(" disk space.")
	}

	fmt.Println()
	color.Reset()
}

func showError(err error) {
	color.Set(color.LightRed)
	fmt.Println("---------- Error! ----------")
	fmt.Println(err)
	color.Reset()
}
