//- Utilize atomic para consertar a condição de corrida do exercício #3.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

const quantidade = 10

//var x int64 = quantidade

var contador int64

func main() {
	//contador = 0
	//fmt.Println("CPUs:", runtime.NumCPU())

	demonstre(quantidade)
	wg.Wait()

	fmt.Println("Goroutines:", quantidade)
	fmt.Println("Contador:", contador)
}

func demonstre(g int) {
	wg.Add(g)
	for j := 0; j < g; j++ {
		go func() {
			atomic.AddInt64(&contador, 1)
			v := atomic.LoadInt64(&contador)
			fmt.Println(v)
			wg.Done()
		}()
	}

	//fmt.Println("Contador:", atomic.LoadInt64(&contador))
	//fmt.Println("Quantidade:", atomic.LoadInt64(&x))
}
