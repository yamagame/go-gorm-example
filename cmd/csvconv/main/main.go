package main

import (
	"fmt"
	"os"
	"sample/go-gorm-example/pkgs/csvconv"
)

type CSVSubSub struct {
	Value1 string
	Value2 uint
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

func main() {
	result := []*CSVRecord{}
	for i := 0; i < 10; i++ {
		mapping := []*csvconv.Mapping[CSVRecord]{
			{"Value1", ".Value1.Value1", nil},
			{"Value2", ".Value1.Value2[0]", nil},
			{"Value3", ".Value1.Value3.Value1", nil},
			{"Value4", ".Value2[0].Value1", nil},
		}
		fp, _ := os.Open("./cmd/csvconv/main/testdata/sample.csv")
		defer fp.Close()
		ret, _ := csvconv.FromCSV[CSVRecord](fp, mapping)
		result = append(result, ret...)
	}
	fmt.Println(result)
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
