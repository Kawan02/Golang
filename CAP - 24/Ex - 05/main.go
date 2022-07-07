//- Nos capítulos seguintes, uma das coisas que veremos é testes.
//- Para testar sua habilidade de se virar por conta própria... desafio:
//    - Utilizando as seguintes fontes: https://godoc.org/testing & https://www.golang-book.com/books/intro/12
//    - Tente descobrir por conta própria como funcionam testes em Go.
//    - Pode usar tradutor automático, pode rodar código na sua máquina, pode procurar no Google. Vale tudo.
//   - O exercício é: crie um teste simples de uma função ou método ou pedaço qualquer de código.

package main

//package math

import (
	"fmt"
	"testing"
)

func main() {
	Mostra(0)

}
func Mostra(x int) int {
	total := x + 10
	fmt.Println(total)
	return x
}

func TestMostra(t *testing.T) {
	v := Mostra(5)
	if v != 1 {
		t.Errorf("Mostrar(1) -> valor inserido %d", v)
	}
}
