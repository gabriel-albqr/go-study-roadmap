package main // define o programa principal

import "fmt" // importa a biblioteca fmt para formatação de entrada e saída

func main() { // função principal onde a execução do programa começa

	var peso int = 70         // declara uma variável inteira chamada peso e atribui o valor 70
	var altura float64 = 1.75 // declara uma variável float64 chamada altura e atribui o valor 1.75

	fmt.Printf("Peso: %d kg\n", peso)      // imprime o peso formatado
	fmt.Printf("Altura: %.2f m\n", altura) // imprime a altura formatada

	fmt.Println("Hello, World!") // imprime "Hello, World!" no console
}
