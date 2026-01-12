package main

import (
	"fmt"
	"structs/model"
	"time"
)

func main() {
	person := model.NewPerson(model.Person{
		Name: "Jonh Doe",
		Age:  23,
		Address: model.Address{
			Street: "Rua X",
			Number: 134,
			City:   "X",
		},
		DateBirth: time.Date(2002, 11, 25, 0, 0, 0, 0, time.Local),
	})
	fmt.Println(person)
	fmt.Println(person.CurrentAge())

	person.SetName("John X")

	fmt.Println("================== // ==================")

	car := model.Car{
		Automovel: model.Automovel{
			Model: "Ret",
			Year:  2022,
		},
		Doors: 4,
		Power: 150,
	}

	fmt.Println(car)
}
