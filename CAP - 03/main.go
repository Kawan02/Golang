// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type hotdog int

var x hotdog = 42
var y int

func main() {
	// var v int
	// v = 42
	// x := 40
	fmt.Printf("O valor de X: %v do tipo %T\n", x, x)
	y = int(x)
	fmt.Printf("O valor de Y: %v do tipo %T\n", y, y)
}

//  1. Demonstre o valor da variável "x"
//    2. Demonstre o tipo da variável "x"
//    3. Atribua 42 à variável "x" utilizando o operador "="
//    4. Demonstre o valor da variável "x"
