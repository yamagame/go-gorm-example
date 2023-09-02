package st2map

import (
	"fmt"
	"reflect"
)

type Field struct {
	Name  string
	Type  string
	Tag   string
	Value interface{}
}

func StructToMap[T any](v T, tag string) (map[string]Field, error) {
	rv := reflect.ValueOf(v) // return reflect.Value
	if rv.Type().Kind() != reflect.Struct {
		return nil, fmt.Errorf("is not struct type")
	}
	return structToMap(rv, tag), nil
}

func fieldValue(fv reflect.Value, tag string) interface{} {
	switch fv.Kind() {
	case reflect.Bool:
		return fv.Bool()
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return fv.Int()
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		return fv.Uint()
	case reflect.Float32, reflect.Float64:
		return fv.Float()
	case reflect.String:
		return fv.String()
	case reflect.Struct:
		return structToMap(fv, tag)
	case reflect.Pointer:
		v := fv.Elem() // return reflect.Value
		return fieldValue(v, tag)
	default:
		panic(fmt.Errorf("no support type %v", fv.Kind()))
	}
}

func structToMap(rv reflect.Value, tag string) map[string]Field {
	rt := rv.Type()
	h := map[string]Field{}
	if rt.Kind() != reflect.Struct {
		return h
	}
	for i := 0; i < rv.NumField(); i++ {
		field := Field{}
		f := rv.Type().Field(i)      // return reflect.StructField
		fv := rv.FieldByName(f.Name) // return reflect.Value
		field.Tag = f.Tag.Get(tag)
		field.Type = fv.Type().String()
		field.Value = fieldValue(fv, tag)
		h[f.Name] = field
	}
	return h
}
