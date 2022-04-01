//Alunos: Igor Dourado e Rafael Dias
package main

import "fmt"
import "sync"


func gera(canal chan int, s int) {
for i := 0; ; i++ { canal <- ((s) * i) }
}

func merge(cs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for v := range c {
				out <- v
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func aux() {
ch1 := make(chan int)
ch2 := make(chan int)
ch3 := make(chan int)

go gera(ch1, 2)
go gera(ch2, 3)
go gera(ch3, 5)


for { fmt.Print(" e " , <- merge(ch1, ch3, ch2) ) } // O "e" é para separar os números
} 
func main() { aux() }