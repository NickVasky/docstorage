package dto

import (
	"testing"
)

func toPtr[T any](value T) *T {
	return &value
}

func TestValidator(t *testing.T) {

	tests := []struct {
		name    string
		arg     MetaUploadRequest
		wantErr bool
	}{
		{
			name: "Empty (nil) mime)",
			arg: MetaUploadRequest{
				Name:     toPtr("Name"),
				IsFile:   toPtr(true),
				IsPublic: toPtr(true),
				Mime:     nil, //toPtr("application/json"),
				Grant:    []string{},
			},
			wantErr: true,
		},
		{
			name: "Empty (nil) bool",
			arg: MetaUploadRequest{
				Name:     toPtr("Name"),
				IsFile:   toPtr(true),
				IsPublic: nil, //toPtr(true),
				Mime:     toPtr("application/json"),
				Grant:    []string{},
			},
			wantErr: true,
		},
		{
			name: "Empty string as mime",
			arg: MetaUploadRequest{
				Name:     toPtr("Name"),
				IsFile:   toPtr(true),
				IsPublic: toPtr(true),
				Mime:     toPtr(""),
				Grant:    []string{},
			},
			wantErr: true,
		},
		{
			name: "Valid mime",
			arg: MetaUploadRequest{
				Name:     toPtr("Name"),
				IsFile:   toPtr(true),
				IsPublic: toPtr(true),
				Mime:     toPtr("application/json"),
				Grant:    []string{},
			},
			wantErr: false,
		},
		{
			name: "Empty name",
			arg: MetaUploadRequest{
				Name:     toPtr(""),
				IsFile:   toPtr(true),
				IsPublic: toPtr(true),
				Mime:     toPtr("application/json"),
				Grant:    []string{},
			},
			wantErr: true,
		},
		{
			name: "Name len = 1",
			arg: MetaUploadRequest{
				Name:     toPtr("N"),
				IsFile:   toPtr(true),
				IsPublic: toPtr(true),
				Mime:     toPtr("application/json"),
				Grant:    []string{},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		if err := Validate.Struct(tt.arg); (err != nil) != tt.wantErr {
			t.Errorf("Case '%v'\nInput: %v\nError = %v, wantErr %v\n\n", tt.name, tt.arg, err, tt.wantErr)
		}
	}
}
