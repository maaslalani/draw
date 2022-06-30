package main_test

import "testing"
import "github.com/maaslalani/draw"

func Test_filePath(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty string", args{""}, DEFAULT_DRAW_FILE},
		{"text file", args{"some.txt"}, "some.txt"},
		{"png file", args{"some.png"}, "some.png"},
		{"jpeg file", args{"some.jpg"}, "some.jpg"},
		{"jpeg file", args{"some.jpeg"}, "some.jpeg"},
		{"other file", args{"some.gif"}, "some.gif"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := filePath(tt.args.file); got != tt.want {
				t.Errorf("filePath(%v) = %v, want %v", tt.args.file, got, tt.want)
			}
		})
	}
}
