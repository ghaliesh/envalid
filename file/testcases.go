package envfilereader

type TestCase struct {
	name           string
	envFileContent string
	expectedResult EnvKeyValuePairs
}

type TestCases = []TestCase

var testcases TestCases = []TestCase{tc_1, tc_2, tc_3}
