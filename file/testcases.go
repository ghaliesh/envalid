package reader

import (
	"os"
	"path/filepath"
)

var envReader = EnvFileReader{}
var dirname, _ = os.Getwd()
var path = filepath.Join(dirname, "../tmp/.env")
var invalidpath = "in/valid/pa/th"

func funcToPanicDuetoEmptyFile() {
	envReader.ReadEnvFile(path)
}

func funcToPanicDuetoNonExistentPath() {
	envReader.ReadEnvFile(invalidpath)
}

type TestCase struct {
	name           string
	envFileContent string
	expectedResult EnvKeyValuePairs
}

type TestCases = []TestCase

var testcases TestCases = []TestCase{tc_1, tc_2, tc_3, tc_4, tc_5}
