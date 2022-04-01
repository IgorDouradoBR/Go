/*
Exercícios Avaliativos - Semáforos - MC-3:
             
(1) Implemente o problema dos leitores-escritores com semáforos,
veja no livro "Little book of semaphores".

Conforme este problema, leitores e escritores acessam um recurso R da seguinte forma:
Um escritor tem que acessar R em exclusão mútua com qualquer outro processo, leitor ou
escritor. Um leitor pode acessar R concorrentemente com outros leitores.   
Assim, se um escritor está esperando que um leitor acabe a leitura, pode chegar outro leitor
e iniciar a sua leitura. Como isto pode acontecer sucessivas vezes, o escritor pode ser 
postergado indefinidamente.
Como você pode implementar esta semântica com semáforos ?
Avalie se processos podem sofrer de inanição (postergação indefinida, ou starvation).
Se isto puder ocorrer, como poderia ser evitado ?
*/

package main

import (
	"bytes"
	"log"
	"sync"

	"github.com/soniakeys/LittleBookOfSemaphores/sem"
)

var (
	readSwitch = sem.NewLightswitch()
	roomEmpty  = sem.NewChanSem(1, 1)
	turnstile  = sem.NewChanSem(1, 1)
	b          bytes.Buffer // protegido por roomEmpty
)

var wg sync.WaitGroup

func writer(nw int) {
	turnstile.Wait()
	roomEmpty.Wait()
	log.Println("escritor", nw, "escreve")
	b.WriteString("w")
	turnstile.Signal()
	roomEmpty.Signal()
	wg.Done()
}

func reader(nr int) {
	turnstile.Wait()
	turnstile.Signal()
	readSwitch.Lock(roomEmpty)
	log.Println("leitor", nr, "vê", b.Len(), "bytes")
	readSwitch.Unlock(roomEmpty)
	wg.Done()
}
const n =  200
const nw = n
const nr = n

func main() {
	wg.Add(nw + nr)
	for i := 1; i <= nw; i++ {
		go writer(i)
	}
	for i := 1; i <= nr; i++ {
		go reader(i)
	}
	wg.Wait()
}
