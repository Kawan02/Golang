package main

import "fmt"

// Maneira que eu fiz:
func main() {
	Esporte := "Futebol"
	switch {
	case Esporte == "Futebol":
		fmt.Println("Seu esporte favorito é Futebol!")
	case Esporte == "Basquete":
		fmt.Println("Seu esporte favorito é Basquete!")
	case Esporte == "Surf ":
		fmt.Println("Seu esporte favorito é Surf!")
	}
}

// Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".

//Maneira que a Instrutora fez:
func Instrutora() {
	Esporte := "Futebol"
	switch Esporte {
	case "Futebol":
		fmt.Println("Seu esporte favorito é Futebol!")
	case "Basquete":
		fmt.Println("Seu esporte favorito é Basquete!")
	case "Surf ":
		fmt.Println("Seu esporte favorito é Surf!")
	}
}
