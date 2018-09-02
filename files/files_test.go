package files

import (
	"testing"
)

func TestCreateDirIfNotExist(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"My first test case",
			args{
				"hello",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateDirIfNotExist(tt.args.dir)
		})
	}
}
