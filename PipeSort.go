// PIPE SORT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const N = 10
const MAX = 100

func main() {

	var resultado chan int = make(chan int)
	var canais [N + 1]chan int
	for i := 0; i <= N; i++ {
		canais[i] = make(chan int)
	}
	for i := 0; i < N; i++ {
		go ordenaCelula(i, canais[i], canais[i+1], resultado, MAX)
	}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < N; i++ {
		valor := rand.Intn(MAX) - rand.Intn(MAX)
		canais[0] <- valor
		fmt.Println(i, " : ", valor)
	}
	canais[0] <- MAX + 1 

	for i := 0; i < N; i++ {
		fmt.Println(i, " : ", <-resultado)
	}
	<-canais[N]
}

func ordenaCelula(i int, entrada chan int, saida chan int, resultado chan int, max int) {
	var valorEscolhido int
	var undefined bool = true
	for {
		n := <-entrada
		if n == max+1 { 
			resultado <- valorEscolhido 
			saida <- n   
			break   
		}
		if undefined {
			valorEscolhido = n
			undefined = false
		} else if n >= valorEscolhido {
			saida <- n
		} else {
			saida <- valorEscolhido
			valorEscolhido = n
		}
	}
}