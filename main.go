package main

import (
	"fmt"
	"math"
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
	declared()
	fmt.Println("_________________________________________________________")
	// basic types in go
	basicTypes()
	fmt.Println("_________________________________________________________")
	// zero values declaration
	declaresZeroValues()
	fmt.Println("_________________________________________________________")

	// type convertions
	convertType()
	fmt.Println("_________________________________________________________")

	// Contants
	declaredContants()
	fmt.Println("_________________________________________________________")

	// for
	forLoop()
	fmt.Println("_________________________________________________________")

	// if
	fmt.Println(sqrt(2), sqrt(-4))
	fmt.Println(pow(2, 5, 10), pow(1, 2, 10))
	fmt.Println("_________________________________________________________")

	// Exercise: Loops and Functions
	var x float64 = 100
	fmt.Printf("~ Result: %v\n", exerciseSqrt(x))
	fmt.Printf("math.Sqrt 81: %v\n", math.Sqrt(x))
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

func declared() {
	// Initialize variable
	var x, y int = 1, 2

	fmt.Println("x:", x, "| y:", y, "| c:", c, "| python:", python, "| java:", java, "| go:", golang)

	// Short variable declaration (:=) in function
	z, h := "Short variable declaration", "no!"

	fmt.Println("z:", z, "| h:", h)
}

// Basic types
func basicTypes() {
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

// Zero values
func declaresZeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

// Type conversions
func convertType() {
	var i, j int = 10, 11
	var f float64 = math.Sqrt(float64(i*i + j*j))
	var z uint = uint(f)
	// or using this syntax
	// i, j := 10, 11
	// f := math.Sqrt(float64(i*i + j*j))
	// z := uint(f)
	fmt.Printf("Type: %T Value: %v\n", i, i)
	fmt.Printf("Type: %T Value: %v\n", j, j)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// Constants
func declaredContants() {
	const name string = "Huy Van"
	// can not use syntax (:=) when declares const
	// Constants can be character, string, boolean, or numeric values.

	fmt.Printf("Type: %T Value: %v\n", name, name)
}

// for in Go
func forLoop() {
	// For
	sum := 0
	for i := 0; i < 100; i++ {
		sum += 1
	}
	fmt.Println("Sum: ", sum)
	// For continued
	// the init and post statment can be optional
	// if the condition expression is omit, it will be loop forever
	counter := 0
	for counter < 10 {
		counter += 1
	}
	fmt.Println("Counter: ", counter)
}

// if in Go
// if and else
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
func pow(x, n, lim float64) float64 {
	// v only use in if scope
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		return lim
	}
}

// Exercise: Loops and Functions
// https://go.dev/tour/flowcontrol/8

func exerciseSqrt(x float64) float64 {
	z := float64(1)
	for math.Abs((z*z-x)/(2*z)) > 1e-6 {
		fmt.Println(math.Abs((z*z-x)/(2*z)), 1e-6)
		z -= (z*z - x) / (2 * z)
		fmt.Printf("z = %v\n", z)
	}
	return z
}
