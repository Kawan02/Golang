/*
Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.s valores e seus índices.
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

	for k, v := range x {
		fmt.Println(k)
		for _, h := range v {
			fmt.Println("\t", h)
		}
		fmt.Println("--------------------------------------")
	}
}
