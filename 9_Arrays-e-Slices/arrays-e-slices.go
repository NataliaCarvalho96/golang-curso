package main

import (
	"fmt"
	//"reflect"
)

func main() {
    fmt.Println("Arrays e Slices!")

	var array1 [5]string
	array1[0] = "Posição 1"
	fmt.Println(array1)

	array2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}
	fmt.Println(array2)
	
	array3 := [...] int{1, 2, 3, 4, 5}
	fmt.Println(array3) //O array é mais inflexível e não muito utilizado nas aplicações em geral

	slice := []int {10, 11, 12, 13, 14, 15, 16, 17} // tem tamanho dinâmico e não sofre limitações, ainda respeitando a tipagem.
	fmt.Println(slice)

	slice = append(slice, 18)
	fmt.Println(slice)

	slice2 := array2[1:3] //O primeiro indice é inclusivo e o segundo indice exclui, neste caso pega da posição 1, até a posição 2, pois a posição 3 é excluida, por isso o retorno é a posição 2 e posição 3.
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2) //o slice nesse caso funciona similar a um ponteiro.

	//Arrays internos
	fmt.Println("----------------")

	slice3 := make([] float32, 10, 11) //Neste exemplo a função 'make' cria um array de 15 posições que pega as 10 primeiras posições.

	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Length
	fmt.Println(cap(slice3)) //Capacity

	fmt.Println("----------------")
	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) 
	fmt.Println(cap(slice3))
	// No exemplo acima, o slice estouraria o tamanho indicado, sendo assim o GO cria outro array para referência e dobra de tamanho, e isso é contínuo. Por esse motivo, o slie não possui um tamanho limite, pois, o mesmo continua a aumentar sempre que necessário.





}