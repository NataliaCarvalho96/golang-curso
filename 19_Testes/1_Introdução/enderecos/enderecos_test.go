//Teste de unidade

package enderecos

import (
	
	"testing"
)

func TestTipoDeEndereco(t *testing.T){
	enderecoParaTeste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "Avenida"
	TipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if TipoDeEnderecoRecebido != tipoDeEnderecoEsperado{
		t.Error("O tipo recebido é diferente do esperado")
	}

}