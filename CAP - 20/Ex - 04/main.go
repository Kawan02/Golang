//- Utilize mutex para consertar a condição de corrida do exercício anterior.

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mt sync.Mutex
var wg sync.WaitGroup

const quantidade = 100

var contador int

func main() {
	//contador := 0

	demonstre(quantidade)

	wg.Wait()
	fmt.Println("Goroutines:", quantidade)
	fmt.Println("Contador:", contador)
}

func demonstre(g int) {
	wg.Add(g)
	for j := 0; j < g; j++ {
		go func() {
			mt.Lock() //Trancar
			v := contador
			runtime.Gosched() // Gosched: escalonador transfere a execução em cada iteração de loop da primeira goroutine para a segunda e vice-versa.
			v++
			contador = v
			wg.Done()
			mt.Unlock() //Desbloquear
		}()
	}
}
