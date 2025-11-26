package main

import (
	"fmt"
	"strconv"
)

func main() {

	// exemplos de conversões em Go (com tratamento quando necessário).
	// Go exige conversões explícitas entre tipos numéricos e usa a biblioteca
	// strconv para conversões entre strings e números.

	// conversões numéricas explícitas
	var a int = 42
	var b float64 = float64(a) // int -> float64
	var c int = int(b)         // float64 -> int (perde parte fracionária)
	fmt.Println("int a:", a)
	fmt.Println("float64 b (from a):", b)
	fmt.Println("int c (from b):", c)

	// int para string
	s1 := strconv.Itoa(a) // converte int para string
	fmt.Println("string s1 (Itoa):", s1)

	// string para int (pode falhar -> tratar erro)
	s2 := "1234"
	n, err := strconv.Atoi(s2) // converte string para int
	if err != nil {
		fmt.Println("Erro ao converter s2 para int:", err)
	} else {
		fmt.Println("int n (Atoi):", n)
	}

	// string para float64 (parse com precisão)
	s3 := "3.1415"
	f, err := strconv.ParseFloat(s3, 64) // 64 bits de precisão
	if err != nil {
		fmt.Println("Erro ao converter s3 para float:", err)
	} else {
		fmt.Println("float64 f (ParseFloat):", f)
	}

	// float para string
	fs := strconv.FormatFloat(f, 'f', 4, 64) // formato 'f', 4 casas, 64 bits
	fmt.Println("string fs (FormatFloat):", fs)

	// bytes <-> string
	str := "Olá, Go!"
	bytes := []byte(str)          // string -> []byte
	strFromBytes := string(bytes) // []byte -> string
	fmt.Println("bytes:", bytes)
	fmt.Println("string from bytes:", strFromBytes)

	// rune (caractere Unicode) para int e string
	r := '世'                  // rune literal (tipo rune = int32)
	fmt.Println("rune r:", r) // valor numérico (code point)
	fmt.Println("rune para string:", string(r))

	// conversão segura com base em prefixos (ex.: hexadecimal)
	hexStr := "1a"                               // representa 26 em hex
	val, err := strconv.ParseInt(hexStr, 16, 64) // base 16, até 64 bits
	if err != nil {
		fmt.Println("Erro ao converter hexStr:", err)
	} else {
		fmt.Println("hexStr em decimal (ParseInt base16):", val)
	}
}
