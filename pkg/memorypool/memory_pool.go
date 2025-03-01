package memorypool

import (
	"sync"
)

// MemoryPool representa um pool de blocos de memória reutilizáveis.
type MemoryPool struct {
	pool sync.Pool
}

// NewMemoryPool cria um novo pool de memória com blocos de um tamanho fixo.
func NewMemoryPool(blockSize int) *MemoryPool {
	return &MemoryPool{
		pool: sync.Pool{
			New: func() interface{} {
				return make([]byte, blockSize) // Aloca um bloco de memória do tamanho definido
			},
		},
	}
}

// Get retorna um bloco de memória do pool.
func (mp *MemoryPool) Get() []byte {
	return mp.pool.Get().([]byte)
}

// Put devolve um bloco de memória ao pool.
func (mp *MemoryPool) Put(block []byte) {
	mp.pool.Put(block)
}
