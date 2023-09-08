package testutils

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

// EqualSnapshot スナップショットと一致比較
func EqualSnapshot(t *testing.T, b []byte, fname string) {
	fpath := filepath.Join("./testdata/", fname)
	rp, err := os.Open(fpath)
	assert.NoError(t, err)
	data, err := io.ReadAll(rp)
	assert.NoError(t, err)
	assert.Equal(t, data, b)
}

// EqualSnapshot スナップショットを保存
func SaveSnapshot(t *testing.T, b []byte, fname string) {
	err := os.MkdirAll("./testdata/", 0777)
	assert.NoError(t, err)
	fpath := filepath.Join("./testdata/", fname)
	os.WriteFile(fpath, b, 0666)
}
