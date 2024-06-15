package main

import "fmt"

func main() {

	usuario := map[string]string{
		"nome":  	"Pedro",
        "Sobrenome": "Silva",

	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "joão",
			"sultimo": "pedro",
		},
		"cursor": {
			"nome": "Engenharia",
			"Campus": "Campus 1",
		},
		
	}	
	fmt.Println(usuario2)
	delete(usuario, "nome")//usado para deletar uma chave, neste caso "nome" foi deletado.
	fmt.Println(usuario2)

}