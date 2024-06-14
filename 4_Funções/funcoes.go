package main

//as funções em GO podem ter parâmetros [dados que vc colaca na função para ela funcionar] e retornos [que é o que a função retorna para você].

import "fmt"

func somar (n1 int8, n2 int8) int8 {
	return n1 + n2
}

	func calculosMatematicos(n1, n2 int8)(int8, int8){
		soma := n1 + n2
		subtracao := n1 - n2
		return soma, subtracao
	}

func main() {

	soma:= somar(14, 22)
	fmt.Println(soma)
	
	//Aqui foi declarado uma variavel do tipo [função] e colocado outra função dentro dela.
	//var f = func(){
	//fmt.Println("Função F")
	//}
	//f()

	var f = func(txt string) string {
		fmt.Println(txt)
		return txt
	}
		resultado := f("Texto da função 1")
		fmt.Println(resultado)

		//*resultadosCalculos :=calculosMatematicos(10, 15)//Erro pois a função retorna dois valores, porem apenas uma variavel foi declarada.

		resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
		fmt.Println(resultadoSoma, resultadoSubtracao) //Usar o _ no lugar do resultado, faz com que o GO entenda que você quer apenas um dos resultados, não gerando um erro.


}