package filesys

import (
	"errors"
	"os"
)

var ErrPathNotExists = errors.New("path not exists")
var ErrPathIsNotDirectory = errors.New("provided path is not a directory")

func IsDirectory(path string) error {
	fi, err := os.Stat(path)

	if os.IsNotExist(err) {
		return ErrPathNotExists
	}

	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return ErrPathIsNotDirectory
	}

	return nil
}
