package csvconv

import (
	"encoding/csv"
	"fmt"
	"io"
)

type GatewayInterface[T any] interface {
	ToCSV(v interface{}) (string, error)
	FromCSV(v string) (interface{}, error)
	SetInfo(mapping []*Mapping[T], record []string, objptr *T, fields map[string]interface{})
	ColumnValue(field string) string
}

type Gateway[T any] struct {
	Self    *T
	Record  []string
	Field   map[string]interface{}
	Mapping []*Mapping[T]
}

func (x *Gateway[T]) SetInfo(mapping []*Mapping[T], record []string, objptr *T, fields map[string]interface{}) {
	x.Mapping = mapping
	x.Record = record
	x.Self = objptr
	x.Field = fields
}

func (Gateway[T]) ToCSV(v interface{}) (string, error) {
	return fmt.Sprintf("%v", v), nil
}

func (Gateway[T]) FromCSV(v string) (interface{}, error) {
	return v, nil
}

func (x *Gateway[T]) ColumnValue(column string) string {
	for i, m := range x.Mapping {
		if m.Column == column {
			return x.Record[i]
		}
	}
	return ""
}

type Mapping[T any] struct {
	Column    string
	FieldPath string
	Gateway   GatewayInterface[T]
}

type csvConstraint[T any] interface {
	*T
}

func FromCSV[T any, PT csvConstraint[T]](fp io.Reader, mapping []*Mapping[T]) ([]*T, error) {
	field := map[string]*Mapping[T]{}
	for _, v := range mapping {
		t := v
		field[v.Column] = t
	}
	reader := csv.NewReader(fp)
	reader.FieldsPerRecord = -1
	line := 0
	var header []string
	ret := []*T{}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if line == 0 {
			header = record
		} else {
			var obj T
			result := PT(&obj)
			fields := map[string]interface{}{}
			for i, h := range header {
				if val, ok := field[h]; ok {
					v := record[i]
					if val.Gateway != nil {
						var err error
						var t interface{}
						val.Gateway.SetInfo(mapping, record, &obj, fields)
						t, err = val.Gateway.FromCSV(v)
						if err != nil {
							return nil, err
						}
						fields[val.FieldPath] = t
					} else {
						fields[val.FieldPath] = v
					}
				}
			}
			if err := FieldToStruct(result, fields); err != nil {
				return nil, err
			}
			ret = append(ret, &obj)
		}
		line++
	}
	return ret, nil
}

func ToCSV[T any, PT csvConstraint[T]](records []*T, mapping []*Mapping[T], fp io.Writer) error {
	writer := csv.NewWriter(fp)
	header := []string{}
	keys := []string{}
	for _, v := range mapping {
		header = append(header, v.Column)
		keys = append(keys, v.FieldPath)
	}
	writer.Write(header)
	for _, obj := range records {
		ret, err := StructToField(obj, keys)
		if err != nil {
			return err
		}
		record := []string{}
		for i, val := range mapping {
			v := ""
			f := ret[i]
			if val.Gateway != nil {
				var err error
				val.Gateway.SetInfo(mapping, record, obj, nil)
				v, err = val.Gateway.ToCSV(f)
				if err != nil {
					return err
				}
			} else if f != nil {
				v = fmt.Sprintf("%v", f)
			}
			record = append(record, v)
		}
		writer.Write(record)
	}
	writer.Flush()
	return nil
}
