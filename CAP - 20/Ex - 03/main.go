/* - Utilizando goroutines, crie um programa incrementador:
   - Tenha uma variável com o valor da contagem
   - Crie um monte de goroutines, onde cada uma deve:
       - Ler o valor do contador
       - Salvar este valor
       - Fazer yield da thread com runtime.Gosched()
       - Incrementar o valor salvo
       - Copiar o novo valor para a variável inicial
   - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
   - Demonstre que há uma condição de corrida utilizando a flag -race
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

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
			v := contador
			runtime.Gosched() // Gosched: escalonador transfere a execução em cada iteração de loop da primeira goroutine para a segunda e vice-versa.
			v++
			contador = v
			wg.Done()
		}()
	}
}
