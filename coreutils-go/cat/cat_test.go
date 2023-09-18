package main

import (
	"path/filepath"
	"testing"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		want     string
		wantErr  bool
	}{
		{
			name:     "display file content",
			filePath: filepath.Join("testdata", "test.txt"),
			want:     "hoge\nbar\nfoo",
			wantErr:  false,
		},
		{
			name:     "file not found",
			filePath: "testdata/notfound.txt",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s, err := read(tt.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && s != tt.want {
				t.Errorf("read() got = %v, want %v", s, tt.want)
			}
		})
	}
}
