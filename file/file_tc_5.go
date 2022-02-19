package reader

var tc5_content = ""

var tc5_result EnvKeyValuePairs = EnvKeyValuePairs{}

var tc_5 TestCase = TestCase{
	name:           "EnvWithEmptyValue",
	envFileContent: tc5_content,
	expectedResult: tc5_result,
}
