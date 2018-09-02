package files

import "os"

type OsInterface interface {
	RemoveAll(string)
	Stat(string) (os.FileInfo, error)
	IsNotExist(error) bool
	MkdirAll(string, os.FileMode) error
}

// CreateDirIfNotExist creates a directory
func CreateDirIfNotExist(dir string, os OsInterface) {
	os.RemoveAll(dir)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
