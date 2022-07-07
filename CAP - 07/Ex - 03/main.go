package main

import (
	"fmt"
)

// Maneira que ela fez
func main() {
	nasceu := 2003
	atual := 2022

	for nasceu <= atual {
		fmt.Print(nasceu, "\t")
		nasceu++
	}

}

// Maneira que eu fiz
func ex() {
	for a := 2003; a <= 2022; a++ {
		fmt.Print(a, "\t")
	}

}

//Crie um loop utilizando a sintaxe: for condition {}
//Utilize-o para demonstrar os anos desde que você nasceu.
