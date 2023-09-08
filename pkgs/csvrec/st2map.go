package csvrec

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
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

/////////////////////////////////////////////////////////////////////
// StructToField
/////////////////////////////////////////////////////////////////////

// StrcutToField 構造体を配列レコードに変換
func StructToField[T any](v T, keys []string) ([]interface{}, error) {
	t := newStructField[T]()
	return t.Marshal(v, keys)
}

// Marshal 構造体を配列レコードに変換
func (x *structField[T]) Marshal(v T, keys []string) ([]interface{}, error) {
	marshal := func(v T) error {
		rv := reflect.ValueOf(v)
		kind := rv.Type().Kind()
		if kind == reflect.Pointer {
			switch rv.Elem().Kind() {
			case reflect.Struct:
				return x.structToField(rv.Elem(), "")
			default:
				return fmt.Errorf("is not struct type")
			}
		}
		if kind == reflect.Struct {
			return x.structToField(rv, "")
		}
		return fmt.Errorf("is not struct type")
	}
	err := marshal(v)
	if err != nil {
		return nil, err
	}
	ret := []interface{}{}
	for _, k := range keys {
		ret = append(ret, x.values[k])
	}
	return ret, nil
}

func (x *structField[T]) fieldMap(fv reflect.Value, ppath string) error {
	switch fv.Kind() {
	case reflect.Bool:
		x.values[ppath] = fv.Bool()
	case reflect.Int:
		x.values[ppath] = int(fv.Int())
	case reflect.Int64:
		x.values[ppath] = int64(fv.Int())
	case reflect.Int32:
		x.values[ppath] = int32(fv.Int())
	case reflect.Int16:
		x.values[ppath] = int16(fv.Int())
	case reflect.Int8:
		x.values[ppath] = int8(fv.Int())
	case reflect.Uint:
		x.values[ppath] = uint(fv.Uint())
	case reflect.Uint64:
		x.values[ppath] = uint64(fv.Uint())
	case reflect.Uint32:
		x.values[ppath] = uint32(fv.Uint())
	case reflect.Uint16:
		x.values[ppath] = uint16(fv.Uint())
	case reflect.Uint8:
		x.values[ppath] = uint8(fv.Uint())
	case reflect.Uintptr:
		x.values[ppath] = uint(fv.Uint())
	case reflect.Float32:
		x.values[ppath] = float32(fv.Float())
	case reflect.Float64:
		x.values[ppath] = float64(fv.Float())
	case reflect.String:
		x.values[ppath] = fv.String()
	case reflect.Struct:
		return x.structToField(fv, ppath)
	case reflect.Array, reflect.Slice:
		for i := 0; i < fv.Len(); i++ {
			v := fv.Index(i)
			err := x.fieldMap(v, ppath+fmt.Sprintf("[%d]", i))
			if err != nil {
				return err
			}
		}
	case reflect.Func, reflect.Map:
	case reflect.Pointer:
		if !fv.IsNil() {
			v := fv.Elem()
			x.fieldMap(v, ppath)
		}
	default:
		return fmt.Errorf("no support type %v", fv.Kind())
	}
	return nil
}

func (x *structField[T]) structToField(rv reflect.Value, mark string) error {
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Type().Field(i)
		fv := rv.FieldByName(f.Name)
		if err := x.fieldMap(fv, mark+"."+f.Name); err != nil {
			return err
		}
	}
	return nil
}

/////////////////////////////////////////////////////////////////////
// FieldToStruct
/////////////////////////////////////////////////////////////////////

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
		return x.fieldToStruct(rv.Elem(), "")
	}
	return fmt.Errorf("is not struct pointer type")
}

func (x *structField[T]) mapField(fv reflect.Value, ppath string) error {
	kind := fv.Kind()
	switch kind {
	case reflect.Bool:
		if val, ok := x.values[ppath]; ok {
			if reflect.TypeOf(val).Kind() == reflect.String {
				val = val == "true"
			}
			fv.Set(reflect.ValueOf(val))
		}
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		if val, ok := x.values[ppath]; ok {
			if reflect.TypeOf(val).Kind() == reflect.String {
				v, err := strconv.ParseInt(val.(string), 10, 64)
				if err != nil {
					return err
				}
				switch kind {
				case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
					fv.SetInt(v)
				default:
					return fmt.Errorf("invalid int value %v", val)
				}
				break
			}
			fv.Set(reflect.ValueOf(val))
		}
	case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uintptr:
		if val, ok := x.values[ppath]; ok {
			if reflect.TypeOf(val).Kind() == reflect.String {
				v, err := strconv.ParseUint(val.(string), 10, 64)
				if err != nil {
					return err
				}
				switch kind {
				case reflect.Uint, reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8, reflect.Uintptr:
					fv.SetUint(v)
				default:
					return fmt.Errorf("invalid uint value %v", val)
				}
				break
			}
			fv.Set(reflect.ValueOf(val))
		}
	case reflect.Float32, reflect.Float64:
		if val, ok := x.values[ppath]; ok {
			if reflect.TypeOf(val).Kind() == reflect.String {
				v, err := strconv.ParseFloat(val.(string), 64)
				if err != nil {
					return err
				}
				switch kind {
				case reflect.Float32:
					val = float32(v)
				case reflect.Float64:
					val = v
				default:
					return fmt.Errorf("invalid float value %v", val)
				}
			}
			fv.Set(reflect.ValueOf(val))
		}
	case reflect.String:
		if val, ok := x.values[ppath]; ok {
			fv.Set(reflect.ValueOf(val))
		}
	case reflect.Struct:
		err := x.fieldToStruct(fv, ppath)
		if err != nil {
			return err
		}
	case reflect.Array, reflect.Slice:
		r := regexp.MustCompile(ppath + "\\[(\\d)\\](.*)")
		for k := range x.values {
			if r.MatchString(k) {
				m := r.FindStringSubmatch(k)
				idx, err := strconv.Atoi(m[1])
				if err != nil {
					return err
				}
				switch kind {
				case reflect.Array:
					if err != nil {
						return err
					}
					if len(m) > 1 {
						values := x.values
						_values := map[string]interface{}{}
						_values[m[2]] = x.values[k]
						x.values = _values
						x.mapField(fv.Index(idx), "")
						x.values = values
					}
				default:
					if len(m) > 1 {
						for fv.Len() <= idx {
							v := reflect.New(fv.Type().Elem()).Elem()
							slice := reflect.Append(fv, v)
							fv.Set(slice)
						}
						values := x.values
						_values := map[string]interface{}{}
						_values[m[2]] = x.values[k]
						x.values = _values
						x.mapField(fv.Index(idx), "")
						x.values = values
					}
				}
			}
		}
	case reflect.Func, reflect.Map:
	case reflect.Pointer:
		for k := range x.values {
			if strings.HasPrefix(k, ppath) {
				if fv.IsNil() {
					fv.Set(reflect.New(fv.Type().Elem()))
				}
				v := fv.Elem()
				x.mapField(v, ppath)
				break
			}
		}
	default:
		return fmt.Errorf("no support type %v", kind)
	}
	return nil
}

func (x *structField[T]) fieldToStruct(rv reflect.Value, mark string) error {
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		fieldvalue := rv.FieldByName(field.Name)
		err := x.mapField(fieldvalue, mark+"."+field.Name)
		if err != nil {
			return err
		}
	}
	return nil
}
