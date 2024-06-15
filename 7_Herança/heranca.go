package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	idade uint8
	altura uint8
	
}

//Todos os campos que estão no struct de pessoa, foram "herdados" dentro do struc estudante, sendo assim as informções serão consideradas quandoa acessadas. *Não é necessário especificar um tipo.

type estudante struct {
	pessoa
	curso string
	faculade string


}


func main() {
    fmt.Println("Herança")

	p1 := pessoa{"Natalia", "Carvalho", 27, 165}
	fmt.Println(p1)

	e1 := estudante{p1, "Análise e Desenvolvimento de Sistemas", "Descomplica"}
	fmt.Println(e1)
}