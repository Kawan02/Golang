//Partindo do código abaixo, utilize unmarshal e demonstre os valores.
package main

import (
	"encoding/json"
	"fmt"
)

type Pessoa []struct {
	First   string   `json: "Nome"`
	Last    string   `json: "Sobrenome"`
	Age     int      `json: "Idade"`
	Sayings []string `json: "Provérbios"`
}

func main() {
	s := (`[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`)
	fmt.Println("JSON:", s, "\n")

	var pessoa Pessoa
	err := json.Unmarshal([]byte(s), &pessoa)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("GO:", pessoa)

}
