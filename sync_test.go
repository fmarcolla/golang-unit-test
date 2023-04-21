package main

import (
	"testing"
	"sync"
)

func TestContador(t *testing.T){
	t.Run("Incrementar o valor 3 vezes resulta no valor 3", func (t *testing.T){
		contador := Contador{}
		contador.Incrementa()
		contador.Incrementa()
		contador.Incrementa()

		expected := 3
		verificaContador(t, contador, expected)
	})
}

func TestContadorSync(t *testing.T){
	t.Run("Roda concorrência com segurança", func (t *testing.T){
		expected := 1000
		contador := Contador{}

		var wg sync.WaitGroup
		// Add é o total de goroutines
		wg.Add(expected)
		
		for i := 0; i < expected; i++ {
			go func(w *sync.WaitGroup){
				contador.Incrementa()
				// Done finaliza a goroutine
				w.Done()
			}(&wg)
		}
			
		// Aguarda todas goroutines finalizar
		wg.Wait()
		
		verificaContador(t, contador, expected)
	})
}

func verificaContador(t *testing.T, resultado Contador, esperado int) {
    t.Helper()
    if resultado.Valor() != esperado {
        t.Errorf("resultado %d, esperado %d", resultado.Valor(), esperado)
    }
}