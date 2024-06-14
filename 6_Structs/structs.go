//Structs são coleções de campos com tipos diferentes relativas a mesma coisa.

package main

import "fmt"

//Análogo a POO, seria como a Classe

type usuario struct {
	nome string
	idade uint8
	endereco endereco
}

type endereco struct { //Aqui temos um struct que será usado dentro de outro struct
	logradouro string
	numero uint8
}

func main() {
    fmt.Println("Arquivo Structs")

	var u usuario
	u.nome = "Pedro"
	u.idade = 27
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua das rosas", 14}

	usuario2 := usuario{"Pedro", 27, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{nome:"Pedro"} //nome: usado para passar apenas uma informação e ignorar a outra.
	fmt.Println(usuario3)




}