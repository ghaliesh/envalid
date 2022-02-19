package envalidator

import (
	"strconv"

	"github.com/ghaliesh/envalid/utils"
)

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
