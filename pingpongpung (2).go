//Alunos: Igor Dourado e Rafael Dias
package main

func pinng(entrada chan bool, saida chan bool, id int) {
	for {
		<-entrada
		switch id {
		case 1:
			println("ping")
			break
		case 2:
			println("pong")
			break
		case 3:
			println("pung")
			break
		}

		saida <- true
	}
}

func main() {
	pingA := make(chan bool)
	pongB := make(chan bool)
	pungB := make(chan bool)

	go pinng(pingA, pongB, 1)
	go pinng(pongB, pungB, 2)
	go pinng(pungB, pingA, 3)

	pingA <- true

	for {
	}
}