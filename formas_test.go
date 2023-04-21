package main

import "testing"

func TestCalculoAreaFormas(t *testing.T){

	testesArea := []struct{
		nome string
		forma Forma
		expected float64
	}{
		{"Triângulo", Triangulo{12.0, 6.0}, 36.0},
		{"Cirulo", Circulo{10.0}, 314.1592653589793},
		{"Retângulo", Retangulo{12.0, 6.0}, 72.0},
	}

	for _, teste := range testesArea {
		result := teste.forma.Area()
		
		if result != teste.expected {
			t.Errorf("Forma: %s  -- resultado %.2f, esperado %.2f", teste.nome, result, teste.expected)
		}
	}
}
