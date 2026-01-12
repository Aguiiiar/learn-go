package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func ViewAreaRect(g Geometry) float64 {
	return g.Area()
}

func ViewAreaCircle(g Geometry) float64 {
	return g.Area()
}

func main() {
	rectangle := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 7}

	fmt.Println(ViewAreaRect(rectangle))
	fmt.Println(ViewAreaCircle(circle))
}
