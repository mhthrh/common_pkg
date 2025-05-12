package directory

import (
	"errors"
	"os"
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
