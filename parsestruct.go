package gotils

import (
  "reflect"
)

func ParseStruct(data map[string]interface{}, result interface{}) (interface{}, error) {
  if data == nil {
    return nil, &errorString{"No input!"}
  }
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
	return result, nil
}
