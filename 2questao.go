//Alunos: Igor Dourado e Rafael Dias
package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

const cadeiras = 7
const N = 80

var vagas = make([]reflect.Value, cadeiras)
var barbeiroVaga = make(chan int)
var sincroniza sync.WaitGroup

func main() {
	sincroniza.Add(N)
	for i := range vagas {vagas[i] = reflect.MakeChan(reflect.TypeOf(barbeiroVaga), 1)}
	go barbeiro()
	for x := 1; x <= N; x++ {
		go execuc(x)
		time.Sleep(time.Duration(rand.Intn(1e8)))}
	sincroniza.Wait()
}

func barbeiro() {
	inativo := make([]reflect.SelectCase, cadeiras+1)
	for i := 0; i < cadeiras; i++ {
		inativo[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: vagas[i],
		}
	}
	inativo[cadeiras] = reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(barbeiroVaga),
	}
	ativo := append([]reflect.SelectCase{}, inativo...)
	ativo[cadeiras] = reflect.SelectCase{
		Dir: reflect.SelectDefault,
	}
	for {
		fmt.Println("dorme o barbeiro")
		_, chega, _ := reflect.Select(inativo)
		fmt.Println("acorda o barbeiro, atende-se o cliente: ", chega.Int())
		time.Sleep(1e8)
		sincroniza.Done()
		for {
			escolha, chega, _ := reflect.Select(ativo)
			if escolha == cadeiras {
				break
			}
			fmt.Println("barbeiro chama cliente ", chega.Int(), " para atendê-lo")
			time.Sleep(1e8)
			sincroniza.Done()
		}
	}
}

func execuc(aux int) {
	time.Sleep(10)
	valores := reflect.ValueOf(aux)
	atendimentos := make([]reflect.SelectCase, cadeiras+1)
	for i := 0; i < cadeiras; i++ {
		atendimentos[i] = reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: vagas[i],
			Send: valores,
		}
	}
	atendimentos[cadeiras].Dir = reflect.SelectDefault
	select {
	case barbeiroVaga <- aux:
		fmt.Println("cliente ", aux, " encontra vaga")
	default:
		if escolha, _, _ := reflect.Select(atendimentos); escolha < cadeiras {
			fmt.Println("cliente ", aux, " espera")
		} else {
			fmt.Println("cliente ", aux, " achou a barbearia lotada e saiu")
			sincroniza.Done()
		}
	}
}
