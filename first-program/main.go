package main

import "fmt"

func main() {
	// go is statically typed (types are fixed once declared)
	var age int = 20         // explicit type
	name := "Hazzaz"         // type inferred
	var price float32 = 20.5 // floating type
	var isCool bool = true

	// multiple variables
	var x, y, z int = 1, 2, 3
	name1, age1 := "Go", 21
	var (
		name2 string = "Javascript"
		age2  int    = 20
		cool  bool   = true
	)

	var zero int
	fmt.Println((zero))

	var ascii uint8 = 'A'
	var unicode int32 = 'h'
	fmt.Println(ascii)
	fmt.Println(unicode)

	// var ub uint = -10 // uint type can hole only no -negative number

	

	fmt.Println(name2, age2, cool)
	fmt.Println(name1, age1)
	fmt.Println(x, y, z)
	fmt.Println(isCool)
	fmt.Println(price)
	fmt.Println(age)
	fmt.Println(name)
}
