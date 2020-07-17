package impl

import (
	"os"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestFileStorage_Mark(t *testing.T) {

	tests := []struct {
		Index uint32
		Flag  bool
		Err   string
	}{
		{100, false, ErrMarkIndexParam},
		{10, true, ""},
		{20, true, ""},
	}

	fs := NewFileStorage(100)

	for _, test := range tests {
		flag, err := fs.Mark(test.Index)
		if nil != err {
			assert.EqualError(t, err, test.Err, "err value")
		} else {
			assert.Nil(t, err)
		}
		assert.Equal(t, test.Flag, flag, "flag value")
	}

}

func TestFileStorage_Find(t *testing.T) {
	tests := []struct {
		Index uint32
		Flag  bool
		Err   error
	}{
		{100, false, errors.New(ErrFindIndexParam)},
		{10, true, nil},
		{11, false, nil},
	}

	fs := NewFileStorage(100)

	fs.Mark(10)
	fs.Mark(12)

	for _, test := range tests {
		flag, err := fs.Find(test.Index)
		if nil != err {
			assert.EqualError(t, err, test.Err.Error(), "err value")
		} else {
			assert.Nil(t, err)
		}
		assert.Equal(t, test.Flag, flag, "flag value")
	}
}
