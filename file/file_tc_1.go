package reader

import "fmt"

var (
	tc1_k1 = "key1"
	tc1_k2 = "key2"
	tc1_v1 = "val1"
	tc1_v2 = "val2"
)
var tc1_content = fmt.Sprintf("%s=%s \n %s=%s", tc1_k1, tc1_v1, tc1_k2, tc1_v2)

var tc1_result EnvKeyValuePairs = EnvKeyValuePairs{
	tc1_k1: tc1_v1,
	tc1_k2: tc1_v2,
}

var tc_1 TestCase = TestCase{
	name:           "HappyPath",
	envFileContent: tc1_content,
	expectedResult: tc1_result,
}
