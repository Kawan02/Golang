//Faça esse código funcionar: https://play.golang.org/p/j-EA6003P0
//Usando uma função anônima auto-executável
//Usando buffer
package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)

	c <- 50

	fmt.Println("Com buffer:", <-c)

	saida(c)
}

func saida(c chan int) {
	go func() {
		c <- 42
	}()

	fmt.Println("Função anônima:", <-c)
}
