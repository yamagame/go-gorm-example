package csvconv

import (
	"encoding/csv"
	"fmt"
	"io"
)

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
			for i, column := range header {
				if g, ok := field[column]; ok {
					val := record[i]
					if g.Gateway != nil {
						var err error
						var t interface{}
						g.Gateway.SetInfo(mapping, record, &obj, fields)
						t, err = g.Gateway.FromCSV(val)
						if err != nil {
							return nil, err
						}
						fields[g.Key] = t
					} else {
						fields[g.Key] = val
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
		keys = append(keys, v.Key)
	}
	writer.Write(header)
	for _, obj := range records {
		ret, err := StructToField(obj, keys)
		if err != nil {
			return err
		}
		record := []string{}
		for i, g := range mapping {
			val := ""
			f := ret[i]
			if g.Gateway != nil {
				var err error
				g.Gateway.SetInfo(mapping, record, obj, nil)
				val, err = g.Gateway.ToCSV(f)
				if err != nil {
					return err
				}
			} else if f != nil {
				val = fmt.Sprintf("%v", f)
			}
			record = append(record, val)
		}
		writer.Write(record)
	}
	writer.Flush()
	return nil
}
