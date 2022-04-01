package main

func pxng(entrada chan bool, saida chan bool, id int) {
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

	go pxng(pingA, pongB, 1)
	go pxng(pongB, pungB, 2)
	go pxng(pungB, pingA, 3)

	pingA <- true

	for {
	}
}
