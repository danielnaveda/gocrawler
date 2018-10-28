package conf

import (
	"errors"
	"reflect"
	"testing"
)

func TestConf_GetConf(t *testing.T) {
	type fields struct {
		Domains                  []string
		API                      string
		WorkersPerDomain         int
		MaxPagesCrawledPerDomain int
		SaveIntoFiles            bool
	}
	type args struct {
		fileName     string
		reader       func(string) ([]byte, error)
		unmarshaller func([]byte, interface{}) error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Conf
		wantErr bool
	}{
		{
			`Normal flow`,
			fields{},
			args{
				`some-file`,
				func(a string) ([]byte, error) { return make([]byte, 1), nil },
				func(a []byte, b interface{}) error { return nil },
			},
			&Conf{},
			false,
		},
		{
			`Error flow`,
			fields{},
			args{
				`some-file`,
				func(a string) ([]byte, error) { return make([]byte, 1), errors.New(`Can't read file`) },
				func(a []byte, b interface{}) error { return nil },
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Conf{
				Domains:                  tt.fields.Domains,
				API:                      tt.fields.API,
				WorkersPerDomain:         tt.fields.WorkersPerDomain,
				MaxPagesCrawledPerDomain: tt.fields.MaxPagesCrawledPerDomain,
				SaveIntoFiles:            tt.fields.SaveIntoFiles,
			}
			got, err := c.getConfFile(tt.args.fileName, tt.args.reader, tt.args.unmarshaller)
			if (err != nil) != tt.wantErr {
				t.Errorf("Conf.GetConf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Conf.GetConf() = %v, want %v", got, tt.want)
			}
		})
	}
}
