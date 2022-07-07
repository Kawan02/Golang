package main

import "fmt"

func main() {
	fmt.Println(true && true)  // && = Operador END. Para as coisas serem TRUE(Verdadeiro(a)) todas tem que ser TRUE.
	fmt.Println(true && false) // && = Operador END. Para as coisas serem Verdadeira(TRUE) todas tem que ser TRUE, como tem um FALSE, ele vai retornar FALSE.
	fmt.Println(true || true)  // || = OU - OR. Ele retornará TRUE se tiver qualquer TRUE no meio, ele será Verdadeiro.
	fmt.Println(true || false) // || = OU - OR. Ele retornará TRUE se tiver qualquer TRUE no meio, ele será Verdadeiro.
	fmt.Println(!true)         // ! = NOT - NÃO. Basicamente tudo que você quer fazer, com esse operador ele retornará ao contrario, ou seja, Verdaderio será falso
}

// Anote (à mão) o resultado das expressões:
