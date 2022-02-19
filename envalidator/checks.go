package envalidator

import (
	"reflect"
	"strconv"

	reader "github.com/ghaliesh/envalid/file"
	"github.com/ghaliesh/envalid/utils"
)

func checkKeyExist(envFile reader.EnvKeyValuePairs, key string) {
	exists := utils.Exists(envFile, key, true)
	if !exists {
		err := utils.KeyDoesNotExistsError(key)
		panic(err)
	}
}

func checkType(typeof reflect.Kind, key, value string) {
	switch typeof {
	case
		reflect.Int, reflect.Int8,
		reflect.Int16, reflect.Int32,
		reflect.Int64:
		checkInt(value, key)

	case
		reflect.Uint, reflect.Uint16,
		reflect.Uint8, reflect.Uint32,
		reflect.Uint64:
		checkUnitInt(value, key)

	case reflect.Float32, reflect.Float64:
		checkFloat(value, key)

	case reflect.Bool:
		checkBool(value, key)

	case reflect.String:
		checkString(value, key)

	default:
		panic("")

	}
}

func checkForError(err error, key, val, typeof string) {
	if err != nil {
		err = utils.KeyIsNotOfRightTypeError(key, val, typeof)
		panic(err)
	}
}

func checkInt(val string, key string) {
	_, err := strconv.ParseInt(val, 0, 64)
	checkForError(err, key, val, "int")
}

func checkUnitInt(val string, key string) {
	_, err := strconv.ParseUint(val, 0, 64)
	checkForError(err, key, val, "unit")
}

func checkFloat(val string, key string) {
	_, err := strconv.ParseFloat(val, 64)
	checkForError(err, key, val, "float")
}

func checkString(val string, key string) {
	if val == "true" || val == "false" {
		err := utils.KeyIsNotOfRightTypeError(key, val, "string")
		panic(err)
	}

	_, err := strconv.ParseFloat(val, 64)
	if err == nil {
		err = utils.KeyIsNotOfRightTypeError(key, val, "string")
		panic(err)
	}
}

func checkBool(val string, key string) {
	if val != "true" && val != "false" {
		err := utils.KeyIsNotOfRightTypeError(key, val, "bool")
		panic(err)
	}
}
