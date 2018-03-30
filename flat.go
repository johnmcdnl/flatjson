package flatjson

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Flatten transforms a json array into it's keys
func Flatten(s string) (map[string]string, error) {
	return FlattenBytes([]byte(s))
}

// FlattenBytes flattens a json byte array
func FlattenBytes(b []byte) (map[string]string, error) {
	return flattenJSON(b)
}

func flattenJSON(b []byte) (map[string]string, error) {
	var m map[string]interface{}
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}

	return flattenJSONMap(m)
}
func flattenJSONMap(m map[string]interface{}) (map[string]string, error) {

	result := make(map[string]string)

	for k, raw := range m {
		flatten(result, k, reflect.ValueOf(raw))
	}

	return result, nil
}

func flatten(result map[string]string, prefix string, v reflect.Value) {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Bool:
		if v.Bool() {
			result[prefix] = "true"
		} else {
			result[prefix] = "false"
		}
	case reflect.Float64:
		result[prefix] = fmt.Sprintf("%v", v.Float())
	case reflect.String:
		result[prefix] = v.String()
	case reflect.Map:
		flattenMap(result, prefix, v)
	case reflect.Slice:
		flattenSlice(result, prefix, v)

	}
}

func flattenMap(result map[string]string, prefix string, v reflect.Value) {
	for _, k := range v.MapKeys() {
		flatten(result, fmt.Sprintf("%s.%s", prefix, k.String()), v.MapIndex(k))
	}
}

func flattenSlice(result map[string]string, prefix string, v reflect.Value) {
	prefix = prefix + "."

	result[prefix+"#"] = fmt.Sprintf("%d", v.Len())
	for i := 0; i < v.Len(); i++ {
		flatten(result, fmt.Sprintf("%s%d", prefix, i), v.Index(i))
	}
}
