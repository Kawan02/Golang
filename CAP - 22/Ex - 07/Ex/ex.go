//- Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
// - Tire estes números do canal e demonstre-os.

package main

import "fmt"

var canal = make(chan int)

func main() {

	demonstre()

}

func demonstre() {
	for i := 0; i < 10; i++ {
		go func() {
			for x := 1; x < 11; x++ {
				canal <- x
			}
		}()
	}

	for k := 1; k < 101; k++ {
		fmt.Println(k, "\t", "---------------", "\t", <-canal)
	}

}
