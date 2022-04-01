package main

import "fmt"

//---------------------------

var x, y, z int = 0, 0, 0

func px() {
	x = 1
 	x = 2
 }

func py() {
	y = 1
 	y = 2
 }

func pz() {
	z = 1
 	z = 2
 }

func questaoStSp2() {
	go px()
	py()
	for {
	}
}


func questaoStSp3() {
	go px()
	go py()
        pz()
	for {
	}
}
//---------------------------

// considerando-se como estados os valores da tripla x,y,z 
// qual o diagrama de estados e transicoes que representa
// questaoStSp2()  ?
//
// e qual representa questaoStSp3()  ?

func main() {
	questaoStSp2()
	//questaoStSp3()

}
