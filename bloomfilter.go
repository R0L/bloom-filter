package bloomfilter

import (
	"errors"

	"github.com/R0L/bloom-filter/calculate"
	"github.com/R0L/bloom-filter/storage"
)

const (
	ErrAddCalculateErr   = "add calculates is not empty"
	ErrCheckCalculateErr = "check calculates is not empty"
)

// 布隆过滤器
type BloomFilter struct {
	length     uint32
	storage    storage.IStorage
	calculates []calculate.Calculator
}

// 创建布隆过滤器
func NewBloomFilter(storage storage.IStorage, adapters ...calculate.Calculator) *BloomFilter {
	return &BloomFilter{
		length:     storage.Length(),
		storage:    storage,
		calculates: adapters,
	}
}

// 添加值
func (bf *BloomFilter) Add(bytes []byte) error {
	if len(bf.calculates) == 0 {
		return errors.New(ErrAddCalculateErr)
	}

	for _, bfCalculate := range bf.calculates {
		index, err := bfCalculate.Calculate(bytes, bf.length)
		if nil != err {
			return err
		}
		if _, err = bf.storage.Mark(index); nil != err {
			return err
		}
	}
	return nil
}

// 检查值
func (bf *BloomFilter) Check(bytes []byte) error {
	if len(bf.calculates) == 0 {
		return errors.New(ErrCheckCalculateErr)
	}

	for _, bfCalculate := range bf.calculates {
		index, err := bfCalculate.Calculate(bytes, bf.length)
		if nil != err {
			return err
		}
		flag, err := bf.storage.Find(index)
		if nil != err {
			return err
		}
		if !flag {
			return errors.New("not exists")
		}
	}
	return nil
}
