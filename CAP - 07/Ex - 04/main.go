package main

import "fmt"

var nasceu = 2003
var atual = 2022

func main() {
	for {
		if nasceu > atual {
			break
		}
		fmt.Print(nasceu, "\t")
		nasceu++
	}
}

func secundaria() {
	for nasceu <= atual {
		fmt.Print(nasceu, "\t")
		nasceu++
	}
}

//Crie um loop utilizando a sintaxe: for {}
// Utilize-o para demonstrar os anos desde que você nasceu.
