package impl

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFnvCalculate_Calculate(t *testing.T) {
	tests := []struct {
		HashByte []byte
		Err      error
		Index    uint32
	}{
		{[]byte("测试一个"), nil, 37},
		{[]byte(""), nil, 61},
	}

	for _, test := range tests {
		index, err := NewFnvCalculate().Calculate(test.HashByte, 100)
		assert.Equal(t, test.Err, err, "err value")
		assert.Equal(t, test.Index, index, "hash value")
	}
}

func TestMd5Calculate_Calculate(t *testing.T) {
	tests := []struct {
		HashByte []byte
		Err      error
		Index    uint32
	}{
		{[]byte("测试一个"), nil, 1},
		{[]byte(""), nil, 2},
	}

	for _, test := range tests {
		index, err := NewMd5Calculate().Calculate(test.HashByte, 100)
		assert.Equal(t, test.Err, err, "err value")
		assert.Equal(t, test.Index, index, "hash value")
	}
}
