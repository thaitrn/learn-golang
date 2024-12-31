package main

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"math"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"time"

	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
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
	// Method
	u := User{"Huy Van", "dvhuy.dev@gmail.com"}
	fmt.Println(u.String())
	fmt.Println("_________________________________________________________")

	// poiter receivers
	u.updateEmail("huy.dinh@navigosgroup.com")
	fmt.Println(u.String())
	fmt.Println("_________________________________________________________")
	// Methods and pointer indirection
	v := Vertex{3, 4}
	v.Scale(2)

	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
	fmt.Println("_________________________________________________________")

	// Methods and pointer indirection (2)
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
	fmt.Println("_________________________________________________________")

	// interfaces
	var i I
	var t *T
	i = t
	i.M()
	describe(i)
	i = &T{"hello"}
	describe(i)
	var j EmptyInterface
	describeEmpty(j)

	fmt.Println("_________________________________________________________")

	// Exercise: Stringers
	stringers()
	fmt.Println("_________________________________________________________")

	// Exercise: Errors
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-4))
	fmt.Println("_________________________________________________________")

	// Exercise: Stringers
	stringers()
	fmt.Println("_________________________________________________________")

	// Exercise: Errors
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(-4))
	fmt.Println("_________________________________________________________")

	// Exercise: Reader
	reader.Validate(MyReader{})
	fmt.Println("_________________________________________________________")
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Println("_________________________________________________________")

	// Exercise: Images

	m := Image{256, 256}
	pic.ShowImage(m)
	fmt.Println("_________________________________________________________")

	// type parameters
	printType(10)
	printType("10")
	printType(true)

	// generic type
	list := LinkedList[int]{}
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Print() // Output: 1 -> 2 -> 3 -> nil
	fmt.Println("_________________________________________________________")

	// Channels
	sChannels := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go Channels(sChannels[:len(sChannels)/2], c)
	go Channels(sChannels[len(sChannels)/2:], c)
	z, k := <-c, <-c // receive from c
	fmt.Println(z, k, z+k)
	d := make(chan int, 10)
	go fibonacci2(cap(d), d)
	for i := range d {
		fmt.Println(i, "á")
	}
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

// Method
func (u User) String() string {
	return fmt.Sprintf("User: %v, %v", u.Name, u.Email)
}

// poiter receivers

func (u *User) updateEmail(email string) {
	u.Email = email
}

// Methods and pointer indirection

type Vertex struct {
	X, Y float64
}

// Scale is a method that scale a vertex
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// ScaleFunc is a function that scale a vertex
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// Abs is a method that return the absolute value of a vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// AbsFunc is a function that return the absolute value of a vertex
func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Interfaces

type I interface {
	M()
}

type EmptyInterface interface{}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func describe(i I) {
	// tyoe assertions
	v, ok := i.(*T)
	fmt.Println(v, ok, "type assertions")
	fmt.Printf("(%v, %T)\n", i, i)
}

func describeEmpty(i EmptyInterface) {
	// tyoe assertions
	v, ok := i.(*T)
	fmt.Println(v, ok, "type assertions")
	fmt.Printf("(%v, %T)\n", i, i)
}

// Exercise: Stringers
type IPAddr [4]byte

func (ip IPAddr) String() string {
	// Format the address as a dotted quad
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func stringers() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)

	}
}

// Exercise: Errors
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		// Return an instance of ErrNegativeSqrt as an error
		return 0, ErrNegativeSqrt(x)
	}

	// Simple approximation of the square root
	z := x / 2
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

// Exercise: Readers

type MyReader struct{}

// Read implements io.Reader.
func (r MyReader) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = 'A'
	}
	return len(p), nil
}

// Exercise: rot13Reader
type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)

	for i := range p {
		p[i] = rot13(p[i])
	}

	return n, err
}

func rot13(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		print('A'+(b-'A'+13)%26, " ")
		return 'A' + (b-'A'+13)%26
	} else if b >= 'a' && b <= 'z' {
		print('a'+(b-'a'+13)%26, " ")
		return 'a' + (b-'a'+13)%26
	}

	return b
}

// Exercise: Images
type Image struct {
	width, height int
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x ^ y) % 256)
	return color.RGBA{v, v, 255, 255}
}

// type parameter
func printType[T comparable](t T) {
	fmt.Printf("(%v, %T)\n", t, t)
}

// Generic type
// Define a generic ListNode with a type parameter
type ListNode[T any] struct {
	value T
	next  *ListNode[T]
}

// Define a generic LinkedList with a type parameter
type LinkedList[T any] struct {
	head *ListNode[T]
}

// Add a method to insert a value at the end of the list
func (l *LinkedList[T]) Insert(value T) {
	newNode := &ListNode[T]{value: value}
	if l.head == nil {
		l.head = newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	fmt.Print(current)
}

// Add a method to print the values in the list
func (l *LinkedList[T]) Print() {
	current := l.head
	for current != nil {
		fmt.Printf("%v -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

// goroutines and channels
func Channels(s []int, c chan int) {
	fmt.Println(s, c)

	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func fibonacci2(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
