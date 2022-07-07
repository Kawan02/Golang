package main

import "fmt"

func main() {
	CopaDoMundo := 4

	if CopaDoMundo == 5 {
		fmt.Println("Brasil tem 5 Copa do Mundo")
	} else if CopaDoMundo == 3 {
		fmt.Println("Argentina tem 3 Copa do Mundo")
	} else {
		fmt.Println("O Brasil tem 5 Copa do Mundo, mas a Argentina tem 3 Copa do Mundo")
	}
}

// Utilizando a solução anterior, adicione as opções else if e else.
