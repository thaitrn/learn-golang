package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"strings"
	"time"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
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

	// Switch in Go
	swith()
	fmt.Println("_________________________________________________________")

	// Defer in Go
	Defer()
	fmt.Println("_________________________________________________________")

	// Pointers in Go
	pointers()
	fmt.Println("_________________________________________________________")

	// structs in Go
	structs()
	fmt.Println("_________________________________________________________")

	// Arrays in Go
	arr()
	fmt.Println("_________________________________________________________")

	// Range in Go
	rangeInGo()
	fmt.Println("_________________________________________________________")

	// Exercise: Slices
	pic.Show(Pic)
	fmt.Println("_________________________________________________________")

	// Maps in Go
	mapInGo()
	fmt.Println("_________________________________________________________")

	// Exercise: Maps
	wc.Test(WordCount)
	fmt.Println("_________________________________________________________")

	// func value in Go
	fmt.Println("Compute sum of 3 and 4: ", compute(sum))
	fmt.Println("_________________________________________________________")

	// clourse in Go
	pos := adder()
	fmt.Println("Function closures: ", pos(1))
	fmt.Println("_________________________________________________________")
	// Exercise: Fibonacci closure

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Print(f(), ", ")
	}
	fmt.Println()
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

// swith in Go
func swith() {
	// switch
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	// switch evaluation order
	fmt.Print("When's Saturday? ")
	switch today := time.Now().Weekday(); time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	// Switch with no condition
	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// Defer
func Defer() {
	// defer
	defer fmt.Println("Defer in Go")
	for i := 0; i < 3; i++ {
		// Stacking defers
		defer fmt.Println(i + 1)
	}
	fmt.Println("I try to use ")

}

// Key Concepts of Pointers in Go

type Person struct {
	Name string
	Age  int
}

func increment(num *int) {
	*num += 1
}

func pointers() {
	// basic
	var a int = 42
	var p *int = &a // p is a pointer to an int
	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of p (address):", p)
	fmt.Println("Dereferenced p:", *p) // read a through the pointer p

	// Pointer Operations
	*p = 42 / 2
	fmt.Println("Value of a after *p = 42 / 2:", a) // set a through the pointer p

	// Nil Pointers
	var n *int
	fmt.Println("Nil pointer:", n)

	// Pointer to Struct
	s := &Person{"Huy Van", 28}
	fmt.Println(s.Age, s.Name) // Implicitly dereferences the pointer
	s.Age = 31                 // Modifies the value through the pointer
	fmt.Println(s.Age, s.Name) // Implicitly dereferences the pointer

	// Functions and Pointers
	x := 10
	increment(&x)
	fmt.Println("x after pass adddress to func increment: ", x) // x is now 11
}

// Structs in Go

type User struct {
	Name  string
	Email string
}

func structs() {
	u := User{"Huy Van", "dvhuy.dev@gmail.com"}
	fmt.Println("User structs: ", u)

	// Struct Fields
	fmt.Println("User Name: ", u.Name)

	// pointers to structs
	p := &u
	p.Name = "Van Huy"
	fmt.Println("u after modifies User Name using poiter p: ", u.Name)

	// struct Literals
	var (
		u1 = User{Name: "Huy Van", Email: "dvhuy.dev@gmail.com"}
		u2 = User{Email: "thaitran@gmail.com"}
		u3 = User{}

		p1 = &User{Name: "Huy Van", Email: "dvhuy.dev@gmail.com"}
	)

	fmt.Println("User1: ", u1)
	fmt.Println("User2: ", u2)
	fmt.Println("User3: ", u3)
	fmt.Println("Poiter user: ", p1)

}

// Arrays
func arr() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println("Arrays: ", a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Primes: ", primes)

	// slices
	var s []int = primes[1:4]
	fmt.Println("Slices: ", s)
	s[1] = 100
	fmt.Println("Primes after modifies slices: ", primes)

	// slice literals
	s1 := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Slice literals: ", s1)
	s2 := []bool{true, false, true, true, false, true}
	fmt.Println("Slice literals: ", s2)
	s3 := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("Slice literals: ", s3)
	fmt.Println("Slice literals with default hight: ", s3[:4])
	fmt.Println("Slice literals with default low: ", s3[4:])

	// slice length and capacity
	s1 = s1[:3]
	fmt.Printf("len=%d cap=%d %v\n", len(s1), cap(s1), s1)

	// nil slices
	var s4 []int
	fmt.Println("Nil slices: ", s4, len(s4), cap(s4))

	// Creating a slice with make
	s5 := make([]int, 5)
	fmt.Println("Creating a slice with make([]int, 5): ", s5)
	s5 = make([]int, 1, 5)
	fmt.Println("Creating a slice with make([]int, 1, 5): ", s5)

	// slices of slices
	// Create a tic-tac-toe board.
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "X"
	board[1][1] = "X"
	board[1][0] = "O"
	board[0][2] = "O"
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	// Appending to a slice
	fmt.Printf("Before Appending to a slice s1 len=%d cap=%d %v\n", len(s1), cap(s1), s1)
	s1 = append(s1, 2)
	fmt.Printf("Appending  2 to a slice s1 len=%d cap=%d %v\n", len(s1), cap(s1), s1)
}

// Range
func rangeInGo() {
	s := []int{10, 20, 30, 40, 50}
	for i, v := range s {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}
	for _, v := range s {
		fmt.Printf("Value: %d\n", v)
	}
	for i := range s {
		fmt.Printf("Index: %d\n", i)
	}
}

// Exercise: Slices
// [][]uint8
func Pic(dx, dy int) [][]uint8 {
	pic := make([][]uint8, dy)
	for y := range pic {
		pic[y] = make([]uint8, dx)
		for x := range pic[y] {
			pic[y][x] = uint8((x * y) / 2)
		}
	}
	return pic
}

// Maps

func mapInGo() {
	m := map[string]int{
		"Huy":   28,
		"Khang": 35,
	}

	fmt.Println(m, "maps")

	// Mutating Maps
	m["Thai"] = 33
	fmt.Println(m, "maps after add Thai")

	delete(m, "Khang")
	fmt.Println(m, "maps after delete Khang")

	v, ok := m["Iconito"]
	fmt.Println("The value:", v, "Present?", ok)
}

// Exercise: Maps
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}
	fmt.Println("WordCount: ", counts)

	return counts
}

// func value in Go
func compute(sumFnc func(int, int) int) int {
	return sumFnc(3, 4)
}

// clourse in Go
func adder() func(int) int {
	sum := 0
	fn := func(x int) int {
		sum += x
		return sum
	}
	return fn
}

// Exercise: Fibonacci closure
func fibonacci() func() int {
	a := 0
	b := 1
	fn := func() int {
		next := a
		a, b = b, next+b
		return next
	}
	return fn
}
