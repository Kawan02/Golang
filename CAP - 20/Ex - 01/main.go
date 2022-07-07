/* Alem da goroutine principal, crie duas outras goroutines.
- Cada goroutine adicional devem fazer um print separado.
- Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
*/

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	NovasGoroutines()
	wg.Wait()
}

func NovasGoroutines() {
	wg.Add(2)
	go func() {
		fmt.Println("Goroutine: 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Goroutine: 2")
		wg.Done()
	}()
}
