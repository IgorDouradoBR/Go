package main

import (
	"math/rand"
	"time"
)

func main () {
	slice := generateSlice(20)
	mergeSortGoPar(slice)
}

func generateSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func mergeSortGoPar(s []int) []int {
	if len(s) > 1 {
		middle := len(s) / 2

		var s1 []int
		var s2 []int
		c := make(chan struct{}, 2)

		go func() {
			s1 = mergeSortGoPar(s[middle:])
			c <- struct{}{}
		}()
		go func() {
			s2 = mergeSortGoPar(s[:middle])
			c <- struct{}{}
		}()
		<-c
		<-c
		return merge(s1,s2)
	}
	return s
}

