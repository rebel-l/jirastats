package utils

import (
	"os"
)

// FileExists checks if a path or file exists
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

// CreateDirectoryIfNotExists creates a path recursive
func CreateDirectoryIfNotExists(path string) (err error) {
	exist, _ := FileExists(path)
	if exist == false {
		err = os.MkdirAll(path, 0755)
	}
	return
}

func CreateFileIfNotExists(file string) (err error) {
	exist, _ := FileExists(file)
	if exist == false {
		_, err = os.Create(file)
	}
	return
}
