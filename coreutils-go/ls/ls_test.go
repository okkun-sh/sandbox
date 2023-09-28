package main

import (
	"os"
	"testing"
)

func TestListead(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "display filenames",
			want: "bar.txt  foo.txt  hoge.txt",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			os.Chdir("testdata")
			defer os.Chdir("..")

			dir, err := os.Getwd()
			if err != nil {
				t.Error(err)
			}

			output, err := list(dir)
			if err != nil {
				t.Error(err)
			}

			if v, w := output, tt.want; v != w {
				t.Errorf("list() got = %v, want %v", v, w)
			}
		})
	}
}
