package main

import "fmt"

func main() {

	// operadores aritméticos
	a := 10
	b := 3

	soma := a + b          // adição
	subtracao := a - b     // subtração
	multiplicacao := a * b // multiplicação
	divisao := a / b       // divisão inteira
	resto := a % b         // resto da divisão

	// operadores de comparação
	igual := a == b        // igual
	diferente := a != b    // diferente
	maior := a > b         // maior que
	menor := a < b         // menor que
	maiorOuIgual := a >= b // maior ou igual
	menorOuIgual := a <= b // menor ou igual

	// operadores lógicos
	x := true
	y := false

	and := x && y // AND lógico
	or := x || y  // OR lógico
	not := !x     // NOT lógico

	fmt.Println("Soma:", soma)
	fmt.Println("Subtração:", subtracao)
	fmt.Println("Multiplicação:", multiplicacao)
	fmt.Println("Divisão:", divisao)
	fmt.Println("Resto:", resto)

	fmt.Println("Igual:", igual)
	fmt.Println("Diferente:", diferente)
	fmt.Println("Maior:", maior)
	fmt.Println("Menor:", menor)
	fmt.Println("Maior ou igual:", maiorOuIgual)
	fmt.Println("Menor ou igual:", menorOuIgual)

	fmt.Println("AND:", and)
	fmt.Println("OR:", or)
	fmt.Println("NOT:", not)
}
