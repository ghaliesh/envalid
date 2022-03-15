package envalidator

import (
	"fmt"
	"reflect"

	"github.com/ghaliesh/envalid/utils"
)

type checkExistTC struct {
	name        string
	keyValPairs dictType
	key         string
	shouldPanic bool
}

type checkExistTCs = []checkExistTC

var key string = utils.RandomString(4)
var keyThatDontExist string = utils.RandomString(5)

var keyValPairs dictType = dictType{key: "value"}

var checkExistTestCases = []checkExistTC{
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

type checkFunc = func(val string, key string) error
type checkFuncData struct {
	name        string
	funcInTest  checkFunc
	shouldPanic bool
}

type checkFuncsData = []checkFuncData

type checkFuncsTC struct {
	arg            string
	checkFuncsData checkFuncsData
}

type checkFuncsTCs = []checkFuncsTC

var checksTc = checkFuncsTCs{
	{
		arg: "0",
		checkFuncsData: checkFuncsData{
			{funcInTest: checkInt, shouldPanic: false, name: "Zero Is Integer"},
			{funcInTest: checkFloat, shouldPanic: false, name: "Zero Is Float"},
			{funcInTest: checkUnitInt, shouldPanic: false, name: "Zero Is UInt"},
			{funcInTest: checkBool, shouldPanic: true, name: "Zero Is Not A Bool"},
			{funcInTest: checkString, shouldPanic: true, name: "Zero is not A String"},
		},
	},
	{
		arg: "1000",
		checkFuncsData: checkFuncsData{
			{funcInTest: checkInt, shouldPanic: false, name: "Int Is Integer"},
			{funcInTest: checkFloat, shouldPanic: false, name: "Int Is Float"},
			{funcInTest: checkUnitInt, shouldPanic: false, name: "Positive Int is UInt"},
			{funcInTest: checkBool, shouldPanic: true, name: "Int Is Not Bool"},
			{funcInTest: checkString, shouldPanic: true, name: "Int Is Not String"},
		},
	},
	{
		arg: "string",
		checkFuncsData: checkFuncsData{
			{funcInTest: checkInt, shouldPanic: true, name: "String Is Not Int"},
			{funcInTest: checkFloat, shouldPanic: true, name: "String Is Not Float"},
			{funcInTest: checkUnitInt, shouldPanic: true, name: "String Is Not Unit"},
			{funcInTest: checkBool, shouldPanic: true, name: "String Is Not Bool"},
			{funcInTest: checkString, shouldPanic: false, name: "String Is String, duh"},
		},
	},
	{
		arg: "1000.1",
		checkFuncsData: checkFuncsData{
			{funcInTest: checkInt, shouldPanic: true, name: "Float Is Not Int"},
			{funcInTest: checkFloat, shouldPanic: false, name: "Float Is Float"},
			{funcInTest: checkUnitInt, shouldPanic: true, name: "Float Is Not Unit"},
			{funcInTest: checkBool, shouldPanic: true, name: "Float Is Not Bool"},
			{funcInTest: checkString, shouldPanic: true, name: "Float Is Not String"},
		},
	},
	{
		arg: "-1",
		checkFuncsData: checkFuncsData{
			{funcInTest: checkInt, shouldPanic: false, name: "Signed Number Is Int"},
			{funcInTest: checkFloat, shouldPanic: false, name: "Signed Number Is Float"},
			{funcInTest: checkUnitInt, shouldPanic: true, name: "Signed Number Is Not UInt"},
			{funcInTest: checkBool, shouldPanic: true, name: "Signed Number Is Not Bool"},
			{funcInTest: checkString, shouldPanic: true, name: "Signed Number Is Not String"},
		},
	},
	{
		arg: "true",
		checkFuncsData: checkFuncsData{
			{funcInTest: checkInt, shouldPanic: true, name: "Bool Number Is Not Int"},
			{funcInTest: checkFloat, shouldPanic: true, name: "Bool Number Is Not Float"},
			{funcInTest: checkUnitInt, shouldPanic: true, name: "Bool Number Is Not Unit"},
			{funcInTest: checkBool, shouldPanic: false, name: "Bool Number Is Bool"},
			{funcInTest: checkString, shouldPanic: true, name: "Bool Number Is Not String"},
		},
	},
}

type checkTypeTC struct {
	name            string
	key             string
	value           string
	matchingTypes   []reflect.Kind
	nonMatchingTyps [][]reflect.Kind
}

type checkTypeTCs = []checkTypeTC

var integers = []reflect.Kind{
	reflect.Int,
	reflect.Int8,
	reflect.Int32,
	reflect.Int16,
	reflect.Int64,
}

var floats = []reflect.Kind{
	reflect.Float32,
	reflect.Float64,
}

var unsignedInts = []reflect.Kind{
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
}

var bools = []reflect.Kind{
	reflect.Bool,
}

var strings = []reflect.Kind{
	reflect.String,
}

var unhandledTypes = []reflect.Kind{
	reflect.Array,
	reflect.Map,
	reflect.Chan,
	reflect.Slice,
	reflect.Func,
	reflect.Invalid,
	reflect.Interface,
	reflect.UnsafePointer,
}

var checkTypeTcs = checkTypeTCs{
	{
		name:            "String match",
		value:           utils.RandomString(5),
		key:             "key",
		matchingTypes:   strings,
		nonMatchingTyps: [][]reflect.Kind{integers, floats, bools, unhandledTypes, unsignedInts},
	},
	{
		name:            "Int match",
		value:           fmt.Sprint(utils.RandomInt()),
		key:             "key",
		matchingTypes:   integers,
		nonMatchingTyps: [][]reflect.Kind{strings, bools, unhandledTypes},
	},
	{
		name:            "Unit match",
		value:           fmt.Sprint(utils.RandomUnit()),
		key:             "key",
		matchingTypes:   unsignedInts,
		nonMatchingTyps: [][]reflect.Kind{strings, bools, unhandledTypes},
	},
	{
		name:            "Float match",
		value:           fmt.Sprint(utils.RandomFloat()),
		key:             "key",
		matchingTypes:   floats,
		nonMatchingTyps: [][]reflect.Kind{strings, bools, unhandledTypes},
	},
	{
		name:            "Bool match",
		value:           fmt.Sprint(true),
		key:             "key",
		matchingTypes:   bools,
		nonMatchingTyps: [][]reflect.Kind{integers, unhandledTypes, floats, unhandledTypes},
	},
}
