package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup
	waitGroup.Add(2) //Especifica a quantidade de rotinas a rodar

	go func (){
		escrever("Olá mundo!")
		waitGroup.Done()
	}()

	go func (){
		escrever("Programando em GO")
		waitGroup.Done()
	}()

	waitGroup.Wait() //Espera a duas rotinas terminarem para prosseguir


}

func escrever(texto string) {
	for  i := 0; i< 5; i++{
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
