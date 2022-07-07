/*
- Crie e use um struct anônimo.
- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/
package main

import "fmt"

func main() {
	ex()
}

func ex() {
	x := struct {
		nome  map[string]string
		valor []int
	}{
		map[string]string{"QualquerCoisa": "Kawan"}, []int{10, 20, 30, 40, 50},
	}
	fmt.Println(x)
}
