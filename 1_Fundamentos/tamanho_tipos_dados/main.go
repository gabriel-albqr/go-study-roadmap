/*
Este código mostra o tamanho em bytes dos principais tipos de dados em Go,
algo importante para entender como a linguagem organiza valores na memória.
Uma das vantagens do Go é justamente sua previsibilidade em relação aos tipos:
inteiros, floats e booleanos têm tamanhos bem definidos, o que facilita escrever
programas mais eficientes e com comportamento consistente em diferentes sistemas.

Além disso, Go diferencia tipos fixos, como int32 e int64, de tipos dependentes
da arquitetura, como int e uint, permitindo que o programador escolha entre
precisão explícita ou desempenho otimizado conforme o ambiente. Esse tipo de
controle é útil para aplicações que exigem alto desempenho, manipulação de dados,
trabalhos com binários ou integração com sistemas de baixo nível.

O código imprime o tamanho de cada tipo usando unsafe.Sizeof, o que ajuda a
visualizar como esses valores são armazenados na prática e reforça como Go lida
com memória de forma simples, direta e sem surpresas.
*/

package main

import (
	"fmt"
	"unsafe"
)

func main() {

	// tamanhos dos tipos inteiros
	fmt.Println("int8:", unsafe.Sizeof(int8(0)), "bytes")
	fmt.Println("int16:", unsafe.Sizeof(int16(0)), "bytes")
	fmt.Println("int32:", unsafe.Sizeof(int32(0)), "bytes")
	fmt.Println("int64:", unsafe.Sizeof(int64(0)), "bytes")

	fmt.Println("uint8:", unsafe.Sizeof(uint8(0)), "bytes")
	fmt.Println("uint16:", unsafe.Sizeof(uint16(0)), "bytes")
	fmt.Println("uint32:", unsafe.Sizeof(uint32(0)), "bytes")
	fmt.Println("uint64:", unsafe.Sizeof(uint64(0)), "bytes")

	// tipos dependentes da arquitetura (32 ou 64 bits)
	fmt.Println("int:", unsafe.Sizeof(int(0)), "bytes")
	fmt.Println("uint:", unsafe.Sizeof(uint(0)), "bytes")

	// tipos float
	fmt.Println("float32:", unsafe.Sizeof(float32(0)), "bytes")
	fmt.Println("float64:", unsafe.Sizeof(float64(0)), "bytes")

	// Boolean
	fmt.Println("bool:", unsafe.Sizeof(true), "bytes")

	// String armazena ponteiro + tamanho
	fmt.Println("string (tamanho estrutural):", unsafe.Sizeof(""), "bytes")

	// Rune (alias para int32)
	fmt.Println("rune:", unsafe.Sizeof(rune(0)), "bytes")

	// Byte (alias para uint8)
	fmt.Println("byte:", unsafe.Sizeof(byte(0)), "bytes")
}
