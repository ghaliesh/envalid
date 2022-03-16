package reader

import (
	"os"
	"path/filepath"
)

var dirname, _ = os.Getwd()
var path = filepath.Join(dirname, "../tmp/.env")

type TestCase struct {
	name           string
	envFileContent string
	expectedResult EnvKeyValuePairs
}

type TestCases = []TestCase

var testcases TestCases = []TestCase{tc_1, tc_2, tc_3, tc_4, tc_5}
