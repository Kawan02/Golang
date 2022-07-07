/*
- Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
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
	x["Messias_Geraldo"] = []string{"AssistirTV", "Viajar"}
	fmt.Println(x, "\n")

	delete(x, "Messias_Geraldo")

	for k, v := range x {
		fmt.Println(k)
		for _, h := range v {
			fmt.Println("\t", h)
		}
		fmt.Println("--------------------------------------")
	}
}
