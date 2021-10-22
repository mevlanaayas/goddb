package goddb

import (
	"reflect"
	"testing"
)

func Test_inMemoryRepository_Put(t *testing.T) {
	storage := make(map[string]string)

	type fields struct {
		storage map[string]string
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "put should save key:value on map",
			fields: fields{
				storage: storage,
			},
			args: args{
				key:   "testKey",
				value: "testValue",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := inMemoryRepository{
				storage: tt.fields.storage,
			}
			err2 := receiver.Put(tt.args.key, tt.args.value)
			if err := err2; (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
			if receiver.storage[tt.args.key] != tt.args.value {
				t.Errorf("Put() error. Put: %s Got: %s", receiver.storage[tt.args.key], tt.args.value)
			}
		})
	}
}

func Test_inMemoryRepository_Retrieve(t *testing.T) {
	storage := make(map[string]string)
	storage["testKey"] = "testValue"
	type fields struct {
		storage map[string]string
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
		want1  string
	}{
		{
			name: "retrieve should get value by key from map",
			fields: fields{
				storage: storage,
			},
			args: args{
				key: "testKey",
			},
			want:  nil,
			want1: "testValue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := inMemoryRepository{
				storage: tt.fields.storage,
			}
			got, got1 := receiver.Retrieve(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Retrieve() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Retrieve() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_inMemoryRepository_Flush(t *testing.T) {
	storage := make(map[string]string)
	storage["testKey"] = "testValue"
	storage["testKey1"] = "testValue1"

	type fields struct {
		storage map[string]string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "flush should remove all key:values from map",
			fields: fields{
				storage: storage,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := &inMemoryRepository{
				storage: tt.fields.storage,
			}
			receiver.Flush()

			if len(receiver.storage) != 0 {
				t.Errorf("After flush size we got = %v, want 0", len(storage))
			}
		})
	}
}

func Test_inMemoryRepository_Get(t *testing.T) {
	storage := make(map[string]string)
	storage["testKey"] = "testValue"
	storage["testKey1"] = "testValue1"
	storage["testKey2"] = "testValue2"

	type fields struct {
		storage map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "get should return all key:values from map",
			fields: fields{
				storage: storage,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := inMemoryRepository{
				storage: tt.fields.storage,
			}
			err2, values := receiver.Get()
			if err := err2; (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(values, storage) {
				t.Errorf("Get() error. Expected map and actual map is different")
			}
		})
	}
}

func Test_inMemoryRepository_Load(t *testing.T) {
	values := make(map[string]string)
	values["testKey"] = "testValue"
	values["testKey1"] = "testValue1"
	values["testKey2"] = "testValue2"

	type fields struct {
		storage map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "load should save all key:values to storage",
			fields: fields{
				storage: make(map[string]string),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := inMemoryRepository{
				storage: tt.fields.storage,
			}
			err2 := receiver.Load(values)
			if err := err2; (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(values, receiver.storage) {
				t.Errorf("Load() error. Expected map and actual map is different")
			}
		})
	}
}
