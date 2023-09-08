package csvrec

import (
	"fmt"
	"sample/go-gorm-example/conv"
	"sample/go-gorm-example/pkgs/testutils"
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

func TestStructToField1(t *testing.T) {
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
	ret, err := StructToField(dummy,
		[]string{
			".NumberField",
			".NumberPtrField",
			".StringField",
			".Sub.Sub2.Number2Field",
		})
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{
		10,
		uint(10),
		"name",
		300,
	}, ret)
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
	ret, err := StructToField(&dummy,
		[]string{
			".NumberField",
			".NumberPtrField",
			".StringField",
			".Sub.Sub2.Number2Field",
		})
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{
		10,
		uint(10),
		"name",
		300,
	}, ret)
}

func TestStructToField3(t *testing.T) {
	val := 100
	_, err := StructToField(&val,
		[]string{
			".NumberField",
			".NumberPtrField",
			".StringField",
			".Sub.Sub2.Number2Field",
		})
	assert.Error(t, err)
}

func TestStructToField4(t *testing.T) {
	dummy := Dummy{
		NumberPtrField: conv.UintP(10),
		// StringPtrField: conv.StrP(""),
	}
	err := FieldToStruct(&dummy,
		map[string]interface{}{
			".NumberField":           10,
			".NumberPtrField":        uint(10),
			".StringField":           "name",
			".Sub.Sub2.Number2Field": 300,
		})
	assert.NoError(t, err)
	ret, err := StructToField(&dummy,
		[]string{
			".NumberField",
			".NumberPtrField",
			".StringField",
			".Sub.Sub2.Number2Field",
		})
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{
		10,
		uint(10),
		"name",
		300,
	}, ret)
}

func TestStructToField5(t *testing.T) {
	type Dummy2 struct {
		Bool    bool
		Int     int
		Int64   int64
		Int32   int32
		Int16   int16
		Int8    int8
		Uint    uint
		Uint64  uint64
		Uint32  uint32
		Uint16  uint16
		Uint8   uint8
		Float32 float32
		Float64 float64
		String  string

		BoolPtr    *bool
		IntPtr     *int
		Int64Ptr   *int64
		Int32Ptr   *int32
		Int16Ptr   *int16
		Int8Ptr    *int8
		UintPtr    *uint
		Uint64Ptr  *uint64
		Uint32Ptr  *uint32
		Uint16Ptr  *uint16
		Uint8Ptr   *uint8
		Float32Ptr *float32
		Float64Ptr *float64
		StringPtr  *string
	}

	Bool := true
	Int := int(10)
	Int64 := int64(10)
	Int32 := int32(10)
	Int16 := int16(10)
	Int8 := int8(10)
	Uint := uint(10)
	Uint64 := uint64(10)
	Uint32 := uint32(10)
	Uint16 := uint16(10)
	Uint8 := uint8(10)
	Float32 := float32(10)
	Float64 := float64(10)
	String := "Hello"
	dummy := Dummy2{
		Bool:    Bool,
		Int:     Int,
		Int64:   Int64,
		Int32:   Int32,
		Int16:   Int16,
		Int8:    Int8,
		Uint:    Uint,
		Uint64:  Uint64,
		Uint32:  Uint32,
		Uint16:  Uint16,
		Uint8:   Uint8,
		Float32: Float32,
		Float64: Float64,
		String:  String,

		BoolPtr:    &Bool,
		IntPtr:     &Int,
		Int64Ptr:   &Int64,
		Int32Ptr:   &Int32,
		Int16Ptr:   &Int16,
		Int8Ptr:    &Int8,
		UintPtr:    &Uint,
		Uint64Ptr:  &Uint64,
		Uint32Ptr:  &Uint32,
		Uint16Ptr:  &Uint16,
		Uint8Ptr:   &Uint8,
		Float32Ptr: &Float32,
		Float64Ptr: &Float64,
		StringPtr:  &String,
	}
	err := FieldToStruct(&dummy,
		map[string]interface{}{
			".Bool":       false,
			".Int":        int(20),
			".Int64":      int64(20),
			".Int32":      int32(20),
			".Int16":      int16(20),
			".Int8":       int8(20),
			".Uint":       uint(20),
			".Uint64":     uint64(20),
			".Uint32":     uint32(20),
			".Uint16":     uint16(20),
			".Uint8":      uint8(20),
			".Float32":    float32(20),
			".Float64":    float64(20),
			".String":     "world",
			".BoolPtr":    false,
			".IntPtr":     int(21),
			".Int64Ptr":   int64(22),
			".Int32Ptr":   int32(23),
			".Int16Ptr":   int16(24),
			".Int8Ptr":    int8(25),
			".UintPtr":    uint(26),
			".Uint64Ptr":  uint64(27),
			".Uint32Ptr":  uint32(28),
			".Uint16Ptr":  uint16(29),
			".Uint8Ptr":   uint8(30),
			".Float32Ptr": float32(31),
			".Float64Ptr": float64(32),
			".StringPtr":  "world",
		})
	assert.NoError(t, err)
	ret, err := StructToField(&dummy,
		[]string{
			".Bool",
			".Int",
			".Int64",
			".Int32",
			".Int16",
			".Int8",
			".Uint",
			".Uint64",
			".Uint32",
			".Uint16",
			".Uint8",
			".Float32",
			".Float64",
			".String",
			".BoolPtr",
			".IntPtr",
			".Int64Ptr",
			".Int32Ptr",
			".Int16Ptr",
			".Int8Ptr",
			".UintPtr",
			".Uint64Ptr",
			".Uint32Ptr",
			".Uint16Ptr",
			".Uint8Ptr",
			".Float32Ptr",
			".Float64Ptr",
			".StringPtr",
		})
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{
		false,
		int(20),
		int64(20),
		int32(20),
		int16(20),
		int8(20),
		uint(20),
		uint64(20),
		uint32(20),
		uint16(20),
		uint8(20),
		float32(20),
		float64(20),
		"world",
		false,
		int(21),
		int64(22),
		int32(23),
		int16(24),
		int8(25),
		uint(26),
		uint64(27),
		uint32(28),
		uint16(29),
		uint8(30),
		float32(31),
		float64(32),
		"world",
	}, ret)
}

func TestStructToField6(t *testing.T) {
	one := uint(1)
	two := uint(2)
	three := uint(3)
	type CSVSub struct {
		Value1 string
		Value2 uint
	}
	type CSVRecord struct {
		Value1 [3]string
		Value2 []*uint
		Value3 []int
		Value4 []CSVSub
	}
	dummy := CSVRecord{
		Value1: [3]string{"1", "2", "3"},
		Value2: []*uint{&one, &two, &three},
		Value3: []int{1, 2, 3},
		Value4: []CSVSub{{Value1: "A", Value2: 1}, {Value1: "B", Value2: 2}},
	}
	mapping := []string{
		".Value1[0]", ".Value1[1]", ".Value1[2]",
		".Value2[0]", ".Value2[1]", ".Value2[2]",
		".Value3[0]", ".Value3[1]", ".Value3[2]",
		".Value4[0].Value1", ".Value4[0].Value2",
		".Value4[1].Value1", ".Value4[1].Value2",
	}
	ret, err := StructToField(&dummy, mapping)
	assert.NoError(t, err)

	testutils.EqualSnapshot(t, []byte(fmt.Sprintf("%v", ret)), "csv-record.out")
	// testutils.SaveSnapshot(t, []byte(fmt.Sprintf("%v", ret)), "csv-record.out")

	var dummy2 CSVRecord
	err = FieldToStruct(&dummy2, map[string]interface{}{
		".Value1[0]":        "hello",
		".Value1[1]":        "world",
		".Value2[0]":        three,
		".Value2[1]":        two,
		".Value2[2]":        one,
		".Value3[0]":        1,
		".Value3[1]":        2,
		".Value3[2]":        3,
		".Value4[0].Value1": "B",
		".Value4[0].Value2": uint(2),
		".Value4[1].Value1": "C",
		".Value4[1].Value2": uint(3),
	})
	assert.NoError(t, err)

	ret2, err := StructToField(&dummy2, mapping)
	assert.NoError(t, err)

	testutils.EqualSnapshot(t, []byte(fmt.Sprintf("%v", ret2)), "csv-record2.out")
	// testutils.SaveSnapshot(t, []byte(fmt.Sprintf("%v", ret2)), "csv-record2.out")
}
