package filekit

import (
	"path/filepath"
)

func Abs(path string) string {
	return filepath.Abs(path)
}