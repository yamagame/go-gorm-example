package csvconv

import (
	"encoding/csv"
	"fmt"
	"io"
)

type Gateway struct {
	ToCSV    func(v interface{}) (string, error)
	ToStruct func(v string) (interface{}, error)
}

type Mapping struct {
	Column    string
	FieldPath string
	Gateway   *Gateway
}

type csvConstraint[T any] interface {
	*T
}

func Read[T any, PT csvConstraint[T]](fp io.Reader, mapping []Mapping) ([]*T, error) {
	field := map[string]*Mapping{}
	for _, v := range mapping {
		t := v
		field[v.Column] = &t
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
			var v T
			result := PT(&v)
			mapping := map[string]interface{}{}
			for i, h := range header {
				if val, ok := field[h]; ok {
					v := record[i]
					if val.Gateway != nil && val.Gateway.ToStruct != nil {
						var err error
						var t interface{}
						t, err = val.Gateway.ToStruct(v)
						if err != nil {
							return nil, err
						}
						mapping[val.FieldPath] = t
					} else {
						mapping[val.FieldPath] = v
					}
				}
			}
			err := FieldToStruct(result, mapping)
			if err != nil {
				return nil, err
			}
			ret = append(ret, &v)
		}
		line++
	}
	return ret, nil
}

func Write[T any, PT csvConstraint[T]](records []*T, mapping []Mapping, fp io.Writer) error {
	writer := csv.NewWriter(fp)
	header := []string{}
	keys := []string{}
	for _, v := range mapping {
		header = append(header, v.Column)
		keys = append(keys, v.FieldPath)
	}
	writer.Write(header)
	for _, r := range records {
		ret, err := StructToField(r, keys)
		if err != nil {
			return err
		}
		record := []string{}
		for i, m := range mapping {
			v := ""
			f := ret[i]
			if f != nil {
				if m.Gateway != nil && m.Gateway.ToCSV != nil {
					var err error
					v, err = m.Gateway.ToCSV(f)
					if err != nil {
						return err
					}
				} else {
					v = fmt.Sprintf("%v", f)
				}
			}
			record = append(record, v)
		}
		writer.Write(record)
	}
	writer.Flush()
	return nil
}
