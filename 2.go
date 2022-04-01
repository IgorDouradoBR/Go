//Alunos: Igor Dourado e Rafael Dias
package main

import "fmt"


func gera(canal chan int, s int) {
for i := 0; i<101; i++ { 
  if i==0 { } else{ canal <- ((s) * i)} }
}

func printa( ch chan int){
  for{
    fmt.Print( <- ch)

    fmt.Print(" e ")}
}

func merge(ch1 chan int, ch2 chan int, ch3 chan int, ch4 chan int) { 

var x1 int = <- ch1
var y2 int = <- ch2
var z3 int = <- ch3

for{ 
    
    if x1 <= y2 && x1 <= z3 {
      ch4 <- x1
      x1 = <- ch1
      
    }
    if y2 <= x1 && y2 <= z3 {
      ch4 <- y2
      y2 = <- ch2
    }
    if z3 <= y2 && z3 <= x1 {
      ch4 <- z3
      z3 = <- ch3
    }
}  
}

func main() { 
ch1 := make(chan int)
ch2 := make(chan int)
ch3 := make(chan int)

go gera(ch1, 2)
go gera(ch2, 3)
go gera(ch3, 5)

ch4 := make(chan int, 50)
go merge(ch1, ch2, ch3, ch4)


go printa(ch4)
for { }
 }