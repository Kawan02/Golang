/*- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
- Demonstre todos esses valores e seus índices.
*/

package main

import "fmt"

func main() {
	ex()
}

func ex() {
	x := map[string][]string{
		"Messias_Kawan": {"JogarBola", "Estudar"},
	}
	for i, y := range x {
		fmt.Println(i)
		for _, a := range y {
			fmt.Println("\t", a)
		}
	}
}
