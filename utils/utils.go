package utils

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ParserPayLoadDataToMap(data string) (map[string]string, error) {
	var mqSchema map[string]interface{}
	err := json.Unmarshal([]byte(data), &mqSchema)
	mqSchemaString := make(map[string]string)
	for key, value := range mqSchema {
		val := reflect.ValueOf(value)
		switch val.Kind() {
		case reflect.String:
			mqSchemaString[key] = val.String()
		default:
			jsonData, err := json.Marshal(val.Interface())
			if err != nil {
				fmt.Println("err:", err)
			}
			mqSchemaString[key] = string(jsonData)
		}
	}
	return mqSchemaString, err
}

func ParserPayLoadDataToString(data string) (head string, body string) {
	mqSchemaString, _ := ParserPayLoadDataToMap(data)
	return mqSchemaString["header"], mqSchemaString["body"]
}
