//Alunos: Igor Dourado e Rafael Dias
package main

import "fmt"
import "strconv"
import "time"
import "math/rand"



func printa( ch chan string){
  for{fmt.Print(" "+ <- ch)}
}

func gera(canal chan int) {
rand.Seed(time.Now().UTC().UnixNano())
for i:=0; i<100; i++ { canal <- rand.Intn(9) }
}

func merge(ch1 chan int, ch2 chan int, ch3 chan string) { for{ ch3 <- strconv.Itoa(<-ch1) + strconv.Itoa(<-ch2) } 
}


func main() { 
ch1 := make(chan int)
ch2 := make(chan int)
go gera(ch1)
go gera(ch2)
ch3 := make(chan string)
go merge(ch1, ch2, ch3)
go printa(ch3) 
for { }
}