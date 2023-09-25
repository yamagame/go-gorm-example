package csvconv

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

/////////////////////////////////////////////////////////////////////
// StructToField
/////////////////////////////////////////////////////////////////////

// StructToField 構造体を配列レコードに変換
func StructToField[T any](v T, keys []string) ([]interface{}, error) {
	values := map[string]interface{}{}
	return marshal(v, keys, values)
}

// marshal 構造体を配列レコードに変換
func marshal[T any](v T, keys []string, values map[string]interface{}) ([]interface{}, error) {
	marshal := func(v T) error {
		rv := reflect.ValueOf(v)
		kind := rv.Type().Kind()
		if kind == reflect.Pointer {
			switch rv.Elem().Kind() {
			case reflect.Struct:
				return structToField[T](rv.Elem(), "", values)
			default:
				return fmt.Errorf("is not struct type")
			}
		}
		if kind == reflect.Struct {
			return structToField[T](rv, "", values)
		}
		return fmt.Errorf("is not struct type")
	}
	err := marshal(v)
	if err != nil {
		return nil, err
	}
	ret := []interface{}{}
	for _, k := range keys {
		ret = append(ret, values[k])
	}
	return ret, nil
}

func fieldMap[T any](fv reflect.Value, ppath string, values map[string]interface{}) error {
	switch fv.Kind() {
	case reflect.Bool:
		values[ppath] = fv.Bool()
	case reflect.Int:
		values[ppath] = int(fv.Int())
	case reflect.Int64:
		values[ppath] = fv.Int()
	case reflect.Int32:
		values[ppath] = int32(fv.Int())
	case reflect.Int16:
		values[ppath] = int16(fv.Int())
	case reflect.Int8:
		values[ppath] = int8(fv.Int())
	case reflect.Uint:
		values[ppath] = uint(fv.Uint())
	case reflect.Uint64:
		values[ppath] = fv.Uint()
	case reflect.Uint32:
		values[ppath] = uint32(fv.Uint())
	case reflect.Uint16:
		values[ppath] = uint16(fv.Uint())
	case reflect.Uint8:
		values[ppath] = uint8(fv.Uint())
	case reflect.Uintptr:
		values[ppath] = uint(fv.Uint())
	case reflect.Float32:
		values[ppath] = float32(fv.Float())
	case reflect.Float64:
		values[ppath] = fv.Float()
	case reflect.String:
		values[ppath] = fv.String()
	case reflect.Struct:
		return structToField[T](fv, ppath, values)
	case reflect.Array, reflect.Slice:
		for i := 0; i < fv.Len(); i++ {
			v := fv.Index(i)
			if err := fieldMap[T](v, ppath+fmt.Sprintf("[%d]", i), values); err != nil {
				return err
			}
		}
	case reflect.Func, reflect.Map:
	case reflect.Pointer:
		if !fv.IsNil() {
			v := fv.Elem()
			if err := fieldMap[T](v, ppath, values); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("no support type %v", fv.Kind())
	}
	return nil
}

func structToField[T any](rv reflect.Value, mark string, values map[string]interface{}) error {
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Type().Field(i)
		fv := rv.FieldByName(f.Name)
		if err := fieldMap[T](fv, mark+"."+f.Name, values); err != nil {
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
	return unmarshal[T](v, values)
}

// unmarshal 構造体を配列レコードに変換
func unmarshal[T any](v T, values map[string]interface{}) error {
	rv := reflect.ValueOf(v)
	kind := rv.Type().Kind()
	if kind == reflect.Pointer && rv.Elem().Kind() == reflect.Struct {
		return fieldToStruct[T](rv.Elem(), "", values)
	}
	return fmt.Errorf("is not struct pointer type")
}

func mapField[T any](fv reflect.Value, ppath string, values map[string]interface{}) error {
	kind := fv.Kind()
	switch kind {
	case reflect.Bool:
		if val, ok := values[ppath]; ok {
			if reflect.TypeOf(val).Kind() == reflect.String {
				val = val == "true"
			}
			fv.Set(reflect.ValueOf(val))
		}
	case reflect.Int, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		if val, ok := values[ppath]; ok {
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
		if val, ok := values[ppath]; ok {
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
		if val, ok := values[ppath]; ok {
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
		if val, ok := values[ppath]; ok {
			fv.Set(reflect.ValueOf(val))
		}
	case reflect.Struct:
		err := fieldToStruct[T](fv, ppath, values)
		if err != nil {
			return err
		}
	case reflect.Array, reflect.Slice:
		r := regexp.MustCompile(ppath + "\\[(\\d)\\](.*)")
		for k := range values {
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
						_values := map[string]interface{}{}
						_values[m[2]] = values[k]
						if err := mapField[T](fv.Index(idx), "", _values); err != nil {
							return err
						}
					}
				default:
					if len(m) > 1 {
						for fv.Len() <= idx {
							v := reflect.New(fv.Type().Elem()).Elem()
							slice := reflect.Append(fv, v)
							fv.Set(slice)
						}
						_values := map[string]interface{}{}
						_values[m[2]] = values[k]
						if err := mapField[T](fv.Index(idx), "", _values); err != nil {
							return err
						}
					}
				}
			}
		}
	case reflect.Func, reflect.Map:
	case reflect.Pointer:
		for k := range values {
			if strings.HasPrefix(k, ppath) {
				if fv.IsNil() {
					fv.Set(reflect.New(fv.Type().Elem()))
				}
				v := fv.Elem()
				if err := mapField[T](v, ppath, values); err != nil {
					return err
				}
				break
			}
		}
	default:
		return fmt.Errorf("no support type %v", kind)
	}
	return nil
}

func fieldToStruct[T any](rv reflect.Value, mark string, values map[string]interface{}) error {
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Type().Field(i)
		fieldvalue := rv.FieldByName(field.Name)
		if err := mapField[T](fieldvalue, mark+"."+field.Name, values); err != nil {
			return err
		}
	}
	return nil
}
