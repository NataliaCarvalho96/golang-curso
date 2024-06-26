package main

import "fmt"


func diaDaSemana(numero int) string {
	switch  numero {
	case 1:
        return "Domingo"
    case 2:
        return "Segunda-feira"
     case 3:
        return "Terça-feira"
    case 4:
        return "Quarta-feira"
    case 5:
        return "Quinta-feira"
    case 6:
        return "Sexta-feira"
    case 7:
        return "Sábado"
    default:
        return "Dia inválido"
    }
	
		
}



func diaDaSemana2(numero int) string {

	var diaDaSemana string

	switch{
	case numero ==1:
		diaDaSemana =  "Domingo"
		fallthrough //Joga o codigo para dentro da proxima condição, no caso segunda-feira
	case numero ==2:
		diaDaSemana =  "Segunda-feira"
    case numero ==3:
		diaDaSemana =  "Terça-feira"
	case numero ==4:
		diaDaSemana =  "Quarta-feira"
    case numero ==5:
        diaDaSemana =  "Quinta-feira"
    case numero ==6:
        diaDaSemana =  "Sexta-feira"
    case numero ==7:
        diaDaSemana =  "Sábado"
	default:
        diaDaSemana =  "Dia inválido"
	}
	
	return diaDaSemana
}

func main() {
    fmt.Println("Switch!")

	dia := diaDaSemana(3)
	fmt.Println(dia)
	
	fmt.Println("----------")

    dia2 := diaDaSemana2(1)
	fmt.Println(dia2)

}