/*
- Callback: passe uma função como argumento a outra função.
*/
package main

import "fmt"

func main() {
	retorno := retorna()

	retorno()
}

func retorna() func() {
	return func() {
		fmt.Println("Qualquer Coisa")
	}
}
