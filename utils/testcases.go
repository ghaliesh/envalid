package utils

type ExistTestCase struct {
	name        string
	dict        map[string]string
	target      string
	lookForKeys bool
	shouldExist bool
}

var existsTestcases []ExistTestCase = []ExistTestCase{
	{
		name:        "Key Exist",
		dict:        map[string]string{"a": "1", "b": "2", "c": "3"},
		target:      "a",
		lookForKeys: true,
		shouldExist: true,
	},
	{
		name:        "Key Doesn't Exist",
		dict:        map[string]string{"a": "1", "b": "2", "c": "3"},
		target:      "d",
		lookForKeys: true,
		shouldExist: false,
	},
	{
		name:        "Value Exist",
		dict:        map[string]string{"a": "1", "b": "2", "c": "3"},
		target:      "1",
		lookForKeys: false,
		shouldExist: true,
	},
	{
		name:        "Value Doesn't Exist",
		dict:        map[string]string{"a": "1", "b": "2", "c": "3"},
		target:      "4",
		lookForKeys: false,
		shouldExist: false,
	},
}
