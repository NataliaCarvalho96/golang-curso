package main


import (
	"fmt"
	//"sync"
	"time"
)
func main(){
	canal := make(chan string)

	go escrever("Olá mundo!", canal)

	fmt.Println("Depois da função escrever começar a ser executada")


	for mensagem := range canal{
		fmt.Println(mensagem)
	}

}


func escrever(texto string, canal chan string) {
	time.Sleep(time.Second * 5)
	for  i := 0; i< 5; i++{
		canal <- texto // mandando um valor para dentro do canal
		time.Sleep(time.Second)
	}

	close(canal)
}
