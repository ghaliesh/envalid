package utils

import (
	"os"
	"path/filepath"
)

var dirname, _ = os.Getwd()
var path = filepath.Join(dirname, "../tmp/.env")

func CreateTmpEnvFile(data []byte) error {
	err := os.WriteFile(path, data, 0777)

	return err
}
