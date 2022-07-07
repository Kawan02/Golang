package main

import (
	"fmt"
)

func main() {
	ex()
}

func ex() {
	pessoa := [][]string{
		[]string{"Kawan", "Silva", "Messias"},
		[]string{"Messias", "Silva", "Kawan"},
		[]string{"Silva", "Messias", "Kawan"},
	}

	fmt.Println(pessoa, "\n")

	for _, nome := range pessoa {
		fmt.Println("Nome:", nome[0])
		fmt.Println("Sobrenome:", nome[1])
		fmt.Println("Hobby Favorito:", nome[2])
		fmt.Println("-----------------------------------------")

	}
}

/**
Crie uma slice contendo slices de strings ([][]string).
Atribua valores a este slice multi-dimensional da seguinte maneira:
"Nome", "Sobrenome", "Hobby favorito"
Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/
