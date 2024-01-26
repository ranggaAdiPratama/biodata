package util

import (
	"os"
	"path/filepath"
	"strings"
)

func DeleteFile(filepath string) error {
	err := os.Remove(filepath)

	if err != nil {
		return err
	}

	return nil
}

func FileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	return !os.IsNotExist(err)
}

func GetFileExtension(filename string) string {
	ext := filepath.Ext(filename)

	return strings.ToLower(ext)
}
