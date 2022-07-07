//Utilizando este código: https://play.golang.org/p/MvL6uamrJP
// ...use um select statement para demonstrar os valores do canal.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	q := make(chan int)
	c := gen(q)

	go receive(c, q)
	wg.Wait()

	fmt.Println("---------------------------------\n", "about to exit")

}

func gen(q <-chan int) <-chan int {
	c := make(chan int)
	wg.Add(1)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		wg.Done()
	}()

	return c
}

func receive(x <-chan int, q <-chan int) {
	for {
		select {
		case msg1 := <-x:
			fmt.Println("Valor final do canal C:", msg1)

		case msg2 := <-q:
			fmt.Println("Valor final do canal Q:", msg2)
			//return

		}

	}

}
