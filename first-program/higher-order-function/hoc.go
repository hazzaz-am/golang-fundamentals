package main

import "fmt"

func processOperation(a, b int, op func(x, y int)) {
	op(a, b)
}

func add(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	processOperation(3, 4, add)
}
