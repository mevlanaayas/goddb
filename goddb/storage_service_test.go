package goddb

import (
	"fmt"
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
	var verifyCalled bool

	type fields struct {
		repository         GetPutFlusher
		persistenceService ReadWriter
	}
	type args struct {
		flushMock func() error
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		wantCalled bool
	}{
		{
			name: "flush should not return error when flusher returns nil error",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				flushMock: func() error {
					verifyCalled = true
					return nil
				},
			},
			wantErr:    false,
			wantCalled: true,
		},
		{
			name: "flush should return error when flusher returns error",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				flushMock: func() error {
					verifyCalled = true
					return fmt.Errorf("error")
				},
			},
			wantErr:    true,
			wantCalled: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flushMock = tt.args.flushMock
			receiver := StorageService{
				repository:         tt.fields.repository,
				persistenceService: tt.fields.persistenceService,
			}
			err2 := receiver.Flush()
			if err := err2; (err != nil) != tt.wantErr {
				t.Errorf("Flush() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantCalled != verifyCalled {
				t.Errorf("call for Flush() method on flusher does not match mock. wanted: %v actual: %v", tt.wantCalled, verifyCalled)
			}
			verifyCalled = false
		})
	}
}

func TestStorageService_Load(t *testing.T) {
	var verifyPutAllCalled bool
	var verifyReadCalled bool
	var putAllValues map[string]string

	type fields struct {
		repository         GetPutFlusher
		persistenceService ReadWriter
	}
	type args struct {
		readMock   func() (error, []byte)
		putAllMock func(values map[string]string) error
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantErr          bool
		wantPutAllCalled bool
		wantReadCalled   bool
		expectedValues   map[string]string
	}{
		{
			name: "load should call reader and pass returned values to putter",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				readMock: func() (error, []byte) {
					verifyReadCalled = true
					return nil, []byte("{\"exkey1\":\"value1\"}")
				},
				putAllMock: func(values map[string]string) error {
					verifyPutAllCalled = true
					putAllValues = values
					return nil
				},
			},
			wantErr:          false,
			wantPutAllCalled: true,
			wantReadCalled:   true,
			expectedValues: map[string]string{
				"exkey1": "value1",
			},
		},
		{
			name: "load should return error when reader returns error",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				readMock: func() (error, []byte) {
					verifyReadCalled = true
					return fmt.Errorf("error"), nil
				},
				putAllMock: func(values map[string]string) error {
					verifyPutAllCalled = true
					putAllValues = values
					return nil
				},
			},
			wantErr:          true,
			wantPutAllCalled: false,
			wantReadCalled:   true,
			expectedValues:   map[string]string{},
		},
		{
			name: "load should return error when putter returns error",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				readMock: func() (error, []byte) {
					verifyReadCalled = true
					return nil, []byte("{\"exkey1\":\"value1\"}")
				},
				putAllMock: func(values map[string]string) error {
					verifyPutAllCalled = true
					putAllValues = values
					return fmt.Errorf("error")
				},
			},
			wantErr:          true,
			wantPutAllCalled: true,
			wantReadCalled:   true,
			expectedValues:   map[string]string{},
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			readMock = tt.args.readMock
			putAllMock = tt.args.putAllMock

			receiver := StorageService{
				repository:         tt.fields.repository,
				persistenceService: tt.fields.persistenceService,
			}
			err2 := receiver.Load()
			if err := err2; (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantPutAllCalled != verifyPutAllCalled {
				t.Errorf("call for PutAll() method on Putter does not match mock. wanted: %v actual: %v", tt.wantPutAllCalled, verifyPutAllCalled)
			}
			if tt.wantReadCalled != verifyReadCalled {
				t.Errorf("call for Read() method on Reader does not match mock. wanted: %v actual: %v", tt.wantReadCalled, verifyReadCalled)
			}
			if !tt.wantErr && !reflect.DeepEqual(tt.expectedValues, putAllValues) {
				t.Errorf("Putter error. expected %v, actual %v", tt.expectedValues, putAllValues)
			}
			verifyPutAllCalled = false
			verifyReadCalled = false
			putAllValues = map[string]string{}
		})
	}
}

func TestStorageService_Put(t *testing.T) {
	var verifyCalled bool
	var putKey string
	var putValue string

	type fields struct {
		repository         GetPutFlusher
		persistenceService ReadWriter
	}
	type args struct {
		request SaveValue
		putMock func(key, value string) error
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		wantCalled bool
	}{
		{
			name: "should return error when request is not valid",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				request: SaveValue{
					Key:   "mevlana",
					Value: "",
				},
				putMock: func(key, value string) error {
					putKey = key
					putValue = value
					verifyCalled = true
					return nil
				},
			},
			wantErr:    true,
			wantCalled: false,
		},
		{
			name: "should return error when putter returns error",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				request: SaveValue{
					Key:   "valid",
					Value: "request",
				},
				putMock: func(key, value string) error {
					putKey = key
					putValue = value
					verifyCalled = true
					return fmt.Errorf("error")
				},
			},
			wantErr:    true,
			wantCalled: true,
		},
		{
			name: "should return error when putter returns error",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				request: SaveValue{
					Key:   "valid",
					Value: "request",
				},
				putMock: func(key, value string) error {
					putKey = key
					putValue = value
					verifyCalled = true
					return fmt.Errorf("error")
				},
			},
			wantErr:    true,
			wantCalled: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			putMock = tt.args.putMock

			receiver := StorageService{
				repository:         tt.fields.repository,
				persistenceService: tt.fields.persistenceService,
			}
			err2 := receiver.Put(tt.args.request)
			if err := err2; (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantCalled != verifyCalled {
				t.Errorf("call for Put() method on Putter does not match mock. wanted: %v actual: %v", tt.wantCalled, verifyCalled)
			}
			if tt.wantCalled && (!reflect.DeepEqual(putKey, "valid") || !reflect.DeepEqual(putValue, "request")) {
				t.Errorf("call for Put() method on Putter does not match mock key:value. wanted: %v actual: %v", "valid:request", fmt.Sprintf("%s:%s", putKey, putValue))
			}
			verifyCalled = false
			putKey = ""
			putValue = ""
		})
	}
}

func TestStorageService_Retrieve(t *testing.T) {

	var verifyCalled bool
	var putKey string

	type fields struct {
		repository         GetPutFlusher
		persistenceService ReadWriter
	}
	type args struct {
		request RetrieveValue
		getMock func(key string) (error, string)
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantErr    bool
		wantCalled bool
		want       string
	}{
		{
			name: "should return error when request is not valid",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				request: RetrieveValue{
					Key: "",
				},
				getMock: func(key string) (error, string) {
					putKey = key
					verifyCalled = true
					return nil, ""
				},
			},
			wantErr:    true,
			wantCalled: false,
			want:       "",
		},
		{
			name: "should return error when getter returns error",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				request: RetrieveValue{
					Key: "valid",
				},
				getMock: func(key string) (error, string) {
					putKey = key
					verifyCalled = true
					return fmt.Errorf("error"), ""
				},
			},
			wantErr:    true,
			wantCalled: true,
			want:       "",
		},
		{
			name: "should return value when getter returns value",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				request: RetrieveValue{
					Key: "valid",
				},
				getMock: func(key string) (error, string) {
					putKey = key
					verifyCalled = true
					return nil, "test"
				},
			},
			wantErr:    false,
			wantCalled: true,
			want:       "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getMock = tt.args.getMock

			receiver := StorageService{
				repository:         tt.fields.repository,
				persistenceService: tt.fields.persistenceService,
			}
			err2, value := receiver.Retrieve(tt.args.request)
			if err := err2; (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(value, tt.want) {
				t.Errorf("Retrieve() got = %v, want %v", value, tt.want)
			}
			if tt.wantCalled != verifyCalled {
				t.Errorf("call for Put() method on Putter does not match mock. wanted: %v actual: %v", tt.wantCalled, verifyCalled)
			}
			if tt.wantCalled && !reflect.DeepEqual(putKey, "valid") {
				t.Errorf("call for Put() method on Putter does not match mock key. wanted: %v actual: %v", "valid", putKey)
			}
			verifyCalled = false
			putKey = ""
		})
	}
}

func TestStorageService_Save(t *testing.T) {

	var verifyGetAllCalled bool
	var verifyWriteCalled bool
	var writeValue []byte

	type fields struct {
		repository         GetPutFlusher
		persistenceService ReadWriter
	}
	type args struct {
		writeMock  func([]byte) error
		getAllMock func() (error, map[string]string)
	}
	tests := []struct {
		name               string
		fields             fields
		args               args
		wantErr            bool
		wantGetAllCalled   bool
		wantWriteCalled    bool
		expectedWriteValue []byte
	}{
		{
			name: "save should call getter and pass returned values to writer",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				writeMock: func(value []byte) error {
					verifyWriteCalled = true
					writeValue = value
					return nil
				},
				getAllMock: func() (error, map[string]string) {
					verifyGetAllCalled = true
					return nil, map[string]string{
						"exkey1": "value1",
					}
				},
			},
			wantErr:            false,
			wantGetAllCalled:   true,
			wantWriteCalled:    true,
			expectedWriteValue: []byte("{\"exkey1\":\"value1\"}"),
		},
		{
			name: "save should return error when getter returns error",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				writeMock: func(value []byte) error {
					verifyWriteCalled = true
					writeValue = value
					return nil
				},
				getAllMock: func() (error, map[string]string) {
					verifyGetAllCalled = true
					return fmt.Errorf("error"), nil
				},
			},
			wantErr:            true,
			wantGetAllCalled:   true,
			wantWriteCalled:    false,
			expectedWriteValue: []byte(""),
		},
		{
			name: "save should return error when writer returns error",
			fields: fields{
				repository:         &getPutFlusherMock{},
				persistenceService: &readWriterMock{},
			},
			args: args{
				writeMock: func(value []byte) error {
					verifyWriteCalled = true
					writeValue = value
					return fmt.Errorf("error")
				},
				getAllMock: func() (error, map[string]string) {
					verifyGetAllCalled = true
					return nil, map[string]string{
						"exkey1": "value1",
					}
				},
			},
			wantErr:            true,
			wantGetAllCalled:   true,
			wantWriteCalled:    true,
			expectedWriteValue: []byte("{\"exkey1\":\"value1\"}"),
		},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			writeMock = tt.args.writeMock
			getAllMock = tt.args.getAllMock

			receiver := StorageService{
				repository:         tt.fields.repository,
				persistenceService: tt.fields.persistenceService,
			}
			err2 := receiver.Save()
			if err := err2; (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantGetAllCalled != verifyGetAllCalled {
				t.Errorf("call for GetAll() method on Getter does not match mock. wanted: %v actual: %v", tt.wantGetAllCalled, verifyGetAllCalled)
			}
			if tt.wantWriteCalled != verifyWriteCalled {
				t.Errorf("call for Write() method on Writer does not match mock. wanted: %v actual: %v", tt.wantWriteCalled, verifyWriteCalled)
			}
			if !tt.wantErr && !reflect.DeepEqual(tt.expectedWriteValue, writeValue) {
				t.Errorf("Putter error. expected %v, actual %v", tt.expectedWriteValue, writeValue)
			}
			verifyGetAllCalled = false
			verifyWriteCalled = false
			writeValue = []byte("")
		})
	}
}
