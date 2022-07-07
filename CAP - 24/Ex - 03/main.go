//- Crie um struct "erroEspecial" que implemente a interface builtin.error.
//- Crie uma função que tenha um valor do tipo error como parâmetro.
//- Crie um valor do tipo "erroEspecial" e passe-o para a função da instrução anterior.

package main

import (
	"log"
)

type erroEspecial struct {
	erro string
}

func main() {
	x := erroEspecial{"Evento inesperado"}
	retorna(x)
}

func (e erroEspecial) Error() string {
	return "Error: 102"
}

func retorna(e error) {
	log.Println("Ocorreu um problema inesperado -", e)
}
