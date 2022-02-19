package envalidator

import (
	"fmt"
	"reflect"

	reader "github.com/ghaliesh/envalid/file"
	"github.com/ghaliesh/envalid/utils"
)

type FieldInfo = map[string]interface{}
type Fields = []FieldInfo

func getEnvFields(g interface{}) Fields {
	var result Fields

	allFields := reflect.VisibleFields(reflect.TypeOf(g))
	for _, f := range allFields {
		fieldInfo := FieldInfo{
			"key":  f.Name,
			"type": f.Type.Kind(),
			"tags": f.Tag,
		}
		result = append(result, fieldInfo)
	}

	return result
}

func Validate(validator interface{}, path string) {
	reader := reader.EnvFileReader{}
	envFile := reader.ReadEnvFile(path)

	validationRules := getEnvFields(validator)

	for _, rule := range validationRules {
		key, ok := rule["key"].(string)
		if !ok {
			panic(fmt.Errorf("all .env file keys should be strings"))
		}

		exists := utils.Exists(envFile, key, true)
		if !exists {
			err := utils.KeyDoesNotExistsError(key)
			panic(err)
		}

		envalue := envFile[key]
		typeof := rule["type"].(reflect.Kind)

		switch typeof {
		case
			reflect.Int, reflect.Int8,
			reflect.Int16, reflect.Int32,
			reflect.Int64:
			checkInt(envalue, key)

		case
			reflect.Uint, reflect.Uint16,
			reflect.Uint8, reflect.Uint32,
			reflect.Uint64:
			checkUnitInt(envalue, key)

		case reflect.Float32, reflect.Float64:
			checkFloat(envalue, key)

		case reflect.Bool:
			checkBool(envalue, key)

		case reflect.String:
			checkString(envalue, key)

		default:
			panic("")

		}
	}

}
