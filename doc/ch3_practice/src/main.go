package main

import (
	"io/fs"
	"os"
	"path/filepath"
)

func MakeBrotherPath(path, name string) string {
	cleaned := filepath.Clean(path)
	parent := filepath.Dir(cleaned)
	brother := filepath.Join(parent, name)
	return brother
}

func GetType(path string) string {
	info, _ := os.Stat(path)
	if !info.IsDir() {
		return "file"
	}
	return "dir"
}

// CopyTree returns files []string, dirs []string, error
func CopyTree(path, name string) ([]string, []string, error) {
	// なにはともあれClean
	path = filepath.Clean(path)
	info, err := os.Stat(path)
	if err != nil {
		return nil, nil, err
	}
	// 与えられたのがファイルパスだったら、{name}/ファイル名を返しておわり
	if !info.IsDir() {
		f := filepath.Join(name, filepath.Base(path))
		return []string{f}, nil, nil
	}

	// 与えられたのがディレクトリだったら順々に走査する
	files, dirs := []string{}, []string{}
	err = filepath.WalkDir(path, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// pathをrootとする相対位置を求めて{name}に結合する
		rel, err := filepath.Rel(path, p)
		if err != nil {
			return err
		}
		dst := filepath.Join(name, rel)
		if !d.IsDir() {
			files = append(files, dst)
			return nil
		}
		dirs = append(dirs, dst)
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return files, dirs, nil
}

func main() {

}
