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
	// basic types in go
	basicTypes()
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

func basicTypes() {
	// Basic types
	var (
		// Boolean
		ToBe bool = false
		// String
		String string = "Hello Go"
		// Numeric
		Int      int        = 42                   // size(bits) 32/64, min -2,147,483,648 / -9,223,372,036,854,775,808, max 2,147,483,647 / 9,223,372,036,854,775,807
		Int8     int8       = 127                  // size(bits) 8, min -128, max 127
		Int16    int16      = 32767                // size(bits) 16, min -32,768, max 32,767
		Int32    int32      = 2147483647           // size(bits) 32, min -2,147,483,648, max 2,147,483,647
		Int64    int64      = 9223372036854775807  // size(bits) 64, min -9,223,372,036,854,775,808, max 9,223,372,036,854,775,807
		Unit     uint       = 42                   // size(bits) 32/64, min 0, max 4,294,967,295 / 18,446,744,073,709,551,615
		Unit8    uint8      = 255                  // size(bits) 8, min 0, max 255
		Unit16   uint16     = 65535                // size(bits) 16, min 0, max 65,535
		Unit32   uint32     = 4294967295           // size(bits) 32, min 0, max 4,294,967,295
		Unit64   uint64     = 18446744073709551615 // size(bits) 64, min 0, max 18,446,744,073,709,551,615
		Float32  float32    = 3.14                 // size(bits) 32, min 1.18E-38, max 3.4E+38
		Float64  float64    = 3.14                 // size(bits) 64, min 2.23E-308, max 1.8E+308
		Cmplx64  complex64  = 3.14 + 12i           // size(bits) 64, min 1.18E-38, max 3.4E+38
		Cmplx128 complex128 = 3.14 + 12i           // size(bits) 128, min 2.23E-308, max 1.8E+308
		Byte     byte       = 255                  // size(bits) 8, min 0, max 255
		Rune     rune       = 2147483647           // size(bits) 32, min -2,147,483,648, max 2,147,483,647
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", String, String)
	fmt.Printf("Type: %T Value: %v\n", Int, Int)
	fmt.Printf("Type: %T Value: %v\n", Int8, Int8)
	fmt.Printf("Type: %T Value: %v\n", Int16, Int16)
	fmt.Printf("Type: %T Value: %v\n", Int32, Int32)
	fmt.Printf("Type: %T Value: %v\n", Int64, Int64)
	fmt.Printf("Type: %T Value: %v\n", Unit, Unit)
	fmt.Printf("Type: %T Value: %v\n", Unit8, Unit8)
	fmt.Printf("Type: %T Value: %v\n", Unit16, Unit16)
	fmt.Printf("Type: %T Value: %v\n", Unit32, Unit32)
	fmt.Printf("Type: %T Value: %v\n", Unit64, Unit64)
	fmt.Printf("Type: %T Value: %v\n", Float32, Float32)
	fmt.Printf("Type: %T Value: %v\n", Float64, Float64)
	fmt.Printf("Type: %T Value: %v\n", Cmplx64, Cmplx64)
	fmt.Printf("Type: %T Value: %v\n", Cmplx128, Cmplx128)
	fmt.Printf("Type: %T Value: %v\n", Byte, Byte)
	fmt.Printf("Type: %T Value: %v\n", Rune, Rune)
}
