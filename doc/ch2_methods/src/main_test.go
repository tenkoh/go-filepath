package main

import (
	"path/filepath"
	"testing"
)

func TestDir(t *testing.T) {
	tests := []struct {
		trgt string
		want string
	}{
		{"foo/bar/", "foo/bar"},
		{"foo/bar", "foo"},
		{"foo/bar/fizz.txt", "foo/bar"},
	}
	for _, tt := range tests {
		d := filepath.Dir(tt.trgt)
		// convert slash to each OS's separator
		w := filepath.FromSlash(tt.want)
		if d != w {
			t.Errorf("want %s, got %s", w, d)
		}
	}
}

func TestClean(t *testing.T) {
	tests := []struct {
		trgt string
		want string
	}{
		{"./foo/bar/", "foo/bar"},
		{"foo/bar/../", "foo"},
	}
	for _, tt := range tests {
		d := filepath.Clean(tt.trgt)
		// convert slash to each OS's separator
		w := filepath.FromSlash(tt.want)
		if d != w {
			t.Errorf("want %s, got %s", w, d)
		}
	}
}

func TestJoin(t *testing.T) {
	tests := []struct {
		trgts []string
		want  string
	}{
		{[]string{"./foo", "bar/"}, "foo/bar"},
	}
	for _, tt := range tests {
		d := filepath.Join(tt.trgts...)
		// convert slash to each OS's separator
		w := filepath.FromSlash(tt.want)
		if d != w {
			t.Errorf("want %s, got %s", w, d)
		}
	}
}

func TestSplit(t *testing.T) {
	tests := []struct {
		trgt     string
		wantDir  string
		wantFile string
	}{
		{"foo/bar/", "foo/bar/", ""},
		{"foo/bar", "foo/", "bar"},
		{"foo/bar/fizz.txt", "foo/bar/", "fizz.txt"},
	}
	for _, tt := range tests {
		d, f := filepath.Split(tt.trgt)
		if d != tt.wantDir {
			t.Errorf("want %s, got %s", tt.wantDir, d)
		}
		if f != tt.wantFile {
			t.Errorf("want %s, got %s", tt.wantFile, f)
		}
	}
}

func TestRel(t *testing.T) {
	tests := []struct {
		root string
		trgt string
		want string
	}{
		{"./foo/", "foo/bar", "bar"},
		{"./foo/", "foo/bar/fizz", "bar/fizz"},
	}
	for _, tt := range tests {
		r, _ := filepath.Rel(tt.root, tt.trgt)
		// convert slash to each OS's separator
		w := filepath.FromSlash(tt.want)
		if r != w {
			t.Errorf("want %s, got %s", w, r)
		}
	}
}
