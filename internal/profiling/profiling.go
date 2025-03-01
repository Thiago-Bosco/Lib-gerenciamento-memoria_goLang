package profiling

import (
	"fmt"
	"runtime"
	"time"
)

// MonitorarUsoDeMemoria exibe o uso de memória e as estatísticas de memória.
func MonitorarUsoDeMemoria() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	fmt.Printf("Memória utilizada (Alloc): %v bytes\n", memStats.Alloc)
	fmt.Printf("Memória total disponível (Sys): %v bytes\n", memStats.Sys)
	fmt.Printf("Memória livre (Free): %v bytes\n", memStats.Frees)
}

// MedirTempoExecucao calcula o tempo decorrido desde o início e exibe.
func MedirTempoExecucao(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("Tempo de execução: %s\n", elapsed)
}
