package envalidator

import (
	"reflect"
	"strconv"

	reader "github.com/ghaliesh/envalid/file"
	"github.com/ghaliesh/envalid/utils"
)

type dictType = reader.EnvKeyValuePairs

func checkKeyExist(dict dictType, key string) error {
	exists := utils.Exists(dict, key, true)
	if !exists {
		err := utils.KeyDoesNotExistsError(key)
		return err
	}

	return nil
}

func checkType(typeof reflect.Kind, key, value string) error {
	switch typeof {
	case
		reflect.Int, reflect.Int8,
		reflect.Int16, reflect.Int32,
		reflect.Int64:
		return checkInt(value, key)

	case
		reflect.Uint, reflect.Uint16,
		reflect.Uint8, reflect.Uint32,
		reflect.Uint64:
		return checkUnitInt(value, key)

	case reflect.Float32, reflect.Float64:
		return checkFloat(value, key)

	case reflect.Bool:
		return checkBool(value, key)

	case reflect.String:
		return checkString(value, key)

	default:
		return utils.ErrUnhandledTypes
	}
}

func checkInt(val string, key string) error {
	_, err := strconv.ParseInt(val, 0, 64)
	if err != nil {
		return utils.KeyIsNotOfRightTypeError(key, val, "int")
	}
	return nil
}

func checkUnitInt(val string, key string) error {
	_, err := strconv.ParseUint(val, 0, 64)
	if err != nil {
		return utils.KeyIsNotOfRightTypeError(key, val, "unit")
	}
	return nil
}

func checkFloat(val string, key string) error {
	_, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return utils.KeyIsNotOfRightTypeError(key, val, "float")
	}

	return nil
}

func checkString(val string, key string) error {

	if val == "true" || val == "false" {
		return utils.KeyIsNotOfRightTypeError(key, val, "string")
	}

	_, err := strconv.ParseFloat(val, 64)
	if err == nil {
		return utils.KeyIsNotOfRightTypeError(key, val, "string")
	}

	return nil
}

func checkBool(val string, key string) error {
	if val != "true" && val != "false" {
		return utils.KeyIsNotOfRightTypeError(key, val, "bool")
	}

	return nil
}
