package main

import (
	"path/filepath"
	"testing"
)

func TestMakeBrotherPath(t *testing.T) {
	tests := []struct {
		src  string
		name string
		want string
	}{
		{"./workdir/src", "dst", "workdir/dst"},
		{"./workdir/src/", "dst", "workdir/dst"},
		{"workdir/src", "dst", "workdir/dst"},
		{"workdir/src/", "dst", "workdir/dst"},
	}
	for _, tt := range tests {
		got := makeBrotherPath(tt.src, tt.name)
		// fit "want" to each OS
		want := filepath.FromSlash(tt.want)
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	}
}
