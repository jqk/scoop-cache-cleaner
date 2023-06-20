package main

import (
	"fmt"

	"github.com/gookit/color"
)

func showVersion() {
	fmt.Println()
	fmt.Println("Copyright (c) 1999-2023 Not a dream Co., Ltd.")
	fmt.Println("scoop cache cleaner (scc) 2.2.0, 2023-06-20")
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
	fmt.Println("\nCommand (casesensitive):")
	fmt.Print("  -h, --help  : ")

	color.Reset()
	fmt.Println(" show this help.")

	color.Set(color.LightYellow)
	fmt.Print("  -l, --list  : ")

	color.Reset()
	fmt.Println(" list the obsolete packages.")

	color.Set(color.LightYellow)
	fmt.Print("  -b, --backup: ")

	color.Reset()
	fmt.Println(" backup the obsolete packages.")

	color.Set(color.LightYellow)
	fmt.Print("  -d, --delete: ")

	color.Reset()
	fmt.Println(" delete the obsolete packages.")

	color.Set(color.LightYellow)
	fmt.Print("  no argument : ")

	color.Reset()
	fmt.Println(" equal to 'scc -l' if %SCOOP% exists, otherwise show help.")

	color.Set(color.LightYellow)
	fmt.Println("\nall other parameters will display the above help information.")
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

// the length of package file extension.
const extlength = 9

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

	// 1. sequence number, right aligned.
	// 2. pacakge name,    left aligned.
	// 3. package version, left aligned.
	// 4. file extension,  right aligned.
	packageInfoFormat = fmt.Sprintf("%%4d %%-%ds  %%-%ds  %%%ds  ", nameLength, versionLength, extlength)

	t := fmt.Sprintf("     %%-%ds  %%-%ds  %%s  %%7s\n", nameLength, versionLength)
	s := fmt.Sprintf(t, "Name", "Version", "Extension", "Size")

	color.Set(color.LightYellow)
	fmt.Println(s)
	color.Reset()
}

func getPackageExtension(pack *PackageInfo) string {
	ext := pack.FileName[len(pack.FileName)-extlength:]
	return ext
}

// package size limit by MB.
const mbColorLimit int64 = 1024 * 1024
const mb100ColorLimit = mbColorLimit * 100
const gbColorLimit = mbColorLimit * 1024

// the counter for sequence number.
var count = 1

func showCleaningItem(pack *PackageInfo) {
	color.Reset()

	ext := getPackageExtension(pack)

	fmt.Printf(packageInfoFormat, count, pack.Name, pack.Version, ext)
	count++

	if pack.Size < mbColorLimit {
		color.Set(color.LightGreen)
	} else if pack.Size < mb100ColorLimit {
		color.Set(color.Cyan)
	} else if pack.Size < gbColorLimit {
		color.Set(color.LightYellow)
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
	fmt.Println()
}
