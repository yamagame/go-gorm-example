package st2rec

import (
	"fmt"
	"reflect"
	"strings"
)

type structField[T any] struct {
	values map[string]interface{}
}

// newStructField コンストラクタ
func newStructField[T any]() structField[T] {
	return structField[T]{
		values: map[string]interface{}{},
	}
}

// StrcutToField 構造体を配列レコードに変換
func StructToField[T any](v T, records []string) ([]interface{}, error) {
	t := newStructField[T]()
	return t.Marshal(v, records)
}

// Marshal 構造体を配列レコードに変換
func (x *structField[T]) Marshal(v T, records []string) ([]interface{}, error) {
	marshal := func(v T) error {
		rv := reflect.ValueOf(v)
		kind := rv.Type().Kind()
		if kind == reflect.Pointer {
			switch rv.Elem().Kind() {
			case reflect.Struct:
				x.structToField(rv.Elem(), "")
				return nil
			default:
				return fmt.Errorf("is not struct type")
			}
		}
		if kind == reflect.Struct {
			x.structToField(rv, "")
			return nil
		}
		return fmt.Errorf("is not struct type")
	}
	err := marshal(v)
	if err != nil {
		return nil, err
	}
	ret := []interface{}{}
	for _, k := range records {
		ret = append(ret, x.values[k])
	}
	return ret, nil
}

// FieldToStruct 構造体を配列レコードに変換
func FieldToStruct[T any](v T, values map[string]interface{}) error {
	t := newStructField[T]()
	t.values = values
	return t.Unmarshal(v)
}

// Unmarshal 構造体を配列レコードに変換
func (x *structField[T]) Unmarshal(v T) error {
	rv := reflect.ValueOf(v)
	kind := rv.Type().Kind()
	if kind == reflect.Pointer && rv.Elem().Kind() == reflect.Struct {
		x.fieldToStruct(rv.Elem(), "")
		return nil
	}
	return fmt.Errorf("is not struct pointer type")
}

func (x *structField[T]) fieldMap(fv reflect.Value, mark string) {
	switch fv.Kind() {
	case reflect.Bool:
		x.values[mark] = fv.Bool()
	case reflect.Int:
		x.values[mark] = int(fv.Int())
	case reflect.Int64:
		x.values[mark] = int64(fv.Int())
	case reflect.Int32:
		x.values[mark] = int32(fv.Int())
	case reflect.Int16:
		x.values[mark] = int16(fv.Int())
	case reflect.Int8:
		x.values[mark] = int8(fv.Int())
	case reflect.Uint:
		x.values[mark] = uint(fv.Uint())
	case reflect.Uint64:
		x.values[mark] = uint64(fv.Uint())
	case reflect.Uint32:
		x.values[mark] = uint32(fv.Uint())
	case reflect.Uint16:
		x.values[mark] = uint16(fv.Uint())
	case reflect.Uint8:
		x.values[mark] = uint8(fv.Uint())
	case reflect.Uintptr:
		x.values[mark] = uint(fv.Uint())
	case reflect.Float32:
		x.values[mark] = float32(fv.Float())
	case reflect.Float64:
		x.values[mark] = float64(fv.Float())
	case reflect.String:
		x.values[mark] = fv.String()
	case reflect.Struct:
		x.structToField(fv, mark)
	case reflect.Array, reflect.Func, reflect.Map:
	case reflect.Pointer:
		if !fv.IsNil() {
			v := fv.Elem()
			x.fieldMap(v, mark)
		}
	default:
		panic(fmt.Errorf("no support type %v", fv.Kind()))
	}
}

func (x *structField[T]) structToField(rv reflect.Value, mark string) interface{} {
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Type().Field(i)
		fv := rv.FieldByName(f.Name)
		x.fieldMap(fv, mark+"."+f.Name)
	}
	return nil
}

func (x *structField[T]) mapField(fv reflect.Value, mark string) {
	kind := fv.Kind()
	switch kind {
	case reflect.Bool:
		if val, ok := x.values[mark]; ok {
			fv.SetBool(val.(bool))
		}
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		if val, ok := x.values[mark]; ok {
			fv.SetInt(val.(int64))
		}
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uintptr:
		if val, ok := x.values[mark]; ok {
			fv.SetUint(val.(uint64))
		}
	case reflect.Float32, reflect.Float64:
		if val, ok := x.values[mark]; ok {
			fv.SetFloat(val.(float64))
		}
	case reflect.String:
		if val, ok := x.values[mark]; ok {
			fv.Set(reflect.ValueOf(val))
		}
	case reflect.Struct:
		x.fieldToStruct(fv, mark)
	case reflect.Array, reflect.Func, reflect.Map:
	case reflect.Pointer:
		for k, _ := range x.values {
			if strings.HasPrefix(k, mark) {
				if fv.IsNil() {
					fv.Set(reflect.New(fv.Type().Elem()))
				}
				v := fv.Elem()
				x.mapField(v, mark)
				break
			}
		}
	default:
		panic(fmt.Errorf("no support type %v", fv.Kind()))
	}
}

func (x *structField[T]) fieldToStruct(rv reflect.Value, mark string) {
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Type().Field(i)
		fv := rv.FieldByName(f.Name)
		x.mapField(fv, mark+"."+f.Name)
	}
}
