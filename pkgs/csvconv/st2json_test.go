package csvconv

import (
	"encoding/json"
	"fmt"
	"sample/go-gorm-example/conv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type flag bool

type category int32

const (
	cat_a category = iota
	cat_b
	cat_c
)

type sub1 struct {
	Sub1 int
	Sub2 string
}

type sub2 struct {
	Sub1 int
	Sub2 string
}

type csvRecord struct {
	Value1 int
	Value2 *string
	Value3 []int
	Value4 sub1
	Value5 []*sub2
	Value6 flag
	Value7 category
}

func TestJSONStructToField(t *testing.T) {
	record := csvRecord{
		Value1: 10,
		Value2: conv.ToPtr("hello"),
		Value3: []int{1, 2, 3, 4},
		Value4: sub1{
			Sub1: 20,
			Sub2: "world",
		},
		Value5: []*sub2{
			{1, "a"},
			{2, "b"},
		},
		Value6: true,
		Value7: cat_b,
	}
	ret, err := StructToField2(record, []string{
		".Value1",
		".Value2",
		".Value3[0]",
		".Value4.Sub1",
		".Value4.Sub2",
		".Value5[1].Sub2",
		".Value6",
		".Value7",
	})
	assert.NoError(t, err)
	fmt.Println(ret)
}

func TestJSONFieldToStruct2(t *testing.T) {
	var record csvRecord
	err := FieldToStruct2(&record, map[string]interface{}{
		".Value1":         10,
		".Value2":         "hello",
		".Value3[1]":      2,
		".Value3[0]":      1,
		".Value4.Sub1":    20,
		".Value4.Sub2":    "world",
		".Value5[1].Sub2": "b",
		".Value6":         true,
		".Value7":         cat_c,
	})
	assert.NoError(t, err)
	bytes, err := json.Marshal(record)
	fmt.Println(string(bytes))
}
