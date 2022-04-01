//Alunos: Igor Dourado e Rafael Dias
package main

func pingOuPong(str string, in chan bool, out chan bool){
  for{
    <-in
    println(str)
    out <-true
  }
}

func main(){
  ch1 := make(chan bool)
  ch2 := make(chan bool)  
  go pingOuPong("ping", ch1, ch2)
  go pingOuPong("pong", ch2, ch1)
 
  ch1<-true
  for{ }

}