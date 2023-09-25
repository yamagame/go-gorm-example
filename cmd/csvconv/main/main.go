package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sample/go-gorm-example/pkgs/csvconv"
)

type Bool int

type CSVSubSub struct {
	Value1 string
	Value2 *Bool
}
type CSVSub struct {
	Value1 string
	Value2 []uint
	Value3 *CSVSubSub
}
type CSVRecord struct {
	Value1 *CSVSub
	Value2 []*CSVSub
}

func (x *CSVRecord) String() string {
	b, _ := json.Marshal(x)
	return string(b)
}

var mapping []*csvconv.Mapping[CSVRecord]

func init() {
	mapping = []*csvconv.Mapping[CSVRecord]{
		{"Value1", ".Value1.Value1", nil},
		{"Value2", ".Value1.Value2[0]", nil},
		{"Value3", ".Value1.Value3.Value1", nil},
		{"Value4", ".Value2[0].Value3.Value2", nil},
	}
}

func main() {
	result := []*CSVRecord{}
	for i := 0; i < 10; i++ {
		fp, _ := os.Open("./cmd/csvconv/main/testdata/sample.csv")
		defer fp.Close()
		ret, _ := csvconv.FromCSV[CSVRecord](fp, mapping)
		result = append(result, ret...)
	}
	for _, v := range result {
		fmt.Println(v)
	}

	{
		fp, _ := os.Open("./cmd/csvconv/main/testdata/sample.csv")
		defer fp.Close()
		cf := csv.NewReader(fp)
		records, _ := cf.ReadAll()
		keys := map[string]string{}
		for _, v := range mapping {
			keys[v.Column] = v.Key
		}
		header := records[0]
		for _, record := range records[1:] {
			result := CSVRecord{}
			fields := map[string]interface{}{}
			for i, v := range record {
				fields[keys[header[i]]] = v
			}
			_ = csvconv.FieldToStruct(&result, fields)
			{
				b, _ := json.Marshal(result)
				fmt.Println(string(b))
			}
		}
	}
}

// dummy := CSVRecord{}
// mapping := []string{
// 	".Value1.Value1",
// 	".Value1.Value2",
// 	".Value1.Value3.Value1",
// 	".Value2[0]",
// }
// ret, _ := csvconv.StructToField(&dummy, mapping)
// fmt.Println(ret)

// result := []*CSVRecord{}
// for i := 0; i < 100; i++ {
// 	dummy2 := &CSVRecord{}
// 	csvconv.FieldToStruct(dummy2, map[string]interface{}{
// 		".Value1.Value1":        "hello1",
// 		".Value1.Value2[0]":     "100",
// 		".Value1.Value3.Value1": "hello3",
// 		".Value2[0].Value1":     "hello4",
// 	})
// 	result = append(result, dummy2)
// }
// fmt.Println(result)
