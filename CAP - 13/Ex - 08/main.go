/*
- Callback: passe uma função como argumento a outra função.
*/
package main

import "fmt"

func main() {
	retorna(inicia)
}

func inicia() {
	fmt.Println("Coisa")
}

func retorna(v func()) {
	fmt.Println("Qualquer")
	v()
}
