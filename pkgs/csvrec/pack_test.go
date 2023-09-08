package csvrec

import (
	"bytes"
	"encoding/json"
	"os"
	"sample/go-gorm-example/pkgs/testutils"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CSVRecord struct {
	Value1 int
	Value2 int
	Value3 int
	Value4 int
}

func (x CSVRecord) String() string {
	r, _ := json.Marshal(x)
	return string(r)
}

func TestPack(t *testing.T) {
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
