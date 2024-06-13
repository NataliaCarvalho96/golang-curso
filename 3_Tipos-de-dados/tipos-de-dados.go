package main
 
import (
	
	"fmt"
	"errors"

)

func main() {
	//int8, int16, int32, int64

	var numero int = 1996
	fmt.Println(numero)

	//uint8, uint16, uint32, uint64 --> é para numeros sem sinal
	var numero2 uint32 = 2208
	fmt.Println(numero2)
	//apelidos dos ints, int32 abaixo, usados para numeros que representam caracteres
	var numero3 rune = 1404
	fmt.Println(numero3)

	//Byte = uint8
	var numero4 byte = 14
	fmt.Println(numero4)


	//Números reais, float32 e float64
	var numeroReal1 float32 = 14.04
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 22.08
	fmt.Println(numeroReal2)

	numeroReal3 := 140422.08
	fmt.Println(numeroReal3)

	//Fim dos numeros reais.

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//Não existe char em GO, o char vai sempre representar algum caracter da tabela ASCII
	char := 'A'
	fmt.Println(char)

	//Fim das Strings

	//Valor Zero --> É o valor que será atribuido a uma variavel quando inicializada sem um valor.

	var texto string
	fmt.Println(texto)

	var numeroZero int16
	fmt.Println(numeroZero)

	//O retorno para strings é uma espaço (de string) vazio, e para númeors é o 0.

	//Fim dos Valores Zero.

	//Booleano e Error

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error
	fmt.Println(erro) //retorna nil se valor não indicado

	var erro2 error = errors.New("Erro interno")//É necessario usar o pacote 'errors.' para atribbuir um erro específico.
	fmt.Println(erro2)

	
	
}

