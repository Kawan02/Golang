/* Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
- Demonstre os valores do map utilizando range.
- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
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

	m := make(map[string]pessoa)

	m["Zezin"] = pessoa{
		nome:      "Zezin",
		sobrenome: "da Silva",
		sabores:   []string{"Flocos", "Chocolate", "Limão"},
	}

	m["Carlos"] = pessoa{
		nome:      "Carlos",
		sobrenome: "Augusto",
		sabores:   []string{"Ninho", "Napolitano", "Maracuja"},
	}

	for _, m := range m {
		fmt.Println("Nome Completo:", m.nome, m.sobrenome)

		for _, v := range m.sabores {

			fmt.Println("Sabores favoritos:", v)

		}

		fmt.Println("-----------------------------------------")
	}
}

/*
	fmt.Println("Nome Completo:", pessoa2.nome, pessoa2.sobrenome)
	for _, sabor := range pessoa2.sabores {
		fmt.Println("Sabores favoritos:", sabor) */
