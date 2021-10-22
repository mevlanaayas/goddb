package goddb

import (
	"reflect"
	"testing"
)

var path = "../tmp/test"

func Test_defaultPersistenceService_Read(t *testing.T) {
	type fields struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		want   error
		want1  []byte
	}{
		{
			name: "Read should read given file",
			fields: fields{
				path: path,
			},
			want:  nil,
			want1: []byte("{\"exkey\":\"value\"}"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := filePersistenceService{
				path: tt.fields.path,
			}
			got, got1 := receiver.Read()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Read() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_defaultPersistenceService_Write(t *testing.T) {
	type fields struct {
		path string
	}
	type args struct {
		value []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Write should write values to given file",
			fields: fields{
				path: "",
			},
			args: args{
				value: []byte("{\"exkey1\":\"value1\"}"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := filePersistenceService{
				path: tt.fields.path,
			}
			err2 := receiver.Write(tt.args.value)
			if err := err2; (err != nil) != tt.wantErr {
				t.Errorf("Write() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
