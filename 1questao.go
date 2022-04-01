//Alunos: Igor Dourado e Rafael Dias

package main

import (
	"bytes"
	"fmt"
	"sync"
	"github.com/soniakeys/LittleBookOfSemaphores/sem"
)

const max =  200

var sincroniza sync.WaitGroup

var (vazio= sem.NewChanSem(1, 1)	
	   catraca= sem.NewChanSem(1, 1)
     leitura= sem.NewLightswitch()
	   buffer bytes.Buffer)

func main() {
	sincroniza.Add(max*2)
	for i := 1; i <= max; i++ {	go writer(i)	}
	for i := 1; i <= max; i++ {	go reader(i)	}
	sincroniza.Wait()
}


func writer(aux int) {
	catraca.Wait()
	vazio.Wait()
	fmt.Println( aux, "º escritor faz escrita")
	buffer.WriteString("w")
	catraca.Signal()
	vazio.Signal()
	sincroniza.Done()
}

func reader(aux int) {
	catraca.Wait()
	catraca.Signal()
	leitura.Lock(vazio)
	fmt.Println(aux, "º leitor enxerga: ", buffer.Len(), "Bytes")
	leitura.Unlock(vazio)
	sincroniza.Done()
}


