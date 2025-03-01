package memorypool

import (
	"testing"
)

func TestDynamicMemoryPool(t *testing.T) {
	sizes := []int{512, 1024, 4096} // Diferentes tamanhos de blocos
	pool := NewDynamicMemoryPool(sizes)

	// Teste de alocação e reutilização de blocos
	for _, size := range sizes {
		block := pool.Get(size)
		if len(block) != size {
			t.Errorf("Esperado bloco de %d bytes, mas recebeu %d", size, len(block))
		}

		pool.Put(block) // Devolvendo ao pool

		// Pegando novamente para verificar reutilização
		newBlock := pool.Get(size)
		if len(newBlock) != size {
			t.Errorf("Bloco reutilizado deveria ter %d bytes, mas recebeu %d", size, len(newBlock))
		}
	}
}
