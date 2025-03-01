package memorypool

import (
	"reflect"
	"testing"
)

func TestConcurrentMemoryPool(t *testing.T) {
	// Criação do pool com tamanhos de blocos 128 e 1024 bytes
	pool := NewConcurrentMemoryPool([]int{128, 1024})

	// Teste de alocação
	block := pool.Get(1024)
	if len(block) != 1024 {
		t.Errorf("Esperado bloco de 1024 bytes, mas obteve %d bytes", len(block))
	}

	// Teste de reutilização
	pool.Put(block)

	block2 := pool.Get(1024)

	// Compara o conteúdo dos slices
	if !reflect.DeepEqual(block, block2) {
		t.Errorf("Blocos não são iguais, esperava que fossem idênticos")
	}
}
