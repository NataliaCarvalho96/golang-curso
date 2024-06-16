package main

import "fmt"

func main() {
    fmt.Println("Estruturas de controle")

    numero := 10

    if numero > 15{
        fmt.Println("Maior que 15")
    }else {
        fmt.Println("Menor ou igual a 15")
    }

    if outroNumero :=numero; outroNumero > 0{
        fmt.Println("Número maior que zero")
    }else if numero < - 10 {
        fmt.Println("Número menor que -10")
    }else{
        fmt.Println("Entre 0 e -10") //Pode ser usado o apenas o else ou else if
    }


}