package envalidator

import (
	"reflect"
)

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

type tc2Types struct {
	K1 int
	K2 string
	K3 bool
}

type tc3Types struct {
	K1 uint8
	K2 string
	K3 bool
}

type tc4Types struct {
	K1 uint16
	K2 string
	K3 bool
}

type tc5Types struct {
	K1 uint32
	K2 string
	K3 bool
}

type tc6Types struct {
	K1 uint64
	K2 string
	K3 bool
}

type tc7Types struct {
	K1 int8
	K2 string
	K3 bool
}

type tc8Types struct {
	K1 int32
	K2 string
	K3 bool
}

type tc9Types struct {
	K1 int64
	K2 string
	K3 bool
}

type tc10Types struct {
	K1 float32
	K2 string
	K3 bool
}

type tc11Types struct {
	K1 float64
	K2 string
	K3 bool
}

var validateTCs = []struct {
	name        string
	envFile     string
	validators  []interface{}
	path        string
	shouldPanic bool
}{
	{
		name: "Should Not Panic", validators: []interface{}{
			tc2Types{}, tc3Types{}, tc4Types{}, tc5Types{}, tc6Types{},
			tc7Types{}, tc8Types{}, tc9Types{}, tc10Types{}, tc11Types{},
		},
		envFile:     "K1=100 \n K2=name \n K3=true",
		shouldPanic: false,
	},
	{
		name: "Should Panic due to missing types", validators: []interface{}{tc2Types{}},
		envFile:     "K1=100 \n K2=name",
		shouldPanic: true,
	},
	{
		name: "Should Panic due to wrong types", validators: []interface{}{tc2Types{}},
		envFile:     "K1=100 \n K2=10",
		shouldPanic: true,
	},
}
