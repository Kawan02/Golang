package main

import "fmt"

func main() {
	ex()
}

func ex() {
	estados := make([]string, 26, 27)
	estados = []string{
		"Acre (AC)", "Alagoas (AL)", "Amapá (AP)",
		"Amazonas (AM)", "Bahia (BA)", "Ceará (CE)",
		"Espírito Santo (ES)", "Goiás (GO)", "Maranhão (MA)",
		"Mato Grosso (MT)", "Mato Grosso do Sul (MS)", "Minas Gerais (MG)",
		"Pará (PA)", "Paraíba (PB)", "Paraná (PR)",
		"Pernambuco (PE)", "Piauí (PI)", "Rio de Janeiro (RJ)",
		"Rio Grande do Norte (RN)", "Rio Grande do Sul (RS)", "Rondônia (RO)",
		"Roraima (RR)", "Santa Catarina (SC)", "São Paulo (SP)", "Sergipe (SE)", "Tocantins (TO)",
	}
	for x := 0; x < len(estados); x++ {
		fmt.Println("Estado:", estados[x])
	}
}

/**
Crie uma slice usando make que possa conter todos os estados do Brasil.
Os estados: "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás", "Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco", "Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina", "São Paulo", "Sergipe", "Tocantins"
Demonstre o len e cap da slice.
Demonstre todos os valores da slice *sem utilizar range.*
*/
