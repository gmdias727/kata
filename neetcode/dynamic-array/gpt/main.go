package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}

	// Criando um novo slice com o dobro do tamanho
	newSlice := make([]int, len(slice)*2)

	// Copiando elementos do slice original para o novo slice
	copy(newSlice, slice)

	fmt.Println("Novo slice:", newSlice)
}
