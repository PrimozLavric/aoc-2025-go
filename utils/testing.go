package utils

import (
	"path/filepath"
	"runtime"
)

func GetTestDataPath(relativePath string) string {
	_, filename, _, _ := runtime.Caller(1)
	dir := filepath.Dir(filename)
	return filepath.Join(dir, "testdata", relativePath)
}
