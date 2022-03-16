package envalidator

import (
	"reflect"

	reader "github.com/ghaliesh/envalid/file"
)

type OnErrorHandler = func(err error)

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

func Validate(validator interface{}, path string) error {
	envFile, err := reader.ReadEnvFile(path)
	if err != nil {
		return err
	}

	validationRules := getEnvFields(validator)

	for _, rule := range validationRules {
		key, _ := rule["key"].(string)

		err := checkKeyExist(envFile, key)
		if err != nil {
			return err
		}

		envalue := envFile[key]
		typeof := rule["type"].(reflect.Kind)

		err = checkType(typeof, key, envalue)
		if err != nil {
			return err
		}
	}

	return nil
}
