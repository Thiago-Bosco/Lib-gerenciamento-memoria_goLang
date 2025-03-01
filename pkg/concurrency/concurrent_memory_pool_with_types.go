package concurrency

import (
	"sync"
	"time"
)

// Estrutura de pool concorrente genérico com shards para reduzir contenção
type ConcurrentMemoryPoolWithTypes[T any] struct {
	shards  [16]map[int]*sync.Pool
	mu      [16]sync.RWMutex
	Metrics Metrics // Estrutura de métricas
}

// Estrutura para armazenar as métricas
type Metrics struct {
	Allocations   int64
	Reuses        int64
	ActiveBlocks  int64
	AllocDuration time.Duration
	ReuseDuration time.Duration
}

// Novo pool concorrente para tipos genéricos
func NewConcurrentMemoryPoolWithTypes[T any](sizes []int) *ConcurrentMemoryPoolWithTypes[T] {
	cmp := &ConcurrentMemoryPoolWithTypes[T]{}
	for i := 0; i < 16; i++ {
		cmp.shards[i] = make(map[int]*sync.Pool)
		for _, size := range sizes {
			blockSize := size
			cmp.shards[i][blockSize] = &sync.Pool{
				New: func() interface{} {
					return make([]T, blockSize) // Agora aloca um slice do tipo T
				},
			}
		}
	}
	return cmp
}

// Obtém um bloco de memória do tamanho especificado
func (cmp *ConcurrentMemoryPoolWithTypes[T]) Get(size int) []T {
	shardIndex := size % 16 // Estratégia de distribuição
	cmp.mu[shardIndex].RLock()
	pool, exists := cmp.shards[shardIndex][size]
	cmp.mu[shardIndex].RUnlock()

	if exists {
		start := time.Now()
		block := pool.Get().([]T)
		cmp.Metrics.AllocDuration += time.Since(start) // Atualiza tempo de alocação
		cmp.Metrics.Allocations++
		return block
	}
	return make([]T, size) // Se não houver no pool, aloca novo bloco
}

// Retorna um bloco ao pool
func (cmp *ConcurrentMemoryPoolWithTypes[T]) Put(block []T) {
	size := len(block)
	shardIndex := size % 16

	cmp.mu[shardIndex].RLock()
	pool, exists := cmp.shards[shardIndex][size]
	cmp.mu[shardIndex].RUnlock()

	if exists {
		start := time.Now()
		pool.Put(block)
		cmp.Metrics.ReuseDuration += time.Since(start) // Atualiza tempo de reutilização
		cmp.Metrics.Reuses++
	}
}

// Método para acessar as métricas
func (cmp *ConcurrentMemoryPoolWithTypes[T]) GetMetrics() Metrics {
	return cmp.Metrics
}
