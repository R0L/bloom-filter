package impl

import (
	"github.com/pkg/errors"
)

// 文件
type fileStorage struct {
	*storage
	bF []int
}

func NewFileStorage(length uint32) *fileStorage {
	return &fileStorage{
		storage: NewStorage(WithLength(length)),
		bF:      make([]int, length),
	}
}

func (fs *fileStorage) Mark(index uint32) (bool, error) {

	if index >= fs.length {
		return false, errors.New(ErrMarkIndexParam)
	}

	fs.bF[index] = 1

	return true, nil
}

func (fs *fileStorage) Find(index uint32) (bool, error) {

	if index >= fs.length {
		return false, errors.New(ErrFindIndexParam)
	}

	if fs.bF[index] == 1 {
		return true, nil
	}

	return false, nil
}
