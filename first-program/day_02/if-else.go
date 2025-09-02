package main

import "fmt"

func ifElse() {
	a := 15

	// positive negative

	if a < 0 {
		fmt.Println("a is negative")
	} else if a > 0 {
		fmt.Println("a is positive")
	} else {
		fmt.Println("a is zero")
	}

	// even or odd
	if a%2 == 0 {
		fmt.Println("a is even")
	} else {
		fmt.Println("a is odd")
	}

	// age checker
	age := 12

	if age >= 0 && age <= 12 {
		fmt.Println("You are child")
	} else if age >= 13 && age <= 19 {
		fmt.Println("You are teenager")
	} else if age >= 20 {
		fmt.Println("You are adult")
	}

	// leap years
	year := 1900

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("leap year")
	} else {
		fmt.Println("no a leap year")
	}

	// letter grade
	marks := 89

	if marks >= 90 && marks <= 100 {
		fmt.Println("A")
	} else if marks >= 80 && marks <= 89 {
		fmt.Println("B")
	} else if marks >= 70 && marks <= 79 {
		fmt.Println("C")
	} else if marks >= 60 && marks <= 69 {
		fmt.Println("D")
	} else if marks < 60 {
		fmt.Println("F")
	}

	// valid number
	number := 89

	if number >= 50 && number <= 100 && number%5 == 0 {
		fmt.Println("Valid number")
	} else {
		fmt.Println("Invalid number")
	}

	// shape detect
	sides1 := 12
	sides2 := 12
	sides3 := 12

	if sides1 == sides2 && sides1 == sides3 {
		fmt.Println("Equilateral")
	} else if sides1 == sides2 || sides1 == sides3 || sides2 == sides3 {
		fmt.Println("Isosceles")
	} else {
		fmt.Println("Scalene")
	}

	num := 45

	if num%3 == 0 && num%7 == 0 {
		fmt.Println("Divisible by 3 and 7")
	} else if num%3 == 0 {
		fmt.Println("Divisible by 3")
	} else if num%7 == 0 {
		fmt.Println("Divisible by 7")
	} else {
		fmt.Println("Not Divisible by 3 or 7")
	}

	membership := true
	discount := 0

	if age < 18 {
		discount += 10
	} else if age >= 60 {
		discount += 20
	}

	if membership {
		discount += 5
	}

	fmt.Printf("Total discont: %d%%\n", discount)

	// temperature classification
	temp := 4

	if temp < 0 {
		fmt.Println("Freezing")
	} else if temp >= 0 && temp < 10 {
		fmt.Println("Very Cold")
	} else if temp >= 10 && temp < 20 {
		fmt.Println("Cold")
	} else if temp >= 20 && temp < 30 {
		fmt.Println("Warm")
	} else if temp >= 30 {
		fmt.Println("Hot")
	}

	day := "sunday"
	student := true
	price := 100

	if age >= 0 && age <= 12 {
		price = 5
	} else if age >= 13 && age <= 18 {
		price = 7
	} else if age >= 19 && age <= 59 {
		price = 10
	} else if age >= 60 {
		price = 6
	}

	if day == "Wednesday" {
		price -= 2
	}

	if student == true {
		price -= 1
	}

	fmt.Println("Ticket price", price)
}
