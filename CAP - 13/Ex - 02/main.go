/*
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
*/

package main

import "fmt"

func main() {
	x := []int{2003, 2004, 2005, 2006, 2007}
	y := []int{2022, 2023, 2024, 2025, 2026}

	fmt.Println(soma1(x...))
	fmt.Println(soma2(y))

}

func soma1(x ...int) int {
	soma := 03
	for _, valor := range x {
		soma += valor
	}
	return soma
}

func soma2(x []int) int {
	s := 22
	for _, valor := range x {
		s += valor
	}
	return s
}
