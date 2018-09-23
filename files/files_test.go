package files

import (
	"errors"
	"os"
	"testing"
)

func TestCreateDirIfNotExist(t *testing.T) {
	type args struct {
		folderName   string
		createFolder func(string, os.FileMode) error
		removeFolder func(string) error
	}
	errors := []error{
		errors.New(`Folder could not be removed`),
	}
	tests := []struct {
		name   string
		args   args
		result error
	}{
		{
			"It should be okay",
			args{
				"some-folder",
				func(a string, b os.FileMode) error { return nil },
				func(a string) error { return nil },
			},
			nil,
		},
		{
			"It should display error",
			args{
				"some-folder",
				func(a string, b os.FileMode) error { return nil },
				func(a string) error { return errors[0] },
			},
			errors[0],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CreateDirIfNotExist(tt.args.folderName, tt.args.createFolder, tt.args.removeFolder)

			if result != tt.result {
				t.Errorf("Expected: \"%v\" is different than actual error: \"%v\"", result, tt.result)
			}
		})
	}
}
