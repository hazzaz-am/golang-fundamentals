package main

import (
	"errors"
	"fmt"
)

func concat(s1, s2 string) string {
	return s1 + s2
}

func getPoint() (int, int) {
	return 3, 4
}

func getNames() (string, string) {
	return "Hazzaz", "Abdul Mannan"
}

func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("can't divide by Zero")
	}
	return dividend / divisor, nil
}

func main() {
	fmt.Println(concat("Hello ", "World"))
	x, _ := getPoint()
	_, y := getPoint()
	fmt.Println(x)
	fmt.Println(y)
	firstName, lastName := getNames()
	fmt.Println(firstName, lastName)
	fmt.Println(divide(10, 0))
}
