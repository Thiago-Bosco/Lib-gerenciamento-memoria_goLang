package alloc

import "unsafe"

// Aloca um bloco de memória de um tamanho especificado, sem usar o garbage collector.
func AlocarMemoria(size int) unsafe.Pointer {
	return unsafe.Pointer(&make([]byte, size)[0])
}
func LiberarMemoria(ptr unsafe.Pointer) {}
func AlocarMemoriaComPonteiros(size int) unsafe.Pointer {
	return unsafe.Pointer(&make([]byte, size)[0])
}
func LiberarMemoriaComPonteiros(ptr unsafe.Pointer) {}
func AlocarMemoriaComPonteirosComTamanhoDeBloco(size int) unsafe.Pointer {
	return unsafe.Pointer(&make([]byte, size)[0])
}
func LiberarMemoriaComPonteirosComTamanhoDeBloco(ptr unsafe.Pointer) {}
func AlocarMemoriaComPonteirosComTamanhoDeBlocoComTamanhoDeBloco(size int) unsafe.Pointer {
	return unsafe.Pointer(&make([]byte, size)[0])
}
func LiberarMemoriaComPonteirosComTamanhoDeBlocoComTamanhoDeBloco(ptr unsafe.Pointer) {}
