package impl

const (
	defaultLength uint32 = 1000
)

type storage struct {
	length uint32
}

func (storage *storage) Length() uint32 {
	return storage.length
}

type Option interface {
	apply(*storage)
}

type optionFunc func(*storage)

func (f optionFunc) apply(storage *storage) {
	f(storage)
}

func WithLength(length uint32) Option {
	return optionFunc(func(storage *storage) {
		storage.length = length
	})
}

func NewStorage(options ...Option) *storage {
	storage := &storage{
		length: defaultLength,
	}
	for _, option := range options {
		option.apply(storage)
	}
	return storage
}
