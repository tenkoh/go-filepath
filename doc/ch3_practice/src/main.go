package main

import "path/filepath"

func makeBrotherPath(path, name string) string {
	cleaned := filepath.Clean(path)
	parent := filepath.Dir(cleaned)
	brother := filepath.Join(parent, name)
	return brother
}

func main() {

}
