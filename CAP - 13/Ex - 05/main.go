/*
- Crie e utilize uma função anônima.
*/

package main

import "fmt"

func main() {
	anonimo()
}

func anonimo() {
	a := 2311

	func(a int) {
		fmt.Println("porcentagem de :", a, "||", "é:", a%10)
	}(a)
}
