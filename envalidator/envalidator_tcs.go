package envalidator

import "reflect"

type GetEnvFieldsTestCases struct {
	name   string
	arg    interface{}
	result Fields
}

type tc1ArgType struct {
	A int
	B uint
	c string
	m float32
	N float64
	K int32
}

var tc1Results Fields = Fields{
	FieldInfo{"key": "A", "type": reflect.Int, "tags": ""},
	FieldInfo{"key": "B", "type": reflect.Uint, "tags": ""},
	FieldInfo{"key": "c", "type": reflect.String, "tags": ""},
	FieldInfo{"key": "m", "type": reflect.Float32, "tags": ""},
	FieldInfo{"key": "N", "type": reflect.Float64, "tags": ""},
	FieldInfo{"key": "K", "type": reflect.Int32, "tags": ""},
}

var getEnvFieldsTC []GetEnvFieldsTestCases = []GetEnvFieldsTestCases{
	{
		name: "Happy Case",
		// c and m are populated to remove warning
		arg:    tc1ArgType{c: "", m: 1.1},
		result: tc1Results,
	},
}
