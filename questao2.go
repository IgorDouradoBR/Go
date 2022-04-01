//Alunos: Igor Dourado e Rafael Dias

package main

import (
	"fmt"
	"math/rand"
)

type LightSwitch struct {
	counter int
	mutex   Semaphore
}

func __init__() *LightSwitch {
	self := &LightSwitch{counter: 0, mutex: *NewSemaphore(1)}
	return self
}

func (self *LightSwitch) __lock__(semaphore *Semaphore) {
	self.mutex.Wait()
	self.counter++
	if self.counter == 1 {
		semaphore.Wait()
		self.mutex.Signal()
	}
}

func (self *LightSwitch) __unlock__(semaphore *Semaphore) {
	self.mutex.Wait()
	self.counter--
	if self.counter == 0 {
		semaphore.Signal()
		self.mutex.Signal()
	}
}

type Semaphore struct { // este semáforo implementa quaquer numero de creditos em "v"
	v    int           // valor do semaforo: negativo significa proc bloqueado
	fila chan struct{} // canal para bloquear os processos se v < 0
	sc   chan struct{} // canal para atomicidade das operacoes wait e signal
}

func NewSemaphore(init int) *Semaphore {
	s := &Semaphore{
		v:    init,                   // valor inicial de creditos
		fila: make(chan struct{}),    // canal sincrono para bloquear processos
		sc:   make(chan struct{}, 1), // usaremos este como semaforo para SC, somente 0 ou 1
	}
	return s
}

func (s *Semaphore) Wait() {
	s.sc <- struct{}{} // SC do semaforo feita com canal
	s.v--              // decrementa valor
	if s.v < 0 {       // se negativo era 0 ou menor, tem que bloquear
		<-s.sc               // antes de bloq, libera acesso
		s.fila <- struct{}{} // bloqueia proc
	} else {
		<-s.sc // libera acesso
	}
}

func (s *Semaphore) Signal() {
	s.sc <- struct{}{} // entra sc
	s.v++
	if s.v <= 0 { // tem processo bloqueado ?
		<-s.fila // desbloqueia
	}
	<-s.sc // libera SC para outra op
}

type Barrier struct {
	mutex    *Semaphore
	catraca1 *Semaphore
	catraca2 *Semaphore
	count    int
	n        int
}

func NewBarrier(val int) *Barrier {
	b := &Barrier{
		mutex:    NewSemaphore(1),
		catraca1: NewSemaphore(0),
		catraca2: NewSemaphore(1),
		count:    0,
		n:        val,
	}

	return b
}

var cont int = 0

func (b *Barrier) Arrive() {
	b.mutex.Wait()
	b.count++
	if b.count == b.n {
		b.catraca2.Wait()
		b.catraca1.Signal()
		fmt.Println("Sessao critica") // para mostrar o ponto critico em relacao aos prints dos processos no exemplo abaixo
	}
	b.mutex.Signal()
	b.catraca1.Wait()
	b.catraca1.Signal()

}

func (b *Barrier) Leave() {
	b.mutex.Wait()
	b.count--
	if b.count == 0 {
		b.catraca1.Wait()
		b.catraca2.Signal()
	}
	b.mutex.Signal()
	b.catraca2.Wait()
	b.catraca2.Signal()
}

var insertMutex = NewSemaphore(1)
var noSearcher = NewSemaphore(1)
var noInserter = NewSemaphore(1)
var searchSwitch = *__init__()
var insertSwitch = *__init__()

const TAM_LISTA_INIC = 50
const NUM_MAX_ELEM = 500
const PROC_INS = 2000
const PROC_SEAR = 200
const PROC_DEL = 2

var lista []int

func main() {
	__preencheLista__(lista)
  //abri mais processos pq ficou inviável colocar um for infinito dentro dos métodos, pois estava travando em minha máquina
	for i := 0; i < 200; i++ {
		go inserter(lista)
		go deleter(lista)
		go searcher(lista)
	}
  for{}
}

func __preencheLista__(list []int) {
	for i := 0; i < TAM_LISTA_INIC; i++ {
		entra := rand.Intn(NUM_MAX_ELEM)
		lista = append(lista, entra)
	}
}

//Não fiz um for infinito dentro dos métodos pois estava travando quando tentei rodar em concorrência, então abri mais processos no main para compensar, caso apareçam poucas linhas de resultado ao rodar o programa, rodar novamente por favor
func deleter(list []int) {
		removeIndex := rand.Intn(len(lista))
		noSearcher.Wait()
		noInserter.Wait()

		fmt.Println("Removido o elemento: ", lista[removeIndex], " na posição ", (removeIndex + 1), "º")
		lista = append(lista[:removeIndex], lista[removeIndex+1:]...)
		noInserter.Signal()
		noSearcher.Signal()
	
}

func inserter(list []int) {
		insere := rand.Intn(NUM_MAX_ELEM)
		insertSwitch.__lock__(noInserter)
		insertMutex.Wait()
		lista = append(lista, insere)
		fmt.Println("Elemento", insere, "adicionado ao final da lista")
		insertMutex.Signal()
		insertSwitch.__unlock__(noInserter)
	
}

func searcher(list []int) {
		busca := rand.Intn(NUM_MAX_ELEM)
		searchSwitch.__lock__(noSearcher)
		for j := 0; j < len(lista); j++ {
			if lista[j] == busca {
				fmt.Println("Na", (j + 1), "º posição do vetor, foi encontrado o número ", busca)
				break
			}
		}
		searchSwitch.__unlock__(noSearcher)

}
