package main

import "fmt"

func main() {

	//ARITMÉTICOS (+, -, /, *, %)

	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10/4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2 //Ocorre erro, quando as  váriaveis são de tipos diferentes.Ex: se for declarado inst16 e int32
	fmt.Println(soma2)
	
	//FIM DOS ARITMÉTICOS

	//ATRIBUIÇÃO

	var variavel1 string = "String"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//FIM DOS OPERADORES DE ATRIBUIÇÃO


	//OPERADORES RELACIONAIS

	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//FIM DOS OPERADORES RELACIONAIS

	//OPERADORES LÓGICOS

	fmt.Println("------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //O && verifica se todas as condições são verdadeiras para retornar um true.

	fmt.Println(verdadeiro || falso) //O || verifica se pelo menos uma das condições é verdadeira para retornar um true.

	fmt.Println(!verdadeiro)
	fmt.Println(!falso) // O ! inverte o valor de uma variável booleana

	// FIM DOS OPERADORES LÓGICOS


	//OPERADORES UNÁRIOS

	numero := 10
	numero ++
	numero += 15 //Mesma coisa que fazer numero = numero + 15
	fmt.Println(numero)

	//Adição e subtração são as unicas que podem ser feitas desta forma

	numero--
	numero -=20 //Mesma coisa que fazer numero = numero - 20

	numero *= 3 //Mesma coisa que fazer numero = numero *3
	numero /= 10
	numero %= 3

	fmt.Println(numero)

	//FIM DOS OPERADORES UNÁRIOS

	//OPERADORES TERNÁRIOS

	var texto string
		if numero > 5 {
			texto = "Maior que 5"
		} else {
            texto = "Menor que 5"
        }
		fmt.Println(texto)

}
