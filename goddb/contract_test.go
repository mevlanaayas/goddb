package goddb

import "testing"

func TestRetrieveValue_Validate(t *testing.T) {
	type fields struct {
		Key string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "validate should return error when key is not valid",
			fields: fields{
				Key: "",
			},
			wantErr: true,
		},
		{
			name: "validate should not return error when key is valid",
			fields: fields{
				Key: "mevlana",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := RetrieveValue{
				Key: tt.fields.Key,
			}
			if err := receiver.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSaveValue_Validate(t *testing.T) {
	type fields struct {
		Key   string
		Value string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "validate should return error when key is not valid",
			fields: fields{
				Key:   "",
				Value: "valid value",
			},
			wantErr: true,
		},
		{
			name: "validate should return error when value is not valid",
			fields: fields{
				Key:   "valid key",
				Value: "",
			},
			wantErr: true,
		},
		{
			name: "validate should return error when both key and value are not valid",
			fields: fields{
				Key:   "",
				Value: "",
			},
			wantErr: true,
		},
		{
			name: "validate should not return error when both key and value are valid",
			fields: fields{
				Key:   "valid key",
				Value: "valid value",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := SaveValue{
				Key:   tt.fields.Key,
				Value: tt.fields.Value,
			}
			if err := receiver.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
