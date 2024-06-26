package main	

import "fmt"


type usuario struct {
	nome string
	idade uint8
}

func(u usuario) salvar(){
	//fmt.Println("dentro do metodo salvar")
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool{
	return u.idade >= 18
}

func(u *usuario) fazerAniversario(){
	u.idade++
}

func main() {
	usuario1 := usuario {"Usuario 1", 20}
	usuario1.salvar()

	usuario2 := usuario{"Davi", 30}
	usuario2.salvar()
	maiorDeIdade := usuario2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)

    //o metodo obrigatoriamente fica grudado em uma estrutura
	//A função possui mais liberdade nessa parte
	//Melhor usar ponteiros.
}