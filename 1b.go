//Alunos: Igor Dourado e Rafael Dias
package main

import "fmt"
import "time"
import "math/rand"



func printa( ch chan int){
  for{
    fmt.Print( <- ch)
    fmt.Print( <- ch)
    fmt.Print( <- ch)// "1 valor de cada canal"(em teoria) juntos e separados por espaço
    fmt.Print(" ")}
}

func gera(canal chan int) {
rand.Seed(time.Now().UTC().UnixNano())
for i:=0; i<100; i++ { canal <- rand.Intn(9) }
}

func merge(ch1 chan int, ch2 chan int, ch3 chan int) { 

for{ select{
    case dado := <- ch1:
      ch3 <- dado
    case dado := <- ch2:
      ch3 <- dado

} } 
}


func main() { 
ch1 := make(chan int)
ch2 := make(chan int)
go gera(ch1)
go gera(ch2)
ch3 := make(chan int)
go merge(ch1, ch2, ch3)//merge da parte 1

ch4 := make(chan int)
go gera(ch4)
ch5 := make(chan int)
go merge(ch3, ch4, ch5)//merge da parte 2 com o outro gerador
go printa(ch5)//consumidor final
for { }
}