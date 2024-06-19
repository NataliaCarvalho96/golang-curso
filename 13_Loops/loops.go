package main

import (
	"fmt"
    "time"
)

func main() {
    i := 0
	 for i < 10 {
		i++
		fmt.Println("Incrementando i") //printou varias vezes na tela enquanto menor que 10
        time.Sleep(time.Second)
        

	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println("Incrementando j", j)
        time.Sleep(time.Second)
    

	}

	nomes :=[3] string{"joão", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for indice, letra := range "PALAVRAD" {
        fmt.Println(indice, string(letra))
    }
	
	usuario := map[string]string{
		"nome": "Leonardo",
		"Sobrenome": "Sabino",
	}

	for chave, valor := range usuario {
        fmt.Println(chave, valor)
    }

	//O range não pode ser usado dentro do struct.
	//type usuarioStruct struct {
		//nome string
		//sobrenome string
	
	//}
	//usuario2 := usuarioStruct{"Zé", "junior"}

	//for chave, valor := range usuario2 {
        //fmt.Println(chave, valor)
    //}

	for {
		fmt.Println("Loop infinito") //para um loop infinito, use apenas o for
        time.Sleep(time.Second)
	}
 }

