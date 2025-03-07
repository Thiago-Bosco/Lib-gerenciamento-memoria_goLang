# Biblioteca de Gerenciamento de Memória em Go

Essa biblioteca oferece uma implementação eficiente de gerenciamento de memória com suporte a pools de memória concorrentes, métricas detalhadas de uso de memória, e monitoramento de alocações e reutilizações. A ideia é minimizar a contenção entre goroutines e fornecer informações valiosas sobre o uso de memória.

## Funcionalidades

- **Pools de memória concorrentes**: Gerencia múltiplos pools de memória para diferentes tamanhos de blocos de dados, otimizando a alocação e reutilização.
- **Métricas de uso de memória**: Coleta informações sobre alocações, reutilizações, blocos ativos, e tempo gasto em alocações e reutilizações.
- **Monitoramento de uso de memória e tempo de execução**: Monitora o uso de memória e mede o tempo de execução das operações.
- **Persistência de blocos de memória**: Permite persistir blocos de memória alocados em arquivos binários.

## Instalação

Para utilizar a biblioteca, basta incluir a dependência no seu projeto Go:

```bash
go get github.com/Thiago-Bosco/Lib-gerenciamento-memoria_goLang

Exemplo Básico:
package main

import (
	"fmt"
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

	// Reutilizar um bloco
	pool.Put(block1)
	fmt.Printf("Bloco reutilizado (1024 bytes): %d\n", len(block1))

	// Mostrar as métricas
	allocation, reuse, active, allocTime, reuseTime := pool.Metrics.ObterMetricas()
	fmt.Printf("Métricas:\n")
	fmt.Printf(" - Alocações: %d\n", allocation)
	fmt.Printf(" - Reutilizações: %d\n", reuse)
	fmt.Printf(" - Blocos ativos: %d\n", active)
	fmt.Printf(" - Tempo total de alocação: %v\n", allocTime)
	fmt.Printf(" - Tempo total de reutilização: %v\n", reuseTime)
}

===================================================================================================================================================================================
Detalhes das Funções
NewConcurrentMemoryPool(sizes []int) *ConcurrentMemoryPool
Cria um novo pool de memória com diferentes tamanhos de blocos. O parâmetro sizes define os tamanhos de memória a serem gerenciados.

Get(size int) []byte
Retorna um bloco de memória com o tamanho especificado. Se o bloco já estiver alocado, ele é reutilizado; caso contrário, um novo bloco é criado.

Put(block []byte)
Devolve um bloco de memória ao pool para reutilização.

Metrics.ObterMetricas()
Retorna as métricas do uso de memória:

Alocações: O número total de alocações realizadas.
Reutilizações: O número total de blocos de memória reutilizados.
Blocos Ativos: O número atual de blocos de memória alocados no sistema.
Tempo total de alocação: O tempo total gasto nas operações de alocação.
Tempo total de reutilização: O tempo total gasto nas operações de reutilização.
Monitoramento e Persistência
A biblioteca também possui funções para monitorar o uso de memória e tempo de execução, além de persistir blocos de memória em arquivos binários, permitindo uma análise mais profunda do comportamento da aplicação.

Monitoramento de Memória
A função MonitorarUsoDeMemoria permite que você monitore o uso de memória durante a execução. O resultado é mostrado em bytes alocados no sistema.

package profiling

import (
	"fmt"
	"runtime"
)

// Função para medir o uso de memória durante a execução
func MonitorarUsoDeMemoria() {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Printf("Memória utilizada: %v bytes\n", memStats.Alloc)
}

================================================================================================================================================
Persistência de Memória
A função PersistirMemoria permite persistir blocos de memória alocados em um arquivo binário. Isso pode ser útil para análise posterior ou para garantir que dados importantes sejam mantidos.

package persistence

import (
	"encoding/binary"
	"fmt"
	"os"
)

// Função para persistir blocos de memória alocados em um arquivo binário
func PersistirMemoria(filename string, block []byte) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo:", err)
		return
	}
	defer file.Close()

	err = binary.Write(file, binary.LittleEndian, block)
	if err != nil {
		fmt.Println("Erro ao escrever no arquivo:", err)
	}
	fmt.Printf("Bloco de memória de %d bytes persistido em '%s'\n", len(block), filename)
}

===============================================================================================================================

Métricas
As métricas coletadas durante o uso da biblioteca incluem:

Alocações: O número total de alocações realizadas.
Reutilizações: O número total de blocos de memória reutilizados.
Blocos Ativos: O número atual de blocos de memória alocados no sistema.
Tempo total de alocação: O tempo total gasto nas operações de alocação.
Tempo total de reutilização: O tempo total gasto nas operações de reutilização.
Essas métricas são úteis para analisar a eficiência da sua aplicação e identificar possíveis gargalos relacionados ao uso de memória.

Contribuindo
Faça um fork do repositório.
Crie uma nova branch: git checkout -b feature/novo-recurso.
Faça as alterações necessárias e commit: git commit -am 'Adiciona novo recurso'.
Faça o push para a branch: git push origin feature/novo-recurso.
Abra um pull request.
Uso do Go.mod
require github.com/Thiago-Bosco/Lib-gerenciamento-memoria_goLang v1.0.1



teste

