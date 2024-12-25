package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello Word, this is my random number from 0 -> 100: ", rand.Intn(100))

	// sum
	fmt.Println("Sum of 2 numb 102 and 22: ", sum(102, 22))

	// multiple result in go fnc
	firtName, lastName := fullName("Van", "Huy")
	fmt.Println("Print my name (Van Huy) with en version: ", lastName, firtName)

}

// func in Go
func sum(x, y int) int {
	return x + y
}

// Multiple results in fnc

func fullName(firtName, lastName string) (string, string) {
	return firtName, lastName
}
