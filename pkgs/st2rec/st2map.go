package st2rec

import (
	"fmt"
	"reflect"
)

type structToRecords[T any] struct {
	values map[string]interface{}
}

// newStructToRecord コンストラクタ
func newStructToRecord[T any]() structToRecords[T] {
	return structToRecords[T]{
		values: map[string]interface{}{},
	}
}

// StrcutToField 構造体を配列レコードに変換
func StructToRecord[T any](v T, records []string) ([]interface{}, error) {
	t := newStructToRecord[T]()
	return t.Marshal(v, records)
}

// Marshal 構造体を配列レコードに変換
func (x *structToRecords[T]) Marshal(v T, records []string) ([]interface{}, error) {
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

func (x *structToRecords[T]) fieldMap(fv reflect.Value, mark string) {
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
		v := fv.Elem()
		x.fieldMap(v, mark)
	default:
		panic(fmt.Errorf("no support type %v", fv.Kind()))
	}
}

func (x *structToRecords[T]) structToField(rv reflect.Value, mark string) interface{} {
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Type().Field(i)
		fv := rv.FieldByName(f.Name)
		x.fieldMap(fv, mark+"."+f.Name)
	}
	return nil
}
