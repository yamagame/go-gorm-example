package csvrec

import (
	"bytes"
	"os"
	"sample/go-gorm-example/pkgs/testutils"
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
	mapping := [][]string{
		{"Value1", ".Value1"},
		{"Value2", ".Value2"},
		{"Value3", ".Value3"},
		{"Value4", ".Value4"},
	}
	fp, _ := os.Open("./testdata/sample.csv")
	defer fp.Close()
	ret, err := CSVToStruct[CSVRecord](fp, mapping)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = StructToCSV(ret, mapping, &buf)
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
	mapping := [][]string{
		{"Value1", ".Value1"},
		{"Value2", ".Value2"},
		{"Value3", ".Value3"},
		{"Value4", ".Value4"},
	}
	fp, _ := os.Open("./testdata/sample.csv")
	defer fp.Close()
	ret, err := CSVToStruct[CSVRecord](fp, mapping)
	assert.NoError(t, err)
	var buf bytes.Buffer
	err = StructToCSV(ret, mapping, &buf)
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
	mapping := [][]string{
		{"Value1", ".Value1"},
		{"Value2", ".Value3[1]"},
		{"Value3", ".Sub1.Value1"},
		{"Value4", ".Sub1.Value2"},
		{"Value5", ".Sub2[0].Value1"},
		{"Value6", ".MissingValue"},
		{"Value7", ".Sub2[0].Value2"},
	}
	fp, _ := os.Open("./testdata/sample.csv")
	defer fp.Close()
	ret, err := CSVToStruct[CSVRecord](fp, mapping)
	assert.NoError(t, err)

	var buf bytes.Buffer
	err = StructToCSV(ret, mapping, &buf)
	assert.NoError(t, err)

	testutils.EqualSnapshot(t, buf.Bytes(), "sample2.csv.out")
	// testutils.SaveSnapshot(t, buf.Bytes(), "sample2.csv.out")

	assert.NoError(t, err)
}
