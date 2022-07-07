//Partindo do código abaixo, ordene os []user por idade e sobrenome.
//Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenarPorIdade []user

func (x ordenarPorIdade) Len() int           { return len(x) }
func (x ordenarPorIdade) Less(i, j int) bool { return x[i].Age < x[j].Age }
func (x ordenarPorIdade) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// --------------------------------------------------------------------------------------------------------
type ordenarPorSobrenome []user

func (x ordenarPorSobrenome) Len() int           { return len(x) }
func (x ordenarPorSobrenome) Less(i, j int) bool { return x[i].Last < x[j].Last }
func (x ordenarPorSobrenome) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// --------------------------------------------------------------------------------------------------------

func main() {
	atividade()
}

func atividade() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	//fmt.Println(users, "\n")

	// Sayings
	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	sort.Sort(ordenarPorIdade(users))
	fmt.Println("Por Idade:\n", users, "\n")
	fmt.Println("Por Idade:\n")

	harmoniosa(users)

	sort.Sort(ordenarPorSobrenome(users))
	fmt.Println("Por Sobrenome:\n", users, "\n")
	fmt.Println("Por Sobrenome:\n")

	harmoniosa(users)
}

// Instr... &&
func harmoniosa(x []user) {
	for i, v := range x {
		fmt.Println(i, "\tIdade:", v.Age, "\tNome completo:", v.First, v.Last, "\n")
		for _, c := range v.Sayings {
			fmt.Println("\t", c)
		}
		fmt.Println("\n")

	}

}
