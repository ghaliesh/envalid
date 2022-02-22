package reader

import "fmt"

var (
	tc4_k1 = "key1"
	tc4_v1 = ""
)
var tc4_content = fmt.Sprintf("%s=%s ", tc4_k1, tc4_v1)

var tc4_result EnvKeyValuePairs = EnvKeyValuePairs{
	tc4_k1: tc4_v1,
}

var tc_4 TestCase = TestCase{
	name:           "EnvWithEmptyValue",
	envFileContent: tc4_content,
	expectedResult: tc4_result,
}
