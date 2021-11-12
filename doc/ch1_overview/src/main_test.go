package main

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// Good
func TestJoin(t *testing.T) {
	dst := "foo"
	savepath := filepath.Join(dst, "bar.txt")

	var want string
	switch runtime.GOOS {
	case "windows":
		want = "foo\\bar.txt"
	default:
		want = "foo/bar.txt"
	}
	if savepath != want {
		t.Errorf("want %s, got %s", want, savepath)
	}
}

// this work on both Unix and Windows. MkdirAll properly handles separators.
func TestMkDir(t *testing.T) {
	dst := "workdir/test"
	if err := os.MkdirAll(dst, 0755); err != nil {
		t.Error(err)
	}
}
