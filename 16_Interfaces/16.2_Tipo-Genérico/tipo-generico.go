package main

import "fmt"



func generica(interf interface{}){

	fmt.Println(interf)
}


func main() {
    generica("string")
	generica(1)
	generica(true)

	fmt.Println(1, 2, "string", false, true, float64(12345)) //Cuidado pois é uma forma muito livre de desenvolver, tornando-se confuso se for usado de forma incorreta.

}