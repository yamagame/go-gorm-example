package st2rec

import (
	"sample/go-gorm-example/conv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Sub struct {
	NumberField int  `dummy:"sub_number"`
	Sub2        Sub2 `dummy:"sub2"`
}

type Sub2 struct {
	Number2Field int `dummy:"sub2_number"`
}

type Dummy struct {
	NumberField    int
	StringField    string
	NumberPtrField *uint
	StringPtrField *string
	Sub            Sub
}

func TestStructToField(t *testing.T) {
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
	ret, err := StructToRecord(dummy,
		[]string{
			".NumberField",
			".NumberPtrField",
			".StringField",
			".Sub.Sub2.Number2Field",
		})
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{10, uint(10), "name", 300}, ret)
}

func TestStructToField2(t *testing.T) {
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
	ret, err := StructToRecord(&dummy,
		[]string{
			".NumberField",
			".NumberPtrField",
			".StringField",
			".Sub.Sub2.Number2Field",
		})
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{10, uint(10), "name", 300}, ret)
}

func TestStructToField3(t *testing.T) {
	val := 100
	_, err := StructToRecord(&val,
		[]string{
			".NumberField",
			".NumberPtrField",
			".StringField",
			".Sub.Sub2.Number2Field",
		})
	assert.Error(t, err)
}
