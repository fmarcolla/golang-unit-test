package main

import "testing"

func TestVerificaPar(t *testing.T){
	
	testes := []struct {
		value int
		expected bool
	}{
		{value: 2, expected: true},
		{value: 3, expected: false},
		{value: 0, expected: true},
		{value: 10, expected: true},
		{value: 11, expected: false},
		{value: -1, expected: false},
		{value: -2, expected: false},
	}

	for _, teste := range testes {
		result := VerificaPar(teste.value)
	
		if result != teste.expected {
			t.Fatalf("Valor testado: %d -> Resultado esperado: %t -> Resultado Ocorrido: %t", teste.value, teste.expected, result)
		}
	}
}