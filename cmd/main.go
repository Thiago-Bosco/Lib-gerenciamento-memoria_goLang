package main

import (
	"fmt"
	"time"

	"github.com/Thiago-Bosco/Lib-gerenciamento-memoria_goLang/pkg/memorypool"
	"github.com/Thiago-Bosco/Lib-gerenciamento-memoria_goLang/pkg/persistence"
	"github.com/Thiago-Bosco/Lib-gerenciamento-memoria_goLang/pkg/profiling"
)

func main() {
	// Tamanhos de blocos para o pool de memória
	sizes := []int{512, 1024, 4096}

	// Criar o pool de memória concorrente
	pool := memorypool.NewConcurrentMemoryPool(sizes)

	// Alocar alguns blocos de memória
	block1 := pool.Get(1024)
	block2 := pool.Get(512)
	block3 := pool.Get(4096)

	// Exibir as alocações
	fmt.Println("Blocos alocados:")
	fmt.Printf(" - 512 bytes: %d\n", len(block2))
	fmt.Printf(" - 1024 bytes: %d\n", len(block1))
	fmt.Printf(" - 4096 bytes: %d\n", len(block3))

	// Monitoramento de tempo de execução
	start := time.Now()

	// Reutilizar um bloco
	pool.Put(block1)
	fmt.Printf("Bloco reutilizado (1024 bytes): %d\n", len(block1))

	// Realizando mais alocações para medir o tempo de alocação
	for i := 0; i < 1000; i++ {
		_ = pool.Get(1024)
	}

	// Monitorar o tempo de execução
	profiling.MedirTempoExecucao(start)

	// Monitorar o uso de memória durante o processo
	profiling.MonitorarUsoDeMemoria()

	// Persistir o bloco de memória alocado para o arquivo (exemplo)
	err := persistence.PersistirMemoriaEmArquivo("memoria_alocada.bin", block2)
	if err != nil {
		fmt.Printf("Erro ao persistir a memória: %v\n", err)
	} else {
		fmt.Println("Bloco de memória de 512 bytes persistido em 'memoria_alocada.bin'")
	}

	// Mostrar as métricas
	allocation, reuse, active, allocTime, reuseTime := pool.Metrics.ObterMetricas() // Acessando 'Metrics' corretamente
	fmt.Printf("Métricas:\n")
	fmt.Printf(" - Alocações: %d\n", allocation)
	fmt.Printf(" - Reutilizações: %d\n", reuse)
	fmt.Printf(" - Blocos ativos: %d\n", active)
	fmt.Printf(" - Tempo total de alocação: %v\n", allocTime)
	fmt.Printf(" - Tempo total de reutilização: %v\n", reuseTime)
}
