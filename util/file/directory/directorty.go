package directory

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func Create(path string) error {
	if Exist(path) {
		return errors.New("path already exist")
	}
	return os.Mkdir(path, 0755)
}

func Remove(path string) error {
	if !Exist(path) {
		return errors.New("path doesnt exist")
	}
	return os.Remove(path)
}

func Exist(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		return false
	}
}

func GetFileName(path string) ([]string, error) {
	if !Exist(path) {
		return nil, errors.New("not exist")
	}
	filename := filepath.Base(path)
	ext := filepath.Ext(filename)
	return []string{filename, ext, strings.TrimSuffix(filename, ext)}, nil
}
