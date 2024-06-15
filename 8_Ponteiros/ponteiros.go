package main

import "fmt"

func main() {

	var variavel1 int = 10 
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA

	var variavel3 int //GUARDA UM VALOR INTEIRO
	var ponteiro *int //GUARDA O ENDEREÇO DE MEMÓRIA DE UM VALOR INTEIRO

	variavel3 = 100
	ponteiro = &variavel3 //COM O & É POSSÍVEL USAR O PONTEIRO CHAMANDO A VARIÁVEL

	fmt.Println(variavel3, *ponteiro) //desreferenciação --> Usando o * irá aparecer o valor da variável e não a referencia de memória onde o ponterio está gravado.

}