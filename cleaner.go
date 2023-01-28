package main

import (
	"os"
	"path"
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
	s := path.Join(scoopPath, time.Now().Format("bak_2006-01-02T15-04-05"))

	if err := os.Mkdir(s, 0777); err != nil {
		return "", err
	}

	return FormatPath(s)
}

// CleanScoopCache moves outdated installation files to the backup directory.
func CleanScoopCache(scoopPath string, showItem ShowCleaningItem) (*CleanResult, error) {
	backupPath, err := prepareBackupPath(scoopPath)
	if err != nil {
		return nil, err
	}

	if err := CheckPathExists(scoopPath); err != nil {
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

	result := CleanResult{0, 0, 0, backupPath}
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

				pack := &PackageInfo{name, version}

				showItem(pack)
			}
		}
	}

	if result.CleanCount == 0 {
		os.Remove(backupPath)
	}

	return &result, nil
}
