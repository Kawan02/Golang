/* - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/
package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {
	ex()
}

func ex() {
	pessoa1 := pessoa{
		nome:      "Zezin",
		sobrenome: "da Silva",
		sabores:   []string{"Flocos", "Chocolate", "Limão"},
	}

	pessoa2 := pessoa{
		nome:      "Carlos",
		sobrenome: "Augusto",
		sabores:   []string{"Ninho", "Napolitano", "Maracuja"},
	}

	fmt.Println("Nome Completo:", pessoa1.nome, pessoa1.sobrenome)
	for _, sabor := range pessoa1.sabores {
		fmt.Println("Sabores favoritos:", sabor)
	}

	fmt.Println("-----------------------------------------")

	fmt.Println("Nome Completo:", pessoa2.nome, pessoa2.sobrenome)
	for _, sabor := range pessoa2.sabores {
		fmt.Println("Sabores favoritos:", sabor)
	}
}
