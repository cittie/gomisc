package gomisc

import (
	"reflect"
)

const (
	maxIterCount = 128
)

// GetFieldNamesRecursively Recursively return field names of an interface
// if input is not a struct, return type name instead
func GetFieldNamesRecursively(data interface{}) []string {
	return getFieldNamesRecur(data, 0)
}

func getFieldNamesRecur(data interface{}, iterCount int) []string {
	if data == nil {
		return nil
	}

	// avoid infinite recursive
	iterCount++
	if iterCount > maxIterCount {
		return nil
	}

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return []string{t.Name()}
	}

	n := t.NumField()

	fieldNames := make([]string, 0, n)

	for i := 0; i < n; i++ {
		switch t.Field(i).Type.Kind() {
		case reflect.Struct:
			if v.Field(i).CanAddr() {
				fieldNames = append(fieldNames, getFieldNamesRecur(v.Field(i).Addr().Interface(), iterCount)...)
			} else {
				fieldNames = append(fieldNames, getFieldNamesRecur(v.Field(i).Interface(), iterCount)...)
			}
		default:
			fieldNames = append(fieldNames, t.Field(i).Name)
		}
	}

	return fieldNames
}
