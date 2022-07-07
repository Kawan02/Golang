package main

import (
	"fmt"
)

var A = 10
var B = 9

func main() {
	for a := 65; a <= 90; a++ {
		fmt.Println(a)
		for b := 1; b <= 3; b++ {
			fmt.Printf("%#U\n", a)
		}
	}

}

// Coloque na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
