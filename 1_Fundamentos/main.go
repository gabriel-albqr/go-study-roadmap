package main

import "fmt"

func main() {
	// Tipo inteiro (int)
	var idade int = 21

	// Tipo float (números decimais)
	var altura float64 = 1.65

	// Tipo string (texto)
	var nome string = "Gabriel"

	// Tipo boolean (verdadeiro ou falso)
	var ativo bool = true

	// Declaração curta (Go detecta o tipo automaticamente)
	cidade := "Recife"

	// Exibindo os valores
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Nome:", nome)
	fmt.Println("Ativo:", ativo)
	fmt.Println("Cidade:", cidade)
}
