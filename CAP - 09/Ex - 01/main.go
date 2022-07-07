package main

import "fmt"

func main() {
	ex()
}

//Resolução que eu fiz
func ex() {
	array := [5]int{2, 4, 6, 8, 10}

	for _, x := range array {
		fmt.Println(x, "\n")
	}
	fmt.Printf("%T", array)
}

//Resolução que a Instrutora fez
func inst() {
	array := [5]int{10, 20, 30, 40, 50}
	for i, v := range array {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", array)
}

//Usando uma literal composta:
//Crie um array que suporte 5 valores to tipo int
//Atribua valores aos seus índices
//Utilize range e demonstre os valores do array.
//Utilizando format printing, demonstre o tipo do array.
