package concurrency

import "sync"

// A estrutura pode ser aprimorada para diferentes tipos de alocação
type ConcurrentMemoryPoolWithTypes[T any] struct {
	shards [16]map[int]*sync.Pool
	mu     [16]sync.RWMutex
}

func NewConcurrentMemoryPoolWithTypes[T any](sizes []int) *ConcurrentMemoryPoolWithTypes[T] {
	cmp := &ConcurrentMemoryPoolWithTypes[T]{}
	for i := 0; i < 16; i++ {
		cmp.shards[i] = make(map[int]*sync.Pool)
		for _, size := range sizes {
			blockSize := size
			cmp.shards[i][blockSize] = &sync.Pool{
				New: func() interface{} {
					return new(T)
				},
			}
		}
	}
	return cmp
}
