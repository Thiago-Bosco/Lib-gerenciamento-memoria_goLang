package memorypool

import (
	"hash/fnv"
	"strconv"
	"sync"
	"time"
)

// shardCount define quantos pools serão usados para minimizar contenção.
const shardCount = 16

// ConcurrentMemoryPool gerencia pools de memória de maneira concorrente, evitando contenção.
type ConcurrentMemoryPool struct {
	shards  [shardCount]map[int]*sync.Pool
	mu      [shardCount]sync.RWMutex
	Metrics *MemoryMetrics // Tornando o campo 'Metrics' público
}

// NewConcurrentMemoryPool inicializa o pool concorrente para diferentes tamanhos de bloco e as métricas.
func NewConcurrentMemoryPool(sizes []int) *ConcurrentMemoryPool {
	cmp := &ConcurrentMemoryPool{
		Metrics: NovaMemoryMetrics(), // Inicializa métricas
	}

	// Inicializa os pools dentro de cada shard
	for i := 0; i < shardCount; i++ {
		cmp.shards[i] = make(map[int]*sync.Pool)
		for _, size := range sizes {
			blockSize := size // Evita captura incorreta da variável
			cmp.shards[i][blockSize] = &sync.Pool{
				New: func() interface{} {
					return make([]byte, blockSize)
				},
			}
		}
	}

	return cmp
}

// getShardIndex calcula um índice baseado no tamanho para determinar qual shard usar.
func getShardIndex(size int) int {
	h := fnv.New32a()
	h.Write([]byte(strconv.Itoa(size)))
	return int(h.Sum32()) % shardCount
}

// Get retorna um bloco de memória do tamanho especificado.
func (cmp *ConcurrentMemoryPool) Get(size int) []byte {
	start := time.Now() // Marca o início da alocação

	shardIndex := getShardIndex(size)

	cmp.mu[shardIndex].RLock()
	pool, exists := cmp.shards[shardIndex][size]
	cmp.mu[shardIndex].RUnlock()

	if !exists {
		// Criando dinamicamente um novo pool se necessário
		cmp.mu[shardIndex].Lock()
		pool = &sync.Pool{
			New: func() interface{} {
				return make([]byte, size)
			},
		}
		cmp.shards[shardIndex][size] = pool
		cmp.mu[shardIndex].Unlock()
	}

	block := pool.Get().([]byte)

	// Medir o tempo de alocação
	elapsed := time.Since(start)
	cmp.Metrics.RegistrarAlocacao()

	// Log do tempo de alocação
	cmp.Metrics.AdicionarTempoDeAlocacao(elapsed)
	return block
}

// Put devolve um bloco ao pool correspondente.
func (cmp *ConcurrentMemoryPool) Put(block []byte) {
	start := time.Now() // Marca o início da reutilização

	size := cap(block)
	shardIndex := getShardIndex(size)

	cmp.mu[shardIndex].RLock()
	pool, exists := cmp.shards[shardIndex][size]
	cmp.mu[shardIndex].RUnlock()

	if exists {
		pool.Put(block)
		// Atualiza as métricas
		cmp.Metrics.RegistrarReutilizacao()

		// Medir o tempo de reutilização
		elapsed := time.Since(start)
		cmp.Metrics.AdicionarTempoDeReutilizacao(elapsed)
	}
}
