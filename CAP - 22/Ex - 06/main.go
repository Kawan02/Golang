//- Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	canal := make(chan int)

	go canal1(canal)

	for x := range canal {
		fmt.Println("Valor finall recebido:", x)
	}
}

func canal1(canal chan int) {
	for i := 0; i < 100; i++ {
		canal <- i
		wg.Wait()
	}
	close(canal)
}

/* func EnviaPraCanal(canal chan int, quit chan int) {
	x := 1
	for {
		select {
		case canal <- x:
			x++

		case <-quit:
			return
		}
		wg.Done()
	}
}
*/
