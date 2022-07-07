package main

import "fmt"

var a = 24

func main() {
	fmt.Printf("%d, %b, %#x", a, a, a)
	fmt.Printf("\n")
	x := a << 1
	fmt.Printf("%d, %b, %#x", x, x, x)
}

// Atribua um valor int a uma variável
// Demonstre este valor em decimal, binário e hexadecimal
// Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
// Demonstre esta outra variável em decimal, binário e hexadecimal
