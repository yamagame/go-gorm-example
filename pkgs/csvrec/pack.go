package csvrec

import (
	"encoding/csv"
	"fmt"
	"io"
)

type CSVRecordConstraint[T any] interface {
	*T
}

func CSVToStruct[T any, PT CSVRecordConstraint[T]](fp io.Reader, mapping [][]string) ([]*T, error) {
	field := map[string]string{}
	for _, v := range mapping {
		field[v[0]] = v[1]
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
					mapping[val] = record[i]
				}
			}
			FieldToStruct(result, mapping)
			ret = append(ret, &v)
		}
		line++
	}
	return ret, nil
}

func StructToCSV[T any, PT CSVRecordConstraint[T]](records []*T, mapping [][]string, fp io.Writer) error {
	writer := csv.NewWriter(fp)
	header := []string{}
	keys := []string{}
	for _, v := range mapping {
		header = append(header, v[0])
		keys = append(keys, v[1])
	}
	writer.Write(header)
	for _, r := range records {
		ret, err := StructToField(r, keys)
		if err != nil {
			return err
		}
		record := []string{}
		for _, r := range ret {
			record = append(record, fmt.Sprintf("%v", r))
		}
		writer.Write(record)
	}
	writer.Flush()
	return nil
}
