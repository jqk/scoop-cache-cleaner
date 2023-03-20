package main

import (
	"fmt"
	"os"
	"strings"

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
	fmt.Println("scoop cache cleaner (scc) 2.1.3, 2023-03-20")
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

	d := strings.LastIndex(ext, ".")
	s := ext[:d]

	a := strings.LastIndex(s, ".")
	b := strings.LastIndex(s, "-")
	c := strings.LastIndex(s, "_")

	if a < b {
		a = b
	}

	if a < c {
		a = c
	}

	a++
	ext = ext[a:]

	return ext
}

// package size limit by MB.
const mbColorLimit int64 = 1024 * 1024
const mb100ColorLimit = mbColorLimit * 100

// the counter for sequence number.
var count = 1

// previous package file extension.
var lastExt = ""

// pervious package info.
var lastPack *PackageInfo = nil

func showCleaningItem(pack *PackageInfo) {
	color.Reset()

	ext := getPackageExtension(pack)
	same := false

	if lastPack != nil && lastPack.Name == pack.Name {
		// current package is as same as last one.
		// version can be same or diff.
		// then check the last extension.
		same = ext == lastExt
	}

	lastPack = pack
	lastExt = ext

	if same {
		ext = ""
	}

	fmt.Printf(packageInfoFormat, count, pack.Name, pack.Version, ext)
	count++

	if pack.Size < mbColorLimit {
		color.Set(color.Magenta)
	} else if pack.Size < mb100ColorLimit {
		color.Set(color.Red)
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
