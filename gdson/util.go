package gdson

import (
	"errors"
	"os"
	"reflect"
)

func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func getStructFields(anyStruct interface{}) (fields []reflect.StructField) {
	val := reflect.ValueOf(anyStruct)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		fields = append(fields, val.Type().Field(i))
	}
	return fields
}

func getStructFieldNames(anyStruct interface{}) (fields []string) {
	val := reflect.ValueOf(anyStruct)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	for i := 0; i < val.NumField(); i++ {
		fields = append(fields, val.Type().Field(i).Name)
	}
	return fields
}

func setStructFieldByName(anyStruct interface{}, fieldname string, value interface{}) error {
	structVal := reflect.ValueOf(anyStruct)
	if structVal.Kind() == reflect.Ptr {
		structVal = structVal.Elem()
	}
	field := structVal.FieldByName(fieldname)
	if !field.IsValid() {
		return errors.New("field name is invalid")
	}
	if !field.CanSet() {
		return errors.New("field value cannot be set")
	}

	fieldType := field.Type()
	val := reflect.ValueOf(value)
	if fieldType != val.Type() {
		return errors.New("field type does not match the value type")
	}
	field.Set(val)
	return nil
}
