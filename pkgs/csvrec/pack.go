package csvrec

import (
	"encoding/csv"
	"fmt"
	"io"
)

type CSVGateway struct {
	ToFile   func(v interface{}) (string, error)
	ToStruct func(v string) (string, error)
}

type CSVMapping struct {
	Column    string
	FieldPath string
	Gateway   *CSVGateway
}

type csvConstraint[T any] interface {
	*T
}

func CSVToStruct[T any, PT csvConstraint[T]](fp io.Reader, mapping []CSVMapping) ([]*T, error) {
	field := map[string]*CSVMapping{}
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
						v, err = val.Gateway.ToStruct(v)
						if err != nil {
							return nil, err
						}
					}
					mapping[val.FieldPath] = v
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

func StructToCSV[T any, PT csvConstraint[T]](records []*T, mapping []CSVMapping, fp io.Writer) error {
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
			r := ret[i]
			if r != nil {
				if m.Gateway != nil && m.Gateway.ToFile != nil {
					var err error
					r, err = m.Gateway.ToFile(r)
					if err != nil {
						return err
					}
				}
				v = fmt.Sprintf("%v", r)
			}
			record = append(record, v)
		}
		writer.Write(record)
	}
	writer.Flush()
	return nil
}
