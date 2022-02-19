package reader

import "fmt"

var (
	tc3_k1 = "key1"
	tc3_k2 = "key2"
	tc3_v1 = "val1=1212"
	tc3_v2 = "val2"
)
var tc3_content = fmt.Sprintf("%s=%s \n \n %s=%s", tc3_k1, tc3_v1, tc3_k2, tc3_v2)

var tc3_result EnvKeyValuePairs = EnvKeyValuePairs{
	tc3_k1: tc3_v1,
	tc3_k2: tc3_v2,
}

var tc_3 TestCase = TestCase{
	name:           "EnvWithEmptyLine",
	envFileContent: tc3_content,
	expectedResult: tc3_result,
}
