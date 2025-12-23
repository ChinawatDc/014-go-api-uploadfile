package utils

import (
	"os"
)

func RemoveFile(path string) error {
	return os.Remove(path)
}
