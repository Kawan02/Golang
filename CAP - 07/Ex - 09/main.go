package main

import "fmt"

func main() {
	cidade := 2
	switch {
	case cidade == 0:
		fmt.Println("Você é Carioca!")
	case cidade == 1:
		fmt.Println("Você é Mineiro!")
	case cidade == 2:
		fmt.Println("Você é Paulista!")
	}
}

//Crie um programa que utilize a declaração switch, sem switch statement especificado
