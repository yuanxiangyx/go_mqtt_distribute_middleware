package utils

import (
	"encoding/json"
	"fmt"
	"mqtt_pro/schemas"
	"reflect"
	"strings"
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
		fieldName := strings.ToLower(typ.Field(i).Name)
		resultMap[fieldName] = fmt.Sprintf("%v", field.Interface())
	}
	return resultMap, nil
}

func ParserPayLoadData(data string) (head string, body string) {
	var mqString schemas.MqStringSchema
	err := json.Unmarshal([]byte(data), &mqString)
	if err == nil {
		return mqString.Header, mqString.Body
	}
	var mqSchema schemas.MqSchema
	json.Unmarshal([]byte(data), &mqSchema)
	headString, _ := json.Marshal(mqSchema.Header)
	bodyString, _ := json.Marshal(mqSchema.Body)
	return string(headString), string(bodyString)
}
