package main

import (
	"os"
	"path/filepath"
	"reflect"
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
		got := MakeBrotherPath(tt.src, tt.name)
		// fit "want" to each OS
		want := filepath.FromSlash(tt.want)
		if got != want {
			t.Errorf("want %s, got %s", want, got)
		}
	}
}

func TestGetType(t *testing.T) {
	tests := []struct {
		target string
		want   string
	}{
		{"./", "dir"},
		{"./main.go", "file"},
	}
	for _, tt := range tests {
		got := GetType(tt.target)
		if got != tt.want {
			t.Errorf("want %s, got %s", tt.want, got)
		}
	}
}

func TestMain(m *testing.M) {
	testDirs := []string{"workdir/src/foo", "workdir/src/bar"}
	for _, dir := range testDirs {
		if err := os.MkdirAll(filepath.Clean(dir), 0755); err != nil {
			return
		}
	}
	defer os.RemoveAll("workdir")

	f, err := os.Create(filepath.Clean("workdir/src/foo/fizz.txt"))
	if err != nil {
		return
	}
	f.Close()

	m.Run()
}

func TestCopyTree(t *testing.T) {
	dst := "workdir/dst"
	tests := []struct {
		src       string
		wantFiles []string
		wantDirs  []string
	}{
		{
			"workdir/src",
			[]string{"workdir/dst/foo/fizz.txt"},
			[]string{"workdir/dst", "workdir/dst/bar", "workdir/dst/foo"},
		},
		{
			"workdir/src/foo/fizz.txt",
			[]string{"workdir/dst/fizz.txt"},
			[]string{},
		},
		{
			"workdir/src/bar/",
			[]string{},
			[]string{"workdir/dst"},
		},
	}
	for _, tt := range tests {
		fs, ds, err := CopyTree(tt.src, dst)
		if err != nil {
			t.Error(err)
			continue
		}

		// replace slash to OS's separator
		wfs, wds := []string{}, []string{}
		for _, f := range tt.wantFiles {
			wfs = append(wfs, filepath.FromSlash(f))
		}
		for _, d := range tt.wantDirs {
			wds = append(wds, filepath.FromSlash(d))
		}

		if len(wfs) == 0 {
			if len(fs) != 0 {
				t.Errorf("want %v, got %v", wfs, fs)
			}
		} else {
			if !reflect.DeepEqual(fs, wfs) {
				t.Errorf("want %v, got %v", wfs, fs)
			}
		}

		if len(wds) == 0 {
			if len(ds) != 0 {
				t.Errorf("want %v, got %v", wds, ds)
			}
		} else {
			if !reflect.DeepEqual(ds, wds) {
				t.Errorf("want %v, got %v", wds, ds)
			}
		}
	}
}
