package main

import "fmt"

func funcao1() {

	fmt.Println("Executando da função 1")

}

func funcao2() {

	fmt.Println("Executando da função 2")

}

func alunoEstaAprovado(n1, n2 float32) bool{
	defer fmt.Println("Média calculada. Resultado será retornado")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		fmt.Println("O aluno está aprovado")
        return true

    }

	fmt.Println("O aluno está reprovado")

    return false

}


func main() {
    defer funcao1() //Defer = Adiar, adia a execução dessa função até o ultimo momento possível
    funcao2()

	fmt.Println(alunoEstaAprovado(7, 8))
}