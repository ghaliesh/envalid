package envalidator

import (
	"reflect"

	reader "github.com/ghaliesh/envalid/file"
)

type fieldInfo = map[string]interface{}
type fields = []fieldInfo

func getEnvFields(g interface{}) fields {
	var result fields

	allFields := reflect.VisibleFields(reflect.TypeOf(g))
	for _, f := range allFields {
		fieldInfo := fieldInfo{
			"key":  f.Name,
			"type": f.Type.Kind(),
			"tags": f.Tag,
		}
		result = append(result, fieldInfo)
	}

	return result
}

func Validate(validator interface{}, path string) {
	envFile := reader.ReadEnvFile(path)

	validationRules := getEnvFields(validator)

	for _, rule := range validationRules {
		key, _ := rule["key"].(string)

		checkKeyExist(envFile, key)

		envalue := envFile[key]
		typeof := rule["type"].(reflect.Kind)
		checkType(typeof, key, envalue)
	}

}
