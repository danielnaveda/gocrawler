package files

import (
	"os"
)

// CreateDirIfNotExist creates a directory
func CreateDirIfNotExist(folderName string, createFolder func(string, os.FileMode) error, removeFolder func(string) error) error {
	err := removeFolder(folderName)

	if err != nil {
		return err
	}

	return createFolder(folderName, 0755)
}
