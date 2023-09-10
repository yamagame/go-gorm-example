package csvconv

import (
	"bytes"
	"sample/go-gorm-example/pkgs/testutils"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func TestGateway(t *testing.T) {
	type SampleRecord struct {
		Value1 string
		Value2 string
	}

	newSampleRecordStatic := func(s string) GatewayInterface[SampleRecord] {
		return &StaticString[SampleRecord]{Value: s}
	}

	//////////////////////////////////////////////////
	// マッピング
	//////////////////////////////////////////////////

	mapping := []*Mapping[SampleRecord]{
		{"Value1", ".Value1", &EmptyString[SampleRecord]{}},
		{"Value2", ".Value2", &ConvString[SampleRecord]{
			To: func(o *SampleRecord) (string, error) {
				return cases.Title(language.Und, cases.NoLower).String(o.Value2), nil
			},
			From: func(v string) (string, error) {
				return strings.ToLower(v), nil
			},
		}},
		{"Value3", "", newSampleRecordStatic("something")},
	}

	//////////////////////////////////////////////////
	// 構造体
	//////////////////////////////////////////////////

	obj := []*SampleRecord{
		{
			Value1: "Hello",
			Value2: "world",
		},
	}

	//////////////////////////////////////////////////
	// 構造体からCSV
	//////////////////////////////////////////////////

	var buf bytes.Buffer
	err := ToCSV(obj, mapping, &buf)
	assert.NoError(t, err)

	testutils.EqualSnapshot(t, buf.Bytes(), "gateway-test1.csv")
	// testutils.SaveSnapshot(t, buf.Bytes(), "gateway-test1.csv")

	//////////////////////////////////////////////////
	// CSVから構造体
	//////////////////////////////////////////////////

	ret, err := FromCSV[SampleRecord](&buf, mapping)
	assert.NoError(t, err)
	assert.Equal(t, "world", ret[0].Value2)
}
