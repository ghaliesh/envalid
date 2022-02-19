package reader

import "fmt"

var (
	tc2_k1 = "key1"
	tc2_k2 = "key2"
	tc2_v1 = "val1=1212"
	tc2_v2 = "val2"
)
var tc2_content = fmt.Sprintf("%s=%s \n %s=%s", tc2_k1, tc2_v1, tc2_k2, tc2_v2)

var tc2_result EnvKeyValuePairs = EnvKeyValuePairs{
	tc2_k1: tc2_v1,
	tc2_k2: tc2_v2,
}

var tc_2 TestCase = TestCase{
	name:           "EnvWithTwoSeperators",
	envFileContent: tc2_content,
	expectedResult: tc2_result,
}
