package main

import (
	"custom/mathlib"
	"fmt"
)

func scope() {
	fmt.Println("Showing package scope")
	mathlib.Add(12, 21)
}