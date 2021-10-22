package goddb

import (
	"reflect"
	"testing"
)

func TestNewStorageService(t *testing.T) {
	type args struct {
		repository         GetPutFlusher
		persistenceService ReadWriter
	}
	tests := []struct {
		name string
		args args
		want StorageService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStorageService(tt.args.repository, tt.args.persistenceService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewStorageService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorageService_Flush(t *testing.T) {
	type fields struct {
		repository         GetPutFlusher
		persistenceService ReadWriter
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := StorageService{
				repository:         tt.fields.repository,
				persistenceService: tt.fields.persistenceService,
			}
			if err := receiver.Flush(); (err != nil) != tt.wantErr {
				t.Errorf("Flush() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageService_Load(t *testing.T) {
	type fields struct {
		repository         GetPutFlusher
		persistenceService ReadWriter
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := StorageService{
				repository:         tt.fields.repository,
				persistenceService: tt.fields.persistenceService,
			}
			if err := receiver.Load(); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageService_Put(t *testing.T) {
	type fields struct {
		repository         GetPutFlusher
		persistenceService ReadWriter
	}
	type args struct {
		request SaveValue
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := StorageService{
				repository:         tt.fields.repository,
				persistenceService: tt.fields.persistenceService,
			}
			if err := receiver.Put(tt.args.request); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorageService_Retrieve(t *testing.T) {
	type fields struct {
		repository         GetPutFlusher
		persistenceService ReadWriter
	}
	type args struct {
		request RetrieveValue
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
		want1  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := StorageService{
				repository:         tt.fields.repository,
				persistenceService: tt.fields.persistenceService,
			}
			got, got1 := receiver.Retrieve(tt.args.request)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Retrieve() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestStorageService_Save(t *testing.T) {
	type fields struct {
		repository         GetPutFlusher
		persistenceService ReadWriter
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := StorageService{
				repository:         tt.fields.repository,
				persistenceService: tt.fields.persistenceService,
			}
			if err := receiver.Save(); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
