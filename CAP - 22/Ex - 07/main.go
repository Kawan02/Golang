//- Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
// - Tire estes números do canal e demonstre-os.

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var canal = make(chan int)

func main() {

	lancar()

}

func lancar() {
	//c := make(chan int)
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

//func receive(c <-chan int, canal <-chan int) {
//	for {
//		select {
//		case msg1 := <-c:
//			fmt.Println("Valor final do canal A:", msg1)
//			fmt.Println("")
//			time.Sleep(100000)
//		case msg1 := <-c:
//		fmt.Println("Valor final do canal B:", msg1)
//			fmt.Println("")
//			time.Sleep(100000)
//
//			/* case msg2 := <-q:
//			fmt.Println("Valor final do canal Q:", msg2)
//			//return
//			*/
//		}
//
//	}
//}
