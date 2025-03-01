package memorypool

import (
	"fmt"
	"sync"
)

// DynamicMemoryPool gerencia pools de memória para diferentes tamanhos de bloco.
type DynamicMemoryPool struct {
	pools map[int]*sync.Pool // Mapeia tamanhos de bloco para seus respectivos pools
	mu    sync.RWMutex       // Controle de concorrência para acesso ao mapa
}

// NewDynamicMemoryPool cria um novo gerenciador de pools de memória.
func NewDynamicMemoryPool(sizes []int) *DynamicMemoryPool {
	dmp := &DynamicMemoryPool{
		pools: make(map[int]*sync.Pool),
	}

	// Inicializa pools para cada tamanho de bloco
	for _, size := range sizes {
		blockSize := size // Evita captura da variável errada em closures
		dmp.pools[blockSize] = &sync.Pool{
			New: func() interface{} {
				return make([]byte, blockSize)
			},
		}
	}

	return dmp
}

// Get retorna um bloco de memória do tamanho especificado.
func (dmp *DynamicMemoryPool) Get(size int) []byte {
	dmp.mu.RLock()
	pool, exists := dmp.pools[size]
	dmp.mu.RUnlock()

	if !exists {
		// Caso não tenha um pool do tamanho solicitado, cria dinamicamente
		dmp.mu.Lock()
		pool = &sync.Pool{
			New: func() interface{} {
				return make([]byte, size)
			},
		}
		dmp.pools[size] = pool
		dmp.mu.Unlock()
	}

	return pool.Get().([]byte)
}

// Put devolve um bloco de memória ao pool correspondente.
func (dmp *DynamicMemoryPool) Put(block []byte) {
	size := cap(block) // Tamanho real do bloco
	dmp.mu.RLock()
	pool, exists := dmp.pools[size]
	dmp.mu.RUnlock()

	if exists {
		pool.Put(block)
	} else {
		fmt.Printf("⚠️  Tentativa de devolver bloco de tamanho %d sem pool correspondente\n", size)
	}
}
