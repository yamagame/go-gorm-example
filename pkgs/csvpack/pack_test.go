package csvpack

import (
	"encoding/json"
	"fmt"
	"os"
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
	fp, _ := os.Open("./testdata/sample.csv")
	defer fp.Close()
	ret, err := CSVToStruct[CSVRecord](fp, map[string]string{
		"Value1": ".Value1",
		"Value2": ".Value2",
		"Value3": ".Value3",
	})
	assert.NoError(t, err)
	fmt.Println(ret)
	StructToCSV(ret, []string{
		".Value1",
		".Value2",
		".Value3",
		".Value4",
	}, os.Stdout)
}
