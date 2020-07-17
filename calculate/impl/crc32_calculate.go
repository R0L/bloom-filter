package impl

import (
	"hash/crc32"
)

const (
	ErrCrc32HashWriteErr = "calculate hash err: hash write err"
)

// CRC32
type crc32Calculate struct {
}

func NewCrc32Calculate() *crc32Calculate {
	return &crc32Calculate{}
}

func (fnvC *crc32Calculate) Calculate(bytes []byte, length uint32) (uint32, error) {
	return crc32.ChecksumIEEE(bytes) % length, nil
}
