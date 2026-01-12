package model

import (
	"fmt"
	"time"
)

// Presente uma classe

type Person struct {
	Name      string
	Age       int
	Address   Address
	DateBirth time.Time
}

// Instancia
func NewPerson(p Person) *Person {
	return &Person{
		Name:      p.Name,
		Age:       p.Age,
		Address:   p.Address,
		DateBirth: p.DateBirth,
	}
}

// Metodo

func (p Person) CurrentAge() int {
	return time.Now().Year() - p.DateBirth.Year()
}

// Metodo setter

func (p *Person) SetName(name string) {
	p.Name = name
	fmt.Println("Nome alterado para:", p.Name)
}
