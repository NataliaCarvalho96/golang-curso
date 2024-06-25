package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	TipoDeEndereco := enderecos.TipoDeEndereco("Rodovia dos imigrantes")
	fmt.Println(TipoDeEndereco)
}