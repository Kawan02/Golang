/*
- Atribua uma função a uma variável.
- Chame essa função.
*/
package main

import "fmt"

func main() {
	x := 2022
	a := func() {
		fmt.Println("Estamos no ano de:", x)
	}
	a()
}
