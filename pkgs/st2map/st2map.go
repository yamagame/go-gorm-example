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

func StructToMap[T any](v T, tag string) map[string]Field {
	rv := reflect.ValueOf(v) // return reflect.Value
	return structToMap(rv, tag)
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
		switch fv.Kind() {
		case reflect.Bool:
			field.Value = fv.Bool()
		case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
			field.Value = fv.Int()
		case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
			field.Value = fv.Uint()
		case reflect.Float32, reflect.Float64:
			field.Value = fv.Float()
		case reflect.String:
			field.Value = fv.String()
		case reflect.Struct:
			field.Value = structToMap(fv, tag)
		case reflect.Pointer:
			v := fv.Elem() // return reflect.Value
			switch v.Kind() {
			case reflect.Bool:
				field.Value = v.Bool()
			case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
				field.Value = v.Int()
			case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
				field.Value = v.Uint()
			case reflect.Float32, reflect.Float64:
				field.Value = v.Float()
			case reflect.String:
				field.Value = v.String()
			default:
				break
			}
		default:
			fmt.Println(fv.String())
		}
		h[f.Name] = field
	}
	return h
}
