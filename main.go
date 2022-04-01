//Alunos: Igor Dourado e Rafael Dias 

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const N = 8 // nro de nodos
const D = 7 // diametro

type Topologia struct {
	orig int    // id origem [0..N-1]
	cont string // conteudo
}

type Message struct {
	orig int // id origem [0..N-1]
	id   int // nro sequencia da mensagem do orig para o dest
	cont int // conteudo
}

func main() {
	net := [N][N]bool{ //  matriz de conexao entre vizinhos:  i conectado com j implica em j conectado com i
		{false, false, false, true, false, false, false, false},
		{false, false, true, true, false, false, false, false},
		{false, true, false, false, true, true, false, false},
		{true, true, false, false, true, false, false, false},
		{false, false, true, true, false, true, false, false},
		{false, false, true, false, true, false, false, false},
		{false, false, false, false, true, false, false, true},
		{false, false, false, false, true, false, true, false}}

	var topologias [N]chan Topologia // cria o canal de entrada de cada nada para montar topologia
	for i := 0; i < N; i++ {
		topologias[i] = make(chan Topologia, N)
	}

	for i := 0; i < N; i++ { // lanca os canais
		go topologia(i, net, topologias)
	}


	var mensagens [N]chan Message // cria o canal de entrada de cada nodo para trocar mensagens
	for i := 0; i < N; i++ {
		mensagens[i] = make(chan Message, N)
	}

	for i := 0; i < N; i++ { // lanca canais para mensagens
		go mensagem(i, net, mensagens)
	}

	for i := 1; ; i++ { // gera mensagems de teste aleatórias
		n := rand.Intn(N)
		c := rand.Intn(100)
		mensagens[n] <- Message{n, i, c}
		time.Sleep(time.Second)
	}


}

//  montagem da topologia
func topologia(id int, net [N][N]bool, topologias [N]chan Topologia) {
	fmt.Println(id, " ativo! ")
	for i := 0; i < D; i++ {

		// Fase 1 envia vizinhos
		m := Topologia{id, ("Topologia: " + strconv.Itoa(i))} // quem recebe de quem
		for j := 0; j < N; j++ {                              // passa adiante para os vizinhos
			if net[id][j] == true { // se i conectado com j
				topologias[j] <- m // coloca no canal de j
			}
		}

		// Fase 2 recebe  vizinhos
		for j := 0; j < N; j++ { // para todos os vizinhos j
			if net[j][id] == true {
				m := <-topologias[id] // aguarda uma entrada
				fmt.Println(i, " recebe de: ", m)
				net[id][j] = true // updatea a topologia
			}
		}
	}
}


// troca de mensagens
func mensagem(i int, net [N][N]bool, mensagens [N]chan Message) {
	fmt.Println(i, " ativo! ")
	for {
		m := <-mensagens[i] // espera entrada
		fmt.Println(i, " está verificando a mensagem ", m)
		for j := 0; j < N; j++ {
			if i == m.orig {
				println("Recebida a mensagem")
				break
			} else if net[i][j] == true {
				mensagens[j] <- m // passa adiante
			}

		}
	}
}