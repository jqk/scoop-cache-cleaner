package main

import (
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
	SoftwareCount int
	BackupPath    string
}

// PackageInfo stores the information of software installation file.
type PackageInfo struct {
	Name    string
	Version string
}

// prepareBackupPath creates the backup directory when necessary.
func prepareBackupPath(scoopPath string) (string, error) {
	s, err := JoinFileName(scoopPath, time.Now().Format("bak_2006-01-02T15-04-05"))

	if err == nil && !IsFileExists(s) {
		if err = os.Mkdir(s, 0777); err != nil {
			return "", err
		}
	}

	return s, err
}

// CleanScoopCache moves outdated installation files to the backup directory.
func CleanScoopCache(scoopPath string, showItem ShowCleaningItem) (*CleanResult, error) {
	backupPath, err := prepareBackupPath(scoopPath)
	if err != nil {
		return nil, err
	}

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

	count := len(files)
	result := CleanResult{0, 0, 0, backupPath}
	newestPackage := PackageInfo{"", ""}

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

				old, _ := JoinFileName(scoopPath, file.Name())
				new, _ := JoinFileName(backupPath, file.Name())

				if err := os.Rename(old, new); err != nil {
					return &result, err
				}

				showItem(currentPackage)
			}
		}
	}

	if result.CleanCount == 0 {
		os.Remove(backupPath)
	}

	return &result, nil
}

// get the package information from file name.
func getPackageInfo(fileName string) (*PackageInfo, bool) {
	// the installation file name in scoop format is "name#version#other-information".
	parts := strings.Split(fileName, "#")

	// ignore invalid file name.
	if len(parts) != 3 {
		return nil, false
	}

	return &PackageInfo{parts[0], parts[1]}, true
}
