package main

import "sync"

type Contador struct {
	mu sync.Mutex
	valor int
}

// Mutex é uma trava, que garante que varias go routines não tentem alterar o "valor" ao mesmo tempo 
func (c *Contador) Incrementa() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

func (c *Contador) Valor() int {
	return c.valor
}