package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum
}

func twoNumbers(num1 int, num2 int) (int, int) {
	sum := num1 + num2
	multiply := num1 * num2

	return sum, multiply
}

func getName(name string) string {
	return fmt.Sprintf("My name is %s", name)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("%.2f cannot divide by zero", a)
	}
	return a / b, nil
}

func printWellcomeMessage() string {
	return "Wellcome to Goodbye Application"
}

func printYourName() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

func inputNumbers() (float64, float64) {
	var num1 float64
	var num2 float64
	fmt.Print("Enter first number: ")
	fmt.Scanln(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&num2)

	return num1, num2
}

func showResult(name string, num1, num2 float64) {
	fmt.Printf("%s, this is your result: %.2f\n", name, num1+num2)
}

func goodByeMessage() {
	fmt.Println("Good Bye")
}

func fn() {
	printWellcomeMessage()
	name := printYourName()
	a, b := inputNumbers()
	showResult(name, a, b)
	goodByeMessage()

	sum := add(10, 20)
	x, y := twoNumbers(10, 20)
	getName := getName("Hazzaz")
	result, error := divide(10, 0)

	if error != nil {
		fmt.Println("ERROR:", error)
	} else {
		fmt.Println(result)
	}

	fmt.Println(getName)
	fmt.Println(sum, x, y)
}
