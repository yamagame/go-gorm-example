package st2map

import (
	"fmt"
	"sample/go-gorm-example/conv"
	"testing"
)

type Sub struct {
	NumberField int  `dummy:"sub_number"`
	Sub2        Sub2 `dummy:"sub2"`
}

type Sub2 struct {
	Number2Field int `dummy:"sub2_number"`
}

type Dummy struct {
	NumberField    int     `dummy:"number"`
	StringField    string  `dummy:"string"`
	NumberPtrField *uint   `dummy:"number_ptr"`
	StringPtrField *string `dummy:"string_ptr"`
	Sub            Sub     `dummy:"sub"`
}

func TestStructToMap(t *testing.T) {
	dummy := Dummy{
		NumberField:    10,
		StringField:    "name",
		NumberPtrField: conv.UintP(10),
		StringPtrField: conv.StrP("name"),
		Sub: Sub{
			NumberField: 200,
			Sub2: Sub2{
				Number2Field: 300,
			},
		},
	}
	h := StructToMap(dummy, "dummy")
	for k, v := range h {
		fmt.Println(k, v)
	}
}
