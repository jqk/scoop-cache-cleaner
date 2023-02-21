package main

import (
	"fmt"
	"os"

	"github.com/gookit/color"
)

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
	r, err := CleanScoopCache(action, showCleaningItem)
	showCleanResult(r, err)
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
	fmt.Println("scoop cache cleaner (scc) 2.0.0, 2023-02-21")
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
	fmt.Println(" list the outdated packages.")

	color.Set(color.LightYellow)
	fmt.Print("  -b: ")
	color.Reset()
	fmt.Println(" backup the outdated packages.")

	color.Set(color.LightYellow)
	fmt.Print("  -d: ")
	color.Reset()
	fmt.Println(" delete the outdated packages.")

	color.Set(color.LightYellow)
	fmt.Println("\nall other parameters will display the above information.")
	fmt.Println()
	color.Reset()
}

func showCleanStart(action *ActionInfo) {
	color.Set(color.LightRed)

	if action.Action == List {
		fmt.Print("List")
	} else if action.Action == Backup {
		fmt.Print("Backup")
	} else { // if action.Action == Delete
		fmt.Print("Delete")
	}

	color.Reset()
	fmt.Print(" outdated packages in: ")
	color.Set(color.LightGreen)
	fmt.Println(action.ScoopPath)
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
	fmt.Println("File found            :", result.FileCount)
	fmt.Println("Package found         :", result.SoftwareCount)
	fmt.Println("Outdated Package found:", result.CleanCount)
	fmt.Print("Outdated Package Size : ")
	color.Set(color.LightRed)

	var kb int64 = 1024
	var mb = kb * kb
	var gb = mb * kb
	var size = result.CleanSize

	if size < kb {
		fmt.Println(size, "bytes")
	} else if size < mb {
		fmt.Printf("%.2f KB\n", float64(size)/float64(kb))
	} else if size < gb {
		fmt.Printf("%.2f MB\n", float64(size)/float64(mb))
	} else {
		fmt.Printf("%.2f GB\n", float64(size)/float64(gb))
	}

	color.Reset()
	fmt.Println("-------------------")

	if result.CleanCount > 0 && result.BackupPath != "" {
		if result.CleanCount == 1 {
			fmt.Print("Outdated Package has been moved to ")
		} else {
			fmt.Print("Outdated Packages have been moved to ")
		}

		color.Set(color.LightGreen)
		fmt.Println(result.BackupPath)
	}

	fmt.Println()
	color.Reset()
}

func showError(err error) {
	color.Set(color.Red)
	fmt.Println("---------- Error! ----------")
	fmt.Println(err)
	color.Reset()
}
