//Partindo do código abaixo, utilize marshal para transformar []user em JSON.

package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
	Nome  string
	Idade int
	Cpf   int
}

func main() {
	user1 := usuario{"Zezin", 30, 2545156858}
	user2 := usuario{"Carlos", 27, 1546466879}

	usuario := []usuario{user1, user2}
	fmt.Println("Usuarios:", usuario, "\n")

	us, err := json.Marshal(usuario)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Com json:", string(us))

}
