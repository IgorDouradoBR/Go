// Merge Sort in Golang
// extratos baixados da internet
// modificacoes feitas por Fernando Dotti - PUCRS
//
// aqui encontram-se 2 implementacoes de mergeSort
//   sequenciais
// o programa avalia o tempo de execucao de cada uma
// go run MergeSortSeqSlice.go

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("------ DIFFERENT MERGESORT IMPLEMENTATIONS -------")

	slice := generateSlice(20)

	fmt.Println("-Unsorted - ", slice)

	start := time.Now()
	v1 := mergeSort(slice)
	fmt.Println("  -> traditional ------ secs: ", time.Since(start).Seconds())
	fmt.Println("--- Sorted -----------------------", v1)

	start1 := time.Now()
	v2 := mergeSortGo(slice)
	fmt.Println("  -> mergeSortGo ------ secs: ", time.Since(start1).Seconds())
	fmt.Println("--- Sorted with mergeSortGo ------", v2)
}

// Generates a slice of size, size filled with random numbers
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

// ---------------------------------------------------------------------
// mergeSortGo: usa facilidades de slices (so isso!)
// SE VOCE NAO ENTENDE SLICES, VEJA NA PARTE DE GO BASICO O CONTEUDO E REFERENCIAS

func mergeSortGo(s []int) []int {
	if len(s) > 1 {
		middle := len(s) / 2
		return merge(mergeSortGo(s[:middle]), mergeSortGo(s[middle:]))
	}
	return s
}

// ---------------------------------------------------------------------
// mergeSort: uma implementacao tradicional

func mergeSort(items []int) []int {
	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
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
