package memorypool

import (
	"sync/atomic"
	"time"
)

// MemoryMetrics mantém informações sobre o uso de memória.
type MemoryMetrics struct {
	Allocations         int32
	Reutilizacoes       int32
	BlocosAtivos        int32
	TempoDeAlocacao     int64 // Tempo em nanossegundos.
	TempoDeReutilizacao int64 // Tempo em nanossegundos.
}

// NewMemoryMetrics cria uma nova instância de MemoryMetrics.
func NewMemoryMetrics() *MemoryMetrics {
	return &MemoryMetrics{}
}

// RegisterAllocation incrementa a contagem de alocações e o número de blocos ativos.
func (m *MemoryMetrics) RegisterAllocation() {
	atomic.AddInt32(&m.Allocations, 1)
	atomic.AddInt32(&m.BlocosAtivos, 1)
}

// RegisterReuse incrementa a contagem de reutilizações e decrementa o número de blocos ativos.
func (m *MemoryMetrics) RegisterReuse() {
	atomic.AddInt32(&m.Reutilizacoes, 1)
	atomic.AddInt32(&m.BlocosAtivos, -1)
}

// AddAllocationTime acumula o tempo total gasto em alocações.
func (m *MemoryMetrics) AddAllocationTime(duration time.Duration) {
	atomic.AddInt64(&m.TempoDeAlocacao, duration.Nanoseconds())
}

// AddReuseTime acumula o tempo total gasto em reutilizações.
func (m *MemoryMetrics) AddReuseTime(duration time.Duration) {
	atomic.AddInt64(&m.TempoDeReutilizacao, duration.Nanoseconds())
}

// GetMetrics retorna as métricas de memória de maneira eficiente.
func (m *MemoryMetrics) GetMetrics() (allocations, reuses, activeBlocks int, allocationTime, reuseTime time.Duration) {
	allocations = int(atomic.LoadInt32(&m.Allocations))
	reuses = int(atomic.LoadInt32(&m.Reutilizacoes))
	activeBlocks = int(atomic.LoadInt32(&m.BlocosAtivos))
	allocationTime = time.Duration(atomic.LoadInt64(&m.TempoDeAlocacao))
	reuseTime = time.Duration(atomic.LoadInt64(&m.TempoDeReutilizacao))
	return
}
