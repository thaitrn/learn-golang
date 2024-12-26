package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("################# Welcome to Go Lang #################")

	fmt.Println("Hello Word, this is my random number from 0 -> 100: ", rand.Intn(100))
	fmt.Println("_________________________________________________________")

	// sum
	fmt.Println("Sum of 2 numb 102 and 22: ", sum(102, 22))
	fmt.Println("_________________________________________________________")

	// multiple result in go fnc
	firtName, lastName := fullName("Van", "Huy")
	fmt.Println("Print my name (Van Huy) with en version: ", lastName, firtName)
	fmt.Println("_________________________________________________________")

	// named return values
	value := date(9, 10, 1997)
	fmt.Println("date 9/10/1997: ", value)
	fmt.Println("_________________________________________________________")

	// variable in go
	declares()
	fmt.Println("_________________________________________________________")
}

// func in Go
func sum(x, y int) int {
	return x + y
}

// Multiple results in fnc
func fullName(firtName, lastName string) (string, string) {
	return firtName, lastName
}

// Named return values
func date(d int, m int, y int) (date string) {
	date = fmt.Sprintf("%2d/%d/%d", d, m, y)
	return
}

// Variable in Go
// Initialize variable
var c, python, java, golang bool = false, false, false, true

func declares() {
	// Initialize variable
	var x, y int = 1, 2

	fmt.Println("x:", x, "| y:", y, "| c:", c, "| python:", python, "| java:", java, "| go:", golang)

	// Short variable declaration (:=) in function
	z, h := "Short variable declaration", "no!"

	fmt.Println("z:", z, "| h:", h)
}
