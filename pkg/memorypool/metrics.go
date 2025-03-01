package memorypool

import (
	"sync"
	"sync/atomic"
	"time"
)

// MemoryMetrics mantém informações sobre o uso de memória.
type MemoryMetrics struct {
	Allocations         int32
	Reutilizacoes       int32
	BlocosAtivos        int32
	TempoDeAlocacao     int64 // Usando int64 para armazenar tempo em nanossegundos.
	TempoDeReutilizacao int64 // Usando int64 para armazenar tempo em nanossegundos.
	mu                  sync.RWMutex
}

// NovaMemoryMetrics cria uma nova instância de MemoryMetrics.
func NovaMemoryMetrics() *MemoryMetrics {
	return &MemoryMetrics{}
}

// RegistrarAlocacao incrementa a contagem de alocações e o número de blocos ativos.
func (m *MemoryMetrics) RegistrarAlocacao() {
	atomic.AddInt32(&m.Allocations, 1)
	atomic.AddInt32(&m.BlocosAtivos, 1)
}

// RegistrarReutilizacao incrementa a contagem de reutilizações e decrementa o número de blocos ativos.
func (m *MemoryMetrics) RegistrarReutilizacao() {
	atomic.AddInt32(&m.Reutilizacoes, 1)
	atomic.AddInt32(&m.BlocosAtivos, -1)
}

// AdicionarTempoDeAlocacao acumula o tempo total gasto em alocações.
func (m *MemoryMetrics) AdicionarTempoDeAlocacao(duration time.Duration) {
	atomic.AddInt64(&m.TempoDeAlocacao, duration.Nanoseconds())
}

// AdicionarTempoDeReutilizacao acumula o tempo total gasto em reutilizações.
func (m *MemoryMetrics) AdicionarTempoDeReutilizacao(duration time.Duration) {
	atomic.AddInt64(&m.TempoDeReutilizacao, duration.Nanoseconds())
}

// ObterMetricas retorna as métricas de memória de maneira eficiente.
func (m *MemoryMetrics) ObterMetricas() (int, int, int, time.Duration, time.Duration) {
	m.mu.RLock() // Usando leitura sem bloqueio
	defer m.mu.RUnlock()

	return int(atomic.LoadInt32(&m.Allocations)),
		int(atomic.LoadInt32(&m.Reutilizacoes)),
		int(atomic.LoadInt32(&m.BlocosAtivos)),
		time.Duration(atomic.LoadInt64(&m.TempoDeAlocacao)),
		time.Duration(atomic.LoadInt64(&m.TempoDeReutilizacao))
}
