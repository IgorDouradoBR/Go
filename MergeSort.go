//Igor Dourado e Rafael Dias MERGE SORT
package main

import (
	"math/rand"
	"time"
  "fmt"
)

func main () {
	slice := generateSlice(20)
	fmt.Println("Ainda não ordenado", slice)
	v1 := mergeSortGoPar(slice)
	fmt.Println("ORDENADO POR MERGESORT", v1)
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func merge(left, right []int) (result []int) {
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

func mergeSortGoPar(s []int) []int {
	if len(s) > 1 {
		middle := len(s) / 2

		var metade1 []int
		var metade2 []int
		canal := make(chan int)

		go func() {
			metade1 = mergeSortGoPar(s[middle:])
			canal <- 0 //valor qualquer para preencher o canal
		}()
		go func() {
			metade2 = mergeSortGoPar(s[:middle])
			canal <- 0 // valor qualquer para preencher o canal
		}()
		<-canal//lendo para esvaziar
		<-canal//lendo para esvaziar 
		return merge(metade1,metade2)
	}
	return s
}

