package main

import (
	"fmt"
	"os"
)

var packageInfoFormat string

func main() {
	showVersion()

	action, err := parseCmdParameter()
	if err != nil {
		showHelp()
		showError(err)
		return
	} else if action == nil {
		showHelp()
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

func parseCmdParameter() (*ActionInfo, error) {
	n := len(os.Args)

	if n == 1 {
		return newAction("", "")
	} else if n == 2 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			return nil, nil
		}
		return newAction(os.Args[1], "")
	} else if n == 3 {
		return newAction(os.Args[1], os.Args[2])
	} else {
		return nil, nil
	}
}

// newAction creates ActionInfo object according to provided strings.
func newAction(cmd string, scoopPath string) (*ActionInfo, error) {
	scoopPath, err := GetScoopPath(scoopPath)

	if cmd == "" {
		if err != nil {
			return nil, nil
		}
		return &ActionInfo{List, scoopPath}, nil
	} else if cmd == "-l" || cmd == "--list" {
		return &ActionInfo{List, scoopPath}, err
	} else if cmd == "-b" || cmd == "--backup" {
		return &ActionInfo{Backup, scoopPath}, err
	} else if cmd == "-d" || cmd == "--delete" {
		return &ActionInfo{Delete, scoopPath}, err
	}

	return nil, fmt.Errorf("unknown command: %s", cmd)
}
