/*
(2)
               Implemente o problema do barbeiro dorminhoco com semáforos
               Veja a descrição no pequeno livro dos semáforos. 
              Uma barbearia tem um único barbeiro.  Tem uma cadeira de corte, e tem N cadeiras de
              espera.    Se o barbeiro não tem clientes, ele dorme na sua cadeira.   Quando um cliente
              chega e a loja está cheia, ele desiste.    Senão ele espera em uma das cadeira.   Se
              o barbeiro estiver dormindo, o cliente acorda o barbeiro para fazer seu corte.
              Após acordado, o barbeiro corta o cabelo do cliente e verifica se há mais clientes.
              Havendo mais clientes ele segue cortando.   Quando não há clientes ele volta a dormir.
*/

package main

import (
	"log"
	"math/rand"
	"reflect"
	"sync"
	"time"
)

const nWRChairs = 6

var wrChairs = make([]reflect.Value, nWRChairs)
var barberChair = make(chan int)
var wg sync.WaitGroup

const nCust = 80

func main() {
	wg.Add(nCust)
	for i := range wrChairs {
		wrChairs[i] = reflect.MakeChan(reflect.TypeOf(barberChair), 1)
	}
	go barber()
	for c := 1; c <= nCust; c++ {
		go customer(c)
		time.Sleep(time.Duration(rand.Intn(1e8)))
	}
	wg.Wait()
}

func barber() {
	sleepingCases := make([]reflect.SelectCase, nWRChairs+1)
	for i := 0; i < nWRChairs; i++ {
		sleepingCases[i] = reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: wrChairs[i],
		}
	}
	sleepingCases[nWRChairs] = reflect.SelectCase{
		Dir:  reflect.SelectRecv,
		Chan: reflect.ValueOf(barberChair),
	}
	awakeCases := append([]reflect.SelectCase{}, sleepingCases...)
	awakeCases[nWRChairs] = reflect.SelectCase{
		Dir: reflect.SelectDefault,
	}
	for {
		log.Print("barbeiro dormindo")
		_, recv, _ := reflect.Select(sleepingCases)
		log.Print("barbeiro acordado e atendendo cliente ", recv.Int())
		time.Sleep(1e8)
		wg.Done()
		// awake
		for {
			chosen, recv, _ := reflect.Select(awakeCases)
			if chosen == nWRChairs {
				break
			}
			log.Print("barbeiro atendendo cliente à espera ", recv.Int())
			time.Sleep(1e8)
			wg.Done()
		}
	}
}

func customer(c int) {
	time.Sleep(1e7)
	cVal := reflect.ValueOf(c)
	cases := make([]reflect.SelectCase, nWRChairs+1)
	for i := 0; i < nWRChairs; i++ {
		cases[i] = reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: wrChairs[i],
			Send: cVal,
		}
	}
	cases[nWRChairs].Dir = reflect.SelectDefault
	select {
	case barberChair <- c:
		log.Print("cliente ", c, " feliz de encontrar um barbeiro livre")
	default:
		if chosen, _, _ := reflect.Select(cases); chosen < nWRChairs {
			log.Print("cliente ", c, " espera")
		} else {
			log.Print("cliente ", c, " encontrou barbearia cheia, saiu")
			wg.Done()
		}
	}
}
