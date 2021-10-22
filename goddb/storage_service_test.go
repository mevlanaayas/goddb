package goddb

import (
	"reflect"
	"testing"
)

var (
	readMock   func() (error, []byte)
	writeMock  func([]byte) error
	getMock    func(key string) (error, string)
	getAllMock func() (error, map[string]string)
	putMock    func(key, value string) error
	putAllMock func(values map[string]string) error
	flushMock  func() error
)

type readWriterMock struct {
}

func (receiver readWriterMock) Read() (error, []byte) {
	return readMock()
}

func (receiver readWriterMock) Write(value []byte) error {
	return writeMock(value)
}

type getPutFlusherMock struct {
}

func (receiver getPutFlusherMock) Put(key, value string) error {
	return putMock(key, value)
}

func (receiver getPutFlusherMock) Get(key string) (error, string) {
	return getMock(key)
}

func (receiver *getPutFlusherMock) Flush() error {
	return flushMock()
}

func (receiver getPutFlusherMock) GetAll() (error, map[string]string) {
	return getAllMock()
}

func (receiver *getPutFlusherMock) PutAll(values map[string]string) error {
	return putAllMock(values)
}

func TestStorageService_Flush(t *testing.T) {
	verifyCalled := false
	flushMock = func() error {
		verifyCalled = true
		return nil
	}
	type fields struct {
		repository         GetPutFlusher
		persistenceService ReadWriter
	}
	tests := []struct {
		name       string
		fields     fields
		wantErr    bool
		wantCalled bool
	}{
		{
			name: "flush should not return error when flusher returns nil error",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			wantErr:    false,
			wantCalled: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := StorageService{
				repository:         tt.fields.repository,
				persistenceService: tt.fields.persistenceService,
			}
			err2 := receiver.Flush()
			if err := err2; (err != nil) != tt.wantErr {
				t.Errorf("Flush() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantCalled != verifyCalled {
				t.Errorf("no call for Flush() method on Flusher ")
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
