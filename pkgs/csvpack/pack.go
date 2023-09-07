package csvpack

import (
	"encoding/csv"
	"fmt"
	"io"
	"sample/go-gorm-example/pkgs/st2rec"
	"strconv"
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
				v, err := strconv.ParseInt(record[i], 10, 64)
				if err != nil {
					return nil, err
				}
				if val, ok := field[h]; ok {
					mapping[val] = v
				}
			}
			st2rec.FieldToStruct(result, mapping)
			ret = append(ret, &v)
		}
		line++
	}
	return ret, nil
}

func StructToCSV[T any, PT CSVRecordConstraint[T]](records []*T, fields []string, fp io.Writer) error {
	writer := csv.NewWriter(fp)
	for _, r := range records {
		ret, err := st2rec.StructToField(r, fields)
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
