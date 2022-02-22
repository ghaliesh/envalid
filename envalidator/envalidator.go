package envalidator

import (
	"reflect"

	reader "github.com/ghaliesh/envalid/file"
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
		key, _ := rule["key"].(string)

		checkKeyExist(envFile, key)

		envalue := envFile[key]
		typeof := rule["type"].(reflect.Kind)
		checkType(typeof, key, envalue)
	}

}
