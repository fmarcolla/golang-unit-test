package main

import "math"

type Forma interface {
	Area() float64
}

type Triangulo struct {
	base float64
	altura float64
}

func (t Triangulo) Area() float64 {
	return (t.base * t.altura) * 0.5
}

type Retangulo struct {
	largura float64
	altura float64
}

func (r Retangulo) Area() float64 {
	return r.largura * r.altura
}

type Circulo struct {
	raio float64
}

func (c Circulo) Area() float64 {
	return c.raio * c.raio * math.Pi
}