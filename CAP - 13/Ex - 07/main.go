/*
- Atribua uma função a uma variável.
- Chame essa função.
*/
package main

import "fmt"

func main() {
	v := retorna()

	v()
}

func retorna() func() {
	return func() {
		fmt.Println("QualquerCoisa")
	}
}
