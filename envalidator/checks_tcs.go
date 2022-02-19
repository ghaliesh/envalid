package envalidator

import (
	"github.com/ghaliesh/envalid/utils"
)

type CheckExistTC struct {
	name        string
	keyValPairs dictType
	key         string
	shouldPanic bool
}

type CheckExistTCs = []CheckExistTC

var key string = utils.RandomString(4)
var keyThatDontExist string = utils.RandomString(5)

var keyValPairs dictType = dictType{key: "value"}

var CheckExistTestCases = []CheckExistTC{
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

type CheckPanicIfErrorTC struct {
	name        string
	err         error
	val         string
	key         string
	typeof      string
	shouldPanic bool
}

var panicVal = utils.RandomString(5)
var panicKey = utils.RandomString(5)
var panicTypeof = utils.RandomString(5)
var panicTCError = utils.KeyIsNotOfRightTypeError(key, panicVal, panicTypeof)
var panicIfErrorTCs = []CheckPanicIfErrorTC{
	{name: "Should Not Panic", err: nil, key: "", val: "", typeof: "", shouldPanic: false},
	{name: "Should Panic", err: panicTCError, key: panicKey, val: panicVal, shouldPanic: true},
}
