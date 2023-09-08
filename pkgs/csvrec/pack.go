package csvrec

import (
	"encoding/csv"
	"fmt"
	"io"
)

type CSVRecordConstraint[T any] interface {
	*T
}

func CSVToStruct[T any, PT CSVRecordConstraint[T]](fp io.Reader, field map[string]string) ([]*T, error) {
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

func StructToCSV[T any, PT CSVRecordConstraint[T]](records []*T, fields []string, fp io.Writer) error {
	writer := csv.NewWriter(fp)
	for _, r := range records {
		ret, err := StructToField(r, fields)
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
