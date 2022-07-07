/* Crie um novo tipo: veículo
    - O tipo subjacente deve ser struct
    - Deve conter os campos: portas, cor
- Crie dois novos tipos: caminhonete e sedan
    - Os tipos subjacentes devem ser struct
    - Ambos devem conter "veículo" como struct embutido
    - O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
    - O tipo sedan deve conter um campo bool chamado "modeloLuxo"
- Usando os structs veículo, caminhonete e sedan:
    - Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
    - Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
- Demonstre estes valores.
- Demonstre um único campo de cada um dos dois.
*/
package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	traçãoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	ex()
}

func ex() {

	x := caminhonete{veiculo: veiculo{2, "verde"}, traçãoNasQuatro: false}

	y := sedan{veiculo: veiculo{4, "prata"}, modeloLuxo: true}

	fmt.Println("Especificações do seu veiculo:", x)
	fmt.Println("--------------------------------------------------")
	fmt.Println("Especificações do seu veiculo:", y)
	fmt.Println("--------------------------------------------------")
	fmt.Println("A cor do seu veiculo é:", x.cor, " . E seu veiculo tem:", x.portas, "portas")
	fmt.Println("--------------------------------------------------")
	fmt.Println("A cor do seu veiculo é:", y.cor, ". E seu veiculo tem:", y.portas, "portas")
}

/*
	fmt.Println("Nome Completo:", pessoa2.nome, pessoa2.sobrenome)
	for _, sabor := range pessoa2.sabores {
		fmt.Println("Sabores favoritos:", sabor) */
