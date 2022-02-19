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

	v := reflect.ValueOf(g)
	allFields := reflect.VisibleFields(reflect.TypeOf(g))
	for i, f := range allFields {

		fieldInfo := FieldInfo{
			"key":   f.Name,
			"value": v.Field(i).Interface(),
			"type":  f.Type,
			"tags":  f.Tag,
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

		_, err := utils.FindOne(envFile, key)
		if err != nil {
			panic(err.Error())
		}

		envalue := envFile[key]
		typeof := rule["type"].(reflect.Type).Kind()

		switch typeof {
		case reflect.Int:
		case reflect.Int16:
		case reflect.Int8:
		case reflect.Int32:
		case reflect.Int64:
			fmt.Println("Integers", envalue)

		case reflect.Uint:
		case reflect.Uint8:
		case reflect.Uint16:
		case reflect.Uint32:
		case reflect.Uint64:
			fmt.Println("Unsigned integers", envalue)

		case reflect.Float32:
		case reflect.Float64:
			fmt.Println("Floats", envalue)

		case reflect.Bool:
			fmt.Println("Boolean", envalue)

		case reflect.String:
			fmt.Println("String", envalue)

		default:
			panic("")

		}
	}

}
