package csvreader

import (
	"reflect"
	"strconv"
	"time"
)

func setField(field reflect.Value, valStr string) error {
	if !field.CanSet() {
		return nil
	}
	switch field.Kind() {
	case reflect.Bool:
		if val, err := strconv.ParseBool(valStr); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}
	case reflect.Int:
		if val, err := strconv.ParseInt(valStr, 10, 64); err == nil {
			field.Set(reflect.ValueOf(int(val)).Convert(field.Type()))
		}
	case reflect.Int8:
		if val, err := strconv.ParseInt(valStr, 10, 8); err == nil {
			field.Set(reflect.ValueOf(int8(val)).Convert(field.Type()))
		}
	case reflect.Int16:
		if val, err := strconv.ParseInt(valStr, 10, 16); err == nil {
			field.Set(reflect.ValueOf(int16(val)).Convert(field.Type()))
		}
	case reflect.Int32:
		if val, err := strconv.ParseInt(valStr, 10, 32); err == nil {
			field.Set(reflect.ValueOf(int32(val)).Convert(field.Type()))
		}
	case reflect.Int64:
		if val, err := time.ParseDuration(valStr); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		} else if val, err := strconv.ParseInt(valStr, 10, 64); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}
	case reflect.Uint:
		if val, err := strconv.ParseUint(valStr, 10, 64); err == nil {
			field.Set(reflect.ValueOf(uint(val)).Convert(field.Type()))
		}
	case reflect.Uint8:
		if val, err := strconv.ParseUint(valStr, 10, 8); err == nil {
			field.Set(reflect.ValueOf(uint8(val)).Convert(field.Type()))
		}
	case reflect.Uint16:
		if val, err := strconv.ParseUint(valStr, 10, 16); err == nil {
			field.Set(reflect.ValueOf(uint16(val)).Convert(field.Type()))
		}
	case reflect.Uint32:
		if val, err := strconv.ParseUint(valStr, 10, 32); err == nil {
			field.Set(reflect.ValueOf(uint32(val)).Convert(field.Type()))
		}
	case reflect.Uint64:
		if val, err := strconv.ParseUint(valStr, 10, 64); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}
	case reflect.Uintptr:
		if val, err := strconv.ParseUint(valStr, 10, 64); err == nil {
			field.Set(reflect.ValueOf(uintptr(val)).Convert(field.Type()))
		}
	case reflect.Float32:
		if val, err := strconv.ParseFloat(valStr, 32); err == nil {
			field.Set(reflect.ValueOf(float32(val)).Convert(field.Type()))
		}
	case reflect.Float64:
		if val, err := strconv.ParseFloat(valStr, 64); err == nil {
			field.Set(reflect.ValueOf(val).Convert(field.Type()))
		}
	case reflect.String:
		field.Set(reflect.ValueOf(valStr).Convert(field.Type()))
	case reflect.Ptr:
		field.Set(reflect.New(field.Type().Elem()))
	}
	return nil
}
