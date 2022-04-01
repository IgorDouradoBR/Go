//Alunos: Igor Dourado e Rafael Dias

package main

import (
	"fmt"
)


// -----------------------------------------
// Definição de Barreira -------------------
// -----------------------------------------

// barreira:  Tam processos devem chegar em um ponto chamado barreira, para então todos seguirem adiante
// do￼￼ livro The Little book of Semaphores￼￼

// turnstile=Semaphore(0)
// turnstile2=Semaphore(1)
// mutex=Semaphore(1)

// Incrementa Tam e veja se todos chegaram. se chegaram, programa turnstile2 para bloquear processos depois de passar
// por turnstile, e sinaliza UM processo para prosseguir
// mutex.wait()
//    count += 1
// if count == n:
//    turnstile2.wait()       // turnstile2 permitia um processo, este passa e os demais bloqueiam
//    turnstile.signal()      // acorda um processo em turnstile
// mutex.signal()

// turnstile.wait()           // processo que acorda, libera outro:
// turnstile.signal()         // o padrao catraca de uso de semaforos

// #criticalpoint

// mutex.wait()
//    count -= 1
//    if count == 0:
//       turnstile.wait()    // ultimo processo da barreira gerou um signal a mais vide acima. aqui zera este semaforo para ninguem passar
//       turnstile2.signal() // ultimo processo libera a catraca turnstile2 e assim todos os processos podem reusar a barreira
// mutex.signal()
//
// turnstile2.wait()
// turnstile2.signal()
// mutex.wait()

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
		fmt.Println("\nMatriz após a ",(cont/16)+1, "ª iteração: \nLinha x Coluna") // para mostrar o ponto critico em relacao aos prints dos processos no exemplo abaixo
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

// -----------------------------------------
// Fim Definição de Barreira ---------------
// -----------------------------------------

// Exemplo de uso a barreira
// Os processos tem cada um um identificador
// eles escrevem o identificador antes e depois do ponto crítico
// como "   Ponto Critico" é escrito na operacao Arrive, pelo ultimo processo
// que chega no arrive, e como "   Ponto Nao Critico" pelo último que sai
// do ponto crítico, cada processo so imprime seu ID uma vez entre
// um e outro (ponto critico e nao critico)


const Tam = 4 //Tamanho da matriz
var matriz = [Tam][Tam]float64{//alterar manualmente caso desejar outros valores ou um maior tamanho, para facilitar os testes de precisão
		{4, 6, 2, 3},
		{2, 9, 4, 7},
		{3, 8, 1, 9},
		{5, 2, 6, 3}}
var f chan struct{}
const Repeticoes = 4 //Número de vezes que irá iterar

func procsComBarreiras() {
	b := NewBarrier(Tam*Tam)
	for i := 0; i < Tam; i++ {
		for j := 0; j < Tam; j++ {
			go useBarrier(b, i, j)
		}
	}
}

func useBarrier(b *Barrier, linha int, coluna int) { 
  //fase 1
	for i := 0; i < Repeticoes; i++ {
    var norte float64= 0
    var sul   float64= 0
    var leste float64= 0
    var oeste float64= 0
    if linha-1 >= 0 {	norte = matriz[linha-1][coluna]
			}else{norte=matriz[Tam-1][coluna]}

			if linha+1 < Tam {	sul = matriz[linha+1][coluna]
			}else{sul=matriz[0][coluna]}

			if coluna+1 < Tam {	leste = matriz[linha][coluna+1]
			}else{leste=matriz[linha][0]}

			if coluna-1 >= 0 {	oeste = matriz[linha][coluna-1]
			}else{oeste=matriz[linha][Tam-1]}

		b.Arrive()
    //fase 2
    matriz[linha][coluna] = (norte + sul + leste + oeste)/4

		fmt.Println("Célula", linha, "x", coluna, "=", matriz[linha][coluna])
    
		b.Leave()
    f <- struct{}{}
    cont++	
	}
}

func main() {
    aux:= Repeticoes*(Tam*Tam)
    fmt.Println("Matriz ",Tam,"x", Tam,"na forma original: ")
    for x := 0; x < Tam; x++ {
			for j := 0; j < Tam; j++ {
        fmt.Print(matriz[x][j]," ")
        if j==3 {fmt.Print("\n")}
      }
    }
	procsComBarreiras()
	f = make(chan struct{}, aux)
	for i := 0; i<aux; i++{<-f}
}

