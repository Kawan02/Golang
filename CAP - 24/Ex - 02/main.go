//- Utilizando este código: https://play.golang.org/p/9a1IAWy5E6
//- ...crie uma mensagem de erro customizada utilizando fmt.Errorf().

package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Println("Ocorreu um erro")
	}

	fmt.Println(string(bs))

}

// toJSON needs to return an error also ----> toJSON precisa retornar um erro também
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{},
			fmt.Errorf("Ocorreu um erro inesperado: %v", err) //%v ele retorna um struct --->https://www.geeksforgeeks.org/fmt-errorf-function-in-golang-with-examples/
	}
	return bs, nil
}
