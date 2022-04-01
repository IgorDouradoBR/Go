//Alunos: Igor Dourado e Rafael Dias

package main


import (
    "fmt"
    "time"
)




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

var elves int = 0
var reindeer int = 0
var santaSem = NewSemaphore(0)
var reindeerSem = NewSemaphore(0)
var elfTex = NewSemaphore(1)
var mutex = NewSemaphore(1)

const N_elves= 21
const N_reindeer= 9

func main(){
  go Santa()
  for i:=0;i<N_reindeer;i++{
    go ReindeerFunc()
  }
  for j:=0;j<N_elves;j++{
    go ElvesFunc()
  }

  for{}
}

func Santa(){

  for{

    santaSem.Wait()
    mutex.Wait()
    if reindeer == 9{
      fmt.Println("Prepare sleigh")
      for i:=0;i<9;i++{reindeerSem.Signal()}
      reindeer=0
    }else if elves==3{
      fmt.Println("Help elves")
    }
    mutex.Signal()

  }
}

//Não fiz um for infinito dentro desse e do próximo método pois estava travando quando tentei rodar em concorrência, fiz apenas um for infinito dentro do metodo Santa como indicado pelo livro
func ReindeerFunc(){
      mutex.Wait()
      reindeer += 1
      if reindeer==9{
        santaSem.Signal()
      }
      
      mutex.Signal()
            time.Sleep(1e8)

      reindeerSem.Wait()
      fmt.Println("Get hitched")
}

func ElvesFunc(){
    elfTex.Wait()
    mutex.Wait()
      elves += 1
      if elves==3{
        santaSem.Signal()
      }else{elfTex.Signal()}
    mutex.Signal()
    fmt.Println("Elves get help")
    time.Sleep(1e8)
    mutex.Wait()
      elves -=1
      if elves == 0{
        elfTex.Signal()
      }
      mutex.Signal()
  
}

