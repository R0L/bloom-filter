package storage

// 布隆过滤器存储标准
type IStorage interface {
	// 长度
	Length() uint32
	// 标记
	Mark(index uint32) (bool, error)
	// 查找
	Find(index uint32) (bool, error)
}
