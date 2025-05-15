package xStruct

import (
	"reflect"
)

func IsStructEmpty(s interface{}) bool {
	v := reflect.ValueOf(s)

	if v.Kind() == reflect.Ptr {
		if v.IsNil() {
			return true
		}
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		panic("IsStructEmpty: input is not a struct or pointer to struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)

		if !field.CanInterface() {
			continue
		}

		zero := reflect.Zero(field.Type()).Interface()
		if !reflect.DeepEqual(field.Interface(), zero) {
			return false
		}
	}
	return true
}
