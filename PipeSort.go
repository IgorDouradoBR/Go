// Igor Dourado e Rafael Dias PIPESORT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 20 //20 valores a serem impressos

func pipeSortGo(canalCorrente chan int, canalSeguinte chan int, canalAux chan int) {
	var booleano bool = false
	var corrente int
	for {
		var aux int = <-canalCorrente
		if aux == 81 { // valor fixo mais a sobra
			canalAux <- corrente
			canalSeguinte <- aux
			break
		}
		if booleano == false {
			booleano = true
			corrente = aux
		} else if aux < corrente {
			canalSeguinte <- corrente
			corrente = aux
		} else {
			canalSeguinte <- aux
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var arrayDeCanal [N + 1]chan int
	var canalDoSort chan int = make(chan int)
	for i := 0; i <= N; i++ {
		arrayDeCanal[i] = make(chan int)
	}
	for i := 0; i < N; i++ {
		go pipeSortGo(arrayDeCanal[i], arrayDeCanal[i+1], canalDoSort)
	}
	for i := 0; i < N; i++ {
		var numero int = rand.Intn(80) // 80 é o valor max dos valores gerados
		arrayDeCanal[0] <- numero
	}
	arrayDeCanal[0] <- 80 + 1 //80 é o valor máximo a ser alcançado mais uma sobra
	fmt.Print("ORDENADO COM PIPESORT:  ")
	for i := 0; i < N; i++ {
		fmt.Print(<-canalDoSort, " ")
	}
	<-arrayDeCanal[N] //liberando
}
