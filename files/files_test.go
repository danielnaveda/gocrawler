package files

import (
	"reflect"
	"testing"
)

func TestConf_GetConf(t *testing.T) {
	type fields struct {
		Domains                  []string
		API                      string
		WorkersPerDomain         int
		MaxPagesCrawledPerDomain int
	}
	tests := []struct {
		name   string
		fields fields
		want   *Conf
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conf{
				Domains:                  tt.fields.Domains,
				API:                      tt.fields.API,
				WorkersPerDomain:         tt.fields.WorkersPerDomain,
				MaxPagesCrawledPerDomain: tt.fields.MaxPagesCrawledPerDomain,
			}
			if got := c.GetConf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conf.GetConf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateDirIfNotExist(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateDirIfNotExist(tt.args.dir)
		})
	}
}
