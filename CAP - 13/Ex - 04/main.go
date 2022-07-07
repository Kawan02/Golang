/*
Crie um tipo struct "pessoa" que contenha os campos:
nome, sobrenome, idade
Crie um método para "pessoa" que demonstre o nome completo e a idade;
Crie um valor de tipo "pessoa"
*/

package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func main() {
	sec()
}

//Método para pessoa
func (p pessoa) d() {
	fmt.Println("Nome completo:", p.nome, p.sobrenome)
	fmt.Println("Idade:", p.idade)
}

// Valor de tipo "pessoa"
func sec() {
	pessoa1 := pessoa{"Kawan", "Silva Messias", 18}
	pessoa1.d()
}

/* Outra maneira de fazer
func sec() {
	pessoa1 := pessoa{
	mome: "Kawan",
	sobrenome: "Silva Messias",
	idade: 18,
	}
	pessoa1.d()
}*/
