package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const N = 50

func main() {
	selectPrime := generateSelection(N)
	p := contaPrimo(selectPrime)
	fmt.Println("\nQuantidades de primos: ", p)
}

func generateSelection(tamanho int) []int {
	selectPrime := make([]int, tamanho, tamanho)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < tamanho; i++ {
		selectPrime[i] = i
	}
	return selectPrime
}

func contaPrimo(s []int) int {
	var primoRun sync.WaitGroup
	primos := make(chan int, N)
	res := 0

	for i := 0; i < N; i++ {
		primoRun.Add(1)
		go isPrime(s[i], i, primos, &primoRun)
	}

	primoRun.Wait()

	for {
		select {
		case i := <-primos:
			fmt.Println(s[i], " Is prime")
			res++
		default:
			return res
		}
	}
}

func isPrime(p int, id int, ans chan int, primoRun *sync.WaitGroup) {
	defer primoRun.Done()

	if p == 0 || p == 1{
		return
	}
	if p%2 == 0 {
		if p > 2 {
			return
		}
	}
	for i := 3; i*i <= p; i += 2 {
		if p%i == 0 {
			return
		}
	}

	//If none of the above, isn't prime
	ans <- id
}
