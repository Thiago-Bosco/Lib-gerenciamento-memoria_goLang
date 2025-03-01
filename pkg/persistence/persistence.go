package persistence

import (
	"fmt"
	"os"
)

// Função para persistir um bloco de memória em um arquivo
func PersistirMemoriaEmArquivo(filename string, data []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("erro ao escrever no arquivo: %w", err)
	}
	return nil
}
