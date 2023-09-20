package main

import (
	"os/exec"
	"path/filepath"
	"testing"
)

func TestRead(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		opt      bool
		want     string
		wantErr  bool
	}{
		{
			name:     "display file content",
			filePath: filepath.Join("testdata", "test.txt"),
			opt:      false,
			want:     cmdExec(t, filepath.Join("testdata", "test.txt"), ""),
			wantErr:  false,
		},
		{
			name:     "display file content without blank lines",
			opt:      true,
			filePath: filepath.Join("testdata", "blank_line.txt"),
			want:     cmdExec(t, filepath.Join("testdata", "blank_line.txt"), "-n"),
			wantErr:  false,
		},
		{
			name:     "file not found",
			filePath: filepath.Join("testdata", "notfound.txt"),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s, err := read(tt.filePath, tt.opt)
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

func cmdExec(t *testing.T, path string, opt string) string {
	var cmd *exec.Cmd
	if opt != "" {
		cmd = exec.Command("cat", opt, path)
	} else {
		cmd = exec.Command("cat", path)
	}

	out, err := cmd.Output()
	if err != nil {
		t.Error(err)
	}

	return string(out)
}
