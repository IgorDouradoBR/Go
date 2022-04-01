package main

import (
	"fmt"
	"time"
)
//
//func ping(entrada chan bool, saida chan bool) {
//	for {
//		<-entrada
//		println("ping")
//		saida <- true
//	}
//}
//
//func pong(entrada chan bool, saida chan bool) {
//	for {
//		<-entrada
//		println("pong")
//		saida <- true
//	}
//}
//
//func main() {
//	pingA := make(chan bool)
//	pongB := make(chan bool)
//
//	go ping(pingA, pongB)
//	go pong(pongB, pingA)
//
//	pingA <- true
//
//	for {
//	}
//}


type Ball struct{ hits int }

func player(name string, table chan *Ball) {
	for {
		ball := <-table // player grabs the ball
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(10 * time.Millisecond)
		table <- ball // pass the ball
	}
}

func main() {
	table := make(chan *Ball)

	go player("ping", table)
	go player("pong", table)
	table <- new(Ball) // game on; toss the ball
	time.Sleep(10 * time.Millisecond)
	<-table // game over, grab the ball
	panic("show me the stack")
}