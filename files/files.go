package files

import (
	"os"
)

// CreateDirIfNotExist creates a directory
func CreateDirIfNotExist(dir string) {
	os.RemoveAll(dir)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
