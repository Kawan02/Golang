/*
- Crie um package "cachorro".
- Este package deverá exportar uma função Idade, que toma como parâmetro um número de anos e retorna a idade
  equivalente em anos caninos. (1 ano humano → 7 anos caninos)
- Documente seu código com comentários, e utilize a função Idade na sua função main.
- Rode seu programa para verificar se ele funciona.
- Rode um local server com godoc e leia sua documentação.
*/

// Package main retorna a idade.
package main

import (
	"fmt"

	"github.com/kawan02/cachorro"
)

// main pega o valor de idade e coloca na tela.
func main() {
	fmt.Println(cachorro.Idade(3))
}
