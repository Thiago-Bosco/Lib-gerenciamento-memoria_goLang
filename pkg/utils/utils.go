package utils

import "log"

// Função para logar erros
func LogError(err error) {
	if err != nil {
		log.Printf("Erro: %v", err)
	}
}
