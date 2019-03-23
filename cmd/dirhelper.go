package cmd

import (
	"fmt"
	"os"
	"strings"
)

func createDirs(dirs []string) error {
	for i, dir := range dirs {
		err := os.MkdirAll(dir, filePermissions)
		if err != nil {
			if delErr := deleteDirs(dirs[:i]); delErr != nil {
				err = fmt.Errorf(`error creating directory "%s": %s, unable to clean up: %s`, dirs[i], err.Error(), delErr.Error())
			}
			return err
		}
	}
	return nil
}

func deleteDirs(dirs []string) error {
	for _, dir := range dirs {
		if err := os.RemoveAll(dir); err != nil {
			return err
		}
	}
	return nil
}

func formatDirName(dir string) (string, error) {
	if len(dir) == 0 {
		return "", fmt.Errorf("Invalid directory %s", dir)
	}
	val := dir
	val = strings.TrimSpace(val)
	val = strings.ToLower(val)
	val = strings.NewReplacer("\n", "", "\r", "", "\t", "", "'", "", " ", "-").Replace(val)
	return val, nil
}
