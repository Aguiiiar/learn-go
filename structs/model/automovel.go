package model

type Automovel struct {
	Model string
	Year  int
}

type Motorcycle struct {
	Automovel
	Cylinder int
}

type Car struct {
	Automovel
	Doors int
	Power int
}
