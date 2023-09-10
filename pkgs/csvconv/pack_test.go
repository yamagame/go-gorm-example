package csvconv

import (
	"bytes"
	"fmt"
	"os"
	"sample/go-gorm-example/pkgs/testutils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPack1(t *testing.T) {
	type CSVRecord struct {
		Value1 int
		Value2 int
		Value3 int
		Value4 int
	}
	mapping := []*Mapping[CSVRecord]{
		{"Value1", ".Value1", nil},
		{"Value2", ".Value2", nil},
		{"Value3", ".Value3", nil},
		{"Value4", ".Value4", nil},
	}
	fp, _ := os.Open("./testdata/sample.csv")
	defer fp.Close()
	ret, err := FromCSV[CSVRecord](fp, mapping)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = ToCSV(ret, mapping, &buf)
	assert.NoError(t, err)

	testutils.EqualSnapshot(t, buf.Bytes(), "sample.csv.out")
	// testutils.SaveSnapshot(t, buf.Bytes(), "sample.csv.out")

	assert.NoError(t, err)
}

func TestPack2(t *testing.T) {
	type CSVRecord struct {
		Value1 *uint
		Value2 *int32
		Value3 *uint32
		Value4 *string
	}
	mapping := []*Mapping[CSVRecord]{
		{"Value1", ".Value1", nil},
		{"Value2", ".Value2", nil},
		{"Value3", ".Value3", nil},
		{"Value4", ".Value4", nil},
	}
	fp, _ := os.Open("./testdata/sample.csv")
	defer fp.Close()
	ret, err := FromCSV[CSVRecord](fp, mapping)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = ToCSV(ret, mapping, &buf)
	assert.NoError(t, err)

	testutils.EqualSnapshot(t, buf.Bytes(), "sample.csv.out")
	// testutils.SaveSnapshot(t, buf.Bytes(), "sample.csv.out")

	assert.NoError(t, err)
}

func TestPack3(t *testing.T) {
	type CSVInt int32
	type CSVSub struct {
		Value1 CSVInt
		Value2 *CSVInt
	}
	type CSVRecord struct {
		Value1 int
		Value2 int
		Value3 [3]int
		Sub1   CSVSub
		Sub2   []CSVSub
	}
	mapping := []*Mapping[CSVRecord]{
		{"Value1", ".Value1", nil},
		{"Value2", ".Value3[1]", nil},
		{"Value3", ".Sub1.Value1", nil},
		{"Value4", ".Sub1.Value2", nil},
		{"Value5", ".Sub2[0].Value1", nil},
		{"Value6", ".MissingValue", nil},
		{"Value7", ".Sub2[0].Value2", nil},
	}
	fp, _ := os.Open("./testdata/sample.csv")
	defer fp.Close()
	ret, err := FromCSV[CSVRecord](fp, mapping)
	assert.NoError(t, err)

	var buf bytes.Buffer
	err = ToCSV(ret, mapping, &buf)
	assert.NoError(t, err)

	testutils.EqualSnapshot(t, buf.Bytes(), "sample2.csv.out")
	// testutils.SaveSnapshot(t, buf.Bytes(), "sample2.csv.out")

	assert.NoError(t, err)
}

type CSVRecord struct {
	Value1 int
	Value2 int
	Value3 int
	Value4 int
	Str1   string
	Str2   string
	Value6 int
}

type sampleLabel struct {
	Form string
	Gateway[CSVRecord]
}

func (x *sampleLabel) ToCSV(v interface{}) (string, error) {
	return fmt.Sprintf(x.Form, v), nil
}

func (x *sampleLabel) FromCSV(v string) (interface{}, error) {
	var r int
	_, err := fmt.Sscanf(v, x.Form, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func newSampleLabel(form string) GatewayInterface[CSVRecord] {
	return &sampleLabel{
		Form: form,
	}
}

type addStr1AndStr2 struct {
	Gateway[CSVRecord]
}

func (x *addStr1AndStr2) ToCSV(v interface{}) (string, error) {
	return fmt.Sprintf("%s_%s", x.Self.Str1, x.Self.Str2), nil
}

func (x *addStr1AndStr2) FromCSV(v string) (interface{}, error) {
	s := strings.Split(v, "_")
	x.Field[".Str1"] = s[0]
	x.Field[".Str2"] = s[1]
	return s, nil
}

type addValue1AndValue2 struct {
	Gateway[CSVRecord]
}

func (x *addValue1AndValue2) ToCSV(v interface{}) (string, error) {
	return fmt.Sprintf("%d", x.Self.Value1+x.Self.Value2), nil
}

func (x *addValue1AndValue2) FromCSV(v string) (interface{}, error) {
	t := x.ColumnValue("Value1")
	return t, nil
}

func TestPack4(t *testing.T) {
	mapping := []*Mapping[CSVRecord]{
		{"Value1", ".Value1", newSampleLabel("%02d")},
		{"Value2", ".Value2", newSampleLabel("%04d")},
		{"Value3", ".Value3", newSampleLabel("%02x")},
		{"Value4", ".Value4", newSampleLabel("%d")},
		{"Value5", ".Value5", &addStr1AndStr2{}},
		{"Value6", ".Value6", &addValue1AndValue2{}},
	}
	fp, _ := os.Open("./testdata/sample3.csv")
	defer fp.Close()
	ret, err := FromCSV[CSVRecord](fp, mapping)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = ToCSV(ret, mapping, &buf)
	assert.NoError(t, err)

	testutils.EqualSnapshot(t, buf.Bytes(), "sample3.csv.out")
	// testutils.SaveSnapshot(t, buf.Bytes(), "sample3.csv.out")

	assert.NoError(t, err)
}
