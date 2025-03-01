package memorypool

import (
	"testing"
)

func TestMemoryPool(t *testing.T) {
	pool := NewMemoryPool(1024) // Criando um pool de blocos de 1KB

	block := pool.Get() // Obtendo um bloco
	if len(block) != 1024 {
		t.Errorf("Esperado bloco de 1024 bytes, mas recebeu %d", len(block))
	}

	pool.Put(block) // Devolvendo o bloco ao pool

	// Pega um novo bloco e verifica se a reutilização está funcionando
	newBlock := pool.Get()
	if len(newBlock) != 1024 {
		t.Errorf("Bloco reutilizado deveria ter 1024 bytes, mas recebeu %d", len(newBlock))
	}
}
