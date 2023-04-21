package main

import "testing"

func TestNewPersonPositive(t *testing.T){
	person := NewPerson(1)

	if person == nil {
		t.Fatalf("Tem que deixar criar uma pessoa com idade positiva")
	}
}

func TestNewPersonNegative(t *testing.T){
	person := NewPerson(-1)

	if person != nil {
		t.Fatalf("Não pode deixar criar uma pessoa com idade negativa")
	}
}

func TestOlderPersonSecondOlder(t *testing.T){
	person1 := NewPerson(10)
	person2 := NewPerson(20)

	if person1.older(person2) {
		t.Fatalf("Esperado que Pessoa2 seja mais velha que a pessoa1")
	}
}

func TestOlderPersonFirstOlder(t *testing.T){
	person1 := NewPerson(30)
	person2 := NewPerson(20)

	if !person1.older(person2) {
		t.Fatalf("Esperado que Pessoa1 seja mais velha que a pessoa2")
	}
}