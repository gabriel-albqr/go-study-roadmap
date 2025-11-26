package main

import "fmt"

func main() {
	// tipo inteiro (int)
	var idade int = 21

	// tipo float (números decimais)
	var altura float64 = 1.65

	// tipo string (texto)
	var nome string = "Gabriel"

	// tipo boolean (verdadeiro ou falso)
	var ativo bool = true

	// declaração curta (Go detecta o tipo automaticamente)
	cidade := "Recife"

	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Nome:", nome)
	fmt.Println("Ativo:", ativo)
	fmt.Println("Cidade:", cidade)
}
