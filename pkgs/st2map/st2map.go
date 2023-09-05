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
	kind := rv.Type().Kind()
	if kind == reflect.Pointer {
		switch rv.Elem().Kind() {
		case reflect.Struct:
			return structToMap(rv.Elem(), tag), nil
		default:
			return nil, fmt.Errorf("is not struct type")
		}
	}
	if kind == reflect.Struct {
		return structToMap(rv, tag), nil
	}
	return nil, fmt.Errorf("is not struct type")
}

func fieldValue(fv reflect.Value, tag string) interface{} {
	switch fv.Kind() {
	case reflect.Bool:
		return fv.Bool()
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return fv.Int()
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uintptr:
		return fv.Uint()
	case reflect.Float32, reflect.Float64:
		return fv.Float()
	case reflect.String:
		return fv.String()
	case reflect.Struct:
		return structToMap(fv, tag)
	case reflect.Array, reflect.Func, reflect.Map:
		return nil
	case reflect.Pointer:
		v := fv.Elem() // return reflect.Value
		return fieldValue(v, tag)
	default:
		panic(fmt.Errorf("no support type %v", fv.Kind()))
	}
}

func structToMap(rv reflect.Value, tag string) map[string]Field {
	h := map[string]Field{}
	for i := 0; i < rv.NumField(); i++ {
		field := Field{}
		f := rv.Type().Field(i)      // return reflect.StructField
		fv := rv.FieldByName(f.Name) // return reflect.Value
		field.Tag = f.Tag.Get(tag)
		field.Type = fv.Type().String()
		fval := fieldValue(fv, tag)
		if fval != nil {
			field.Value = fval
			h[f.Name] = field
		}
	}
	return h
}

type structToField[T any] struct {
	tag    string
	fields []string
}

func StructToField[T any](v T, fields []string, tag string) ([]interface{}, error) {
	t := structToField[T]{
		tag:    tag,
		fields: fields,
	}
	return t.Marchal(v)
}

func (x *structToField[T]) Marchal(v T) ([]interface{}, error) {
	rv := reflect.ValueOf(v) // return reflect.Value
	kind := rv.Type().Kind()
	if kind == reflect.Pointer {
		switch rv.Elem().Kind() {
		case reflect.Struct:
			return x.structToField(rv.Elem()), nil
		default:
			return nil, fmt.Errorf("is not struct type")
		}
	}
	if kind == reflect.Struct {
		return x.structToField(rv), nil
	}
	return nil, fmt.Errorf("is not struct type")
}

func (x *structToField[T]) fieldMap(fv reflect.Value) interface{} {
	switch fv.Kind() {
	case reflect.Bool:
		return fv.Bool()
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		return fv.Int()
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uintptr:
		return fv.Uint()
	case reflect.Float32, reflect.Float64:
		return fv.Float()
	case reflect.String:
		return fv.String()
	case reflect.Struct:
		return x.structToField(fv)
	case reflect.Array, reflect.Func, reflect.Map:
		return nil
	case reflect.Pointer:
		v := fv.Elem() // return reflect.Value
		return x.fieldMap(v)
	default:
		panic(fmt.Errorf("no support type %v", fv.Kind()))
	}
}

func (x *structToField[T]) structToField(rv reflect.Value) []interface{} {
	h := map[string]interface{}{}
	for i := 0; i < rv.NumField(); i++ {
		field := Field{}
		f := rv.Type().Field(i)      // return reflect.StructField
		fv := rv.FieldByName(f.Name) // return reflect.Value
		field.Tag = f.Tag.Get(x.tag)
		field.Type = fv.Type().String()
		fval := x.fieldMap(fv)
		if fval != nil {
			field.Value = fval
			for _, fname := range x.fields {
				if f.Name == fname {
					h[f.Name] = fval
					break
				}
			}
		}
	}
	r := []interface{}{}
	for _, k := range x.fields {
		r = append(r, h[k])
	}
	return r
}
