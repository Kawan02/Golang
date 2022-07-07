/*
   - Exercício:
   - Crie uma função que retorne um int
   - Crie outra função que retorne um int e uma string
   - Chame as duas funções
   - Demonstre seus resultados
*/

package main

import "fmt"

func main() {
	fmt.Println(um())
	fmt.Println(dois())
}

func um() int {
	return 2003

}

func dois() (int, string) {
	return 2022, "Ano de Copa do Mundo"
}
