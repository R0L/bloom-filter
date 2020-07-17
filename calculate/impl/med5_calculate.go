package impl

import (
	"crypto/md5"

	"github.com/pkg/errors"
)

// MED5
type md5Calculate struct {
}

func NewMd5Calculate() *md5Calculate {
	return &md5Calculate{}
}

func (fnvC *md5Calculate) Calculate(bytes []byte, length uint32) (uint32, error) {
	md5Hash := md5.New()
	_, err := md5Hash.Write(bytes)
	if nil != err {
		return 0, errors.Wrap(err, ErrHashWriteErr)
	}
	hashByte := md5Hash.Sum(nil)

	return (uint32)((hashByte[0]<<24)+(hashByte[1]<<16)+(hashByte[2]<<8)+hashByte[3]) / length, nil

}
