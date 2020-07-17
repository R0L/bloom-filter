package calculate

// 布隆过滤器计算标准
type Calculator interface {
	// 计算
	Calculate(bytes []byte, length uint32) (uint32, error)
}
