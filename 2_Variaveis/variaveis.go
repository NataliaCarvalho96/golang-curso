package main

import (
    "fmt"
    
)

func main()  {
	var variavel1 string = "Variavel 1"
	//Inferência de tipo quando usado :=
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var(

		variavel3 string = "Melancia"
		variavel4 string = "Maçã"

	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Morango", "Mamão"
	fmt.Println(variavel5, variavel6)
	
	const constante1 string = "Constante 1"
	fmt.Println(constante1)

	variavel7, variavel8 := "Variavel7", "Variavel8"
	fmt.Println(variavel7, variavel8)
	
	//Valores invertidos
	variavel7, variavel8 = variavel8, variavel7
	fmt.Println(variavel7, variavel8)



}