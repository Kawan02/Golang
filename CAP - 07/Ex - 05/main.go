// You can edit this code!
// Click here and start typing.
package main

import "fmt"

var nasceu = 2003
var atual = 2022

func main() {
	secundaria()
}

func secundaria() {
	for nasceu <= atual {
		fmt.Print(nasceu, "\t")
		nasceu++
	}
}
