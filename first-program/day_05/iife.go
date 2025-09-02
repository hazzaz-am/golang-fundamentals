package main

import "fmt"

var subtract = func(a, b int) {
	fmt.Println(a - b)
}

func main() {
	func(a, b int) {
		fmt.Println(a + b)
	}(5, 4)

	add := func(a, b int) {
		fmt.Println(a + b)
	}
	add(20, 20)
	subtract(50, 20)
}

func init() {
	fmt.Println("I am the first")
}
