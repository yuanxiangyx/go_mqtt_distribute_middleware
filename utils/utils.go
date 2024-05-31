package utils

import (
	"fmt"
	"reflect"
)

func StructToMapString(schema interface{}) (map[string]string, error) {
	val := reflect.ValueOf(schema)
	if val.Kind() != reflect.Struct {
		return nil, fmt.Errorf("toMap: %v is not a struct", val)
	}
	typ := val.Type()

	resultMap := make(map[string]string)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Name
		resultMap[fieldName] = fmt.Sprintf("%v", field.Interface())
	}
	return resultMap, nil
}
