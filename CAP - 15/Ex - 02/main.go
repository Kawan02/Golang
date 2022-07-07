//Crie um struct "pessoa"
//Crie uma função chamada mudeMe que tenha *pessoa como parâmetro.
//Essa função deve mudar um valor armazenado no endereço *pessoa.
package main

import "fmt"

type pessoa struct {
	cpf int
}

func main() {
	cpf := pessoa{24648484844}
	fmt.Println(cpf)
	mudeMe(&cpf)
	fmt.Println(cpf)
}

func mudeMe(p *pessoa) {
	p.cpf = 15448488254
}
