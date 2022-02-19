package envalidator

import "github.com/ghaliesh/envalid/utils"

type CheckTestCase struct {
	name        string
	keyValPairs dictType
	key         string
	shouldPanic bool
}

type CheckTestCases = []CheckTestCase

var key string = utils.RandomString(4)
var keyThatDontExist string = utils.RandomString(5)

var keyValPairs dictType = dictType{key: "value"}

var CheckExistTestCases = []CheckTestCase{
	{
		name:        "Key Exists",
		keyValPairs: keyValPairs,
		key:         key,
		shouldPanic: false,
	},
	{
		name:        "Key Dosn't Exists",
		keyValPairs: keyValPairs,
		key:         keyThatDontExist,
		shouldPanic: true,
	},
}
