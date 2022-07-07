//Utilizando este código: https://play.golang.org/p/sfyu4Is3FG
//...use um for range loop para demonstrar os valores do canal.
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := gen()
	go receive(c)
	wg.Wait()

	fmt.Println("about to exit")
}

func gen() <-chan int {
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

//Funcão receive criada e implementada o for range loop
func receive(r <-chan int) {
	for x := range r {
		fmt.Println(x)
	}
	wg.Wait()
	return

	//close(r)
}
