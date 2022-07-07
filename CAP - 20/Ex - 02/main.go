/* - Crie um tipo para um struct chamado "pessoa" OK
   - Crie um metodo "falar" para este tipo que tenha um receiver ponteiro (*pessoa)
   - Crie uma interface, "humanos", que seja implementada por tipos com o método "falar"
   - Crie uma funcao "dizerAlgumaCoisa" cujo parâmetro seja do tipo "humanos" e que chame o método "falar"
   - Demonstre no seu codigo:
       - Que você pode utilizar um valor do tipo "*pessoa" na função "dizerAlgumaCoisa"
       - Que você nao pode utilizar um valor do tipo "pessoa" na função "dizerAlgumaCoisa"
*/
package main

import (
	"fmt"
)

type Pessoa struct {
	nome      string
	sobrenome string
	cpf       int
	nasceu    int
}

type humanos interface {
	Falar()
}

func main() {
	roda()
}

func (m *Pessoa) Falar() {
	fmt.Println("Nome:", m.nome, "\t", "Sobrenome:", m.sobrenome, "\t", "CPF:", m.cpf,
		"\t", "Ano de nascimento:", m.nasceu)
}

func dizerAlgumaCoisa(h humanos) {
	h.Falar()
}

func roda() {
	diga := Pessoa{"Kawan", "Messias", 73178318731, 2003}

	diga.Falar() // <- É um shortcut para (&diga).Falar()
	// (&diga).Falar() // É a maneira "mais correta"

	fmt.Println("")

	dizerAlgumaCoisa(&diga)

	//fmt.Println(&diga)
}
