/*
- Demonstre o funcionamento de um closure.
- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
*/
package main

import "fmt"

func main() {
	a := closure()
	b := closure()

	fmt.Println(a()*1, "|", b()*4)
	fmt.Println(a()*2, "|", b()*5)
	fmt.Println(a()*3, "|", b()*6)
}

func closure() func() int {
	x := 3
	return func() int {
		return x
	}
}
