package main

import "fmt"

type Car struct {
	Make       string
	Model      string
	Height     int
	Width      int
	FrontWheel Wheel
	BackWheel  Wheel
}

type Wheel struct {
	Radius   int
	Material string
}

type car struct {
	make  string
	model string
}

type truck struct {
	car
}

type rect struct {
	Width  int
	Height int
}

func (r rect) area() int {
	return r.Height * r.Width
}

func main() {
	r := rect{
		Width:  4,
		Height: 5,
	}
	fmt.Println(r.area())
	myCar := Car{}
	myCar.Model = "BMW"
	myCar.Make = "Germany"
	myCar.Height = 176
	myCar.Width = 456
	myCar.BackWheel.Radius = 5
	myCar.FrontWheel.Radius = 5
	myCar.FrontWheel.Material = "Plastic"
	myCar.BackWheel.Material = "Plastic"
	banglaTruck := truck{}
	banglaTruck.make = "Gulistan"
	fmt.Println(banglaTruck)

}
