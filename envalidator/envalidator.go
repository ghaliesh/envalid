package envalidator

import (
	"reflect"
)

type FieldInfo = map[string]interface{}
type Fields = []FieldInfo

func getEnvFields(g interface{}) Fields {
	var result Fields

	v := reflect.ValueOf(g)
	typeOfV := v.Type()
	length := v.NumField()

	for i := 0; i < length; i++ {
		field := typeOfV.Field(i)
		val := v.Field(i).Interface()

		key := field.Name
		typeof := field.Type.String()
		tags := field.Tag

		fieldInfo := FieldInfo{
			"key":   key,
			"value": val,
			"type":  typeof,
			"tags":  tags,
		}
		result = append(result, fieldInfo)
	}

	return result
}
