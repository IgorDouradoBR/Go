//Alunos: Igor Dourado e Rafael Dias
package main

func ping(entrada chan bool, saida chan bool) {
	for {
		<-entrada
		println("ping")
		saida <- true
	}
}

func pong(entrada chan bool, saida chan bool) {
	for {
		<-entrada
		println("pong")
		saida <- true
	}
}

func main() {
	pingA := make(chan bool)
	pongB := make(chan bool)

	go ping(pingA, pongB)
	go pong(pongB, pingA)

	pingA <- true

	for {
	}
}
