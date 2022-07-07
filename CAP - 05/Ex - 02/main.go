package main

import "fmt"

var A = (50 == 500)
var B = (50 != 500)
var C = (50 <= 500)
var D = (50 < 500)
var E = (50 >= 500)
var F = (50 > 500)

func main() {

	fmt.Printf("%v\n%v\n%v\n%v\n%v\n%v\n", A, B, C, D, E, F)

}
