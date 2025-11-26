package main

import "fmt"

// constantes globais
const PI = 3.14
const AppName = "Sistema"

func main() {
	// constante local
	const idadeMinima = 18

	// usando as constantes
	fmt.Println("Valor de PI:", PI)
	fmt.Println("Nome do app:", AppName)
	fmt.Println("Idade mínima permitida:", idadeMinima)

	// Constantes não podem ser alteradas
	// idadeMinima = 20  // isso geraria erro
}
