package service

import (
	"path/filepath"
	"testing"
)

func TestGetFullPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				path: "a",
			},
			want:    "a",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFullPath(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFullPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if filepath.Base(got) != filepath.Base(tt.want) {
				t.Errorf("GetFullPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
