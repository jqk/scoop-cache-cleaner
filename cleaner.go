package main

import (
	"errors"
	"os"
	"strings"
	"time"
)

// ShowCleaningItem is used to display current cleaning package information.
type ShowCleaningItem func(pack *PackageInfo)

// CleanResult stores processing result of cleaning.
type CleanResult struct {
	FileCount     int
	CleanCount    int
	CleanSize     int64
	SoftwareCount int
	BackupPath    string
	CleanPackages []*PackageInfo
}

// PackageInfo stores the information of software installation file.
type PackageInfo struct {
	Name     string
	Version  string
	Size     int64
	FileName string
}

// ActionType defines the types of action that can be executed.
type ActionType int

const (
	List ActionType = iota
	Backup
	Delete
)

// ActionInfo stores the clean action information.
type ActionInfo struct {
	Action    ActionType
	ScoopPath string
}

// NewAction creates ActionInfo object according to provided strings.
func NewAction(action string, scoopPath string) *ActionInfo {
	if action == "-l" {
		return &ActionInfo{List, scoopPath}
	} else if action == "-b" {
		return &ActionInfo{Backup, scoopPath}
	} else if action == "-d" {
		return &ActionInfo{Delete, scoopPath}
	}

	return nil
}

// GetScoopPath gets the formal path string from the command parameter
// or environment variable. At last, ensure the path exists.
func GetScoopPath(param string) (string, error) {
	var s string
	var err error

	if param == "" {
		scoop := os.Getenv("SCOOP")
		if scoop == "" {
			return "", errors.New("environment variable SCOOP not found")
		}

		s, err = JoinFileName(scoop, "cache")
	} else {
		s, err = FormatFileName(param)
	}

	if err != nil {
		return "", err
	}

	if FileExists(s) {
		return s, nil
	} else {
		return "", errors.New("scoop cache path [" + s + "] does not exist")
	}
}

// CleanScoopCache moves outdated installation files to the backup directory.
func CleanScoopCache(action *ActionInfo, showItem ShowCleaningItem) (*CleanResult, error) {
	result, err := findOutdatedPackages(action.ScoopPath)
	if err != nil {
		return nil, err
	}

	if result.CleanCount > 0 {
		if action.Action == List {
			for _, p := range result.CleanPackages {
				showItem(p)
			}
		} else if action.Action == Backup {
			var err error
			result.BackupPath, err = prepareBackupPath(action.ScoopPath)
			if err != nil {
				return nil, err
			}

			for _, p := range result.CleanPackages {
				old, _ := JoinFileName(action.ScoopPath, p.FileName)
				new, _ := JoinFileName(result.BackupPath, p.FileName)

				if err := os.Rename(old, new); err != nil {
					return result, err
				}

				showItem(p)
			}
		} else if action.Action == Delete {
			for _, p := range result.CleanPackages {
				old, _ := JoinFileName(action.ScoopPath, p.FileName)

				if err := os.Remove(old); err != nil {
					return result, err
				}

				showItem(p)
			}
		}
	}

	return result, nil
}

// findOutdatedPackages finds outdated packages in specified path.
func findOutdatedPackages(scoopPath string) (*CleanResult, error) {
	f, err := os.Open(scoopPath)
	if err != nil {
		return nil, err
	}

	// get the list of files in the specified directory.
	// file names are in alphabetical ascending order.
	// so the latest package of each software is relatively behind the file list.
	files, err := f.Readdir(0)
	if err != nil {
		return nil, err
	}

	result := &CleanResult{0, 0, 0, 0, "", make([]*PackageInfo, 0)}
	count := len(files)
	newestPackage := PackageInfo{"", "", 0, ""}

	// process files in the list in reverse order.
	// then first package is the newest one.
	for i := count - 1; i >= 0; i-- {
		file := files[i]

		// skip directories.
		if file.IsDir() {
			continue
		}

		result.FileCount++

		// skip none scoop installation files by checking isPackage.
		if currentPackage, isPackage := getPackageInfo(file.Name()); isPackage {
			if currentPackage.Name != newestPackage.Name {
				result.SoftwareCount++

				newestPackage.Name = currentPackage.Name
				newestPackage.Version = currentPackage.Version
			} else if currentPackage.Version != newestPackage.Version {
				result.CleanCount++
				result.CleanSize += file.Size()
				result.CleanPackages = append(result.CleanPackages, currentPackage)

				currentPackage.Size = file.Size()
				currentPackage.FileName = file.Name()
			}
		}
	}

	return result, nil
}

// prepareBackupPath creates the backup directory when necessary.
func prepareBackupPath(scoopPath string) (string, error) {
	s, err := JoinFileName(scoopPath, time.Now().Format("bak_2006-01-02T15-04-05"))

	if err == nil && !FileExists(s) {
		if err = os.Mkdir(s, 0777); err != nil {
			return "", err
		}
	}

	return s, err
}

// get the package information from file name.
func getPackageInfo(fileName string) (*PackageInfo, bool) {
	// the installation file name in scoop format is "name#version#other-information".
	parts := strings.Split(fileName, "#")

	// ignore invalid file name.
	if len(parts) != 3 {
		return nil, false
	}

	return &PackageInfo{parts[0], parts[1], 0, ""}, true
}
