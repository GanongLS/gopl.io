// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
	stringConversion()
	parseInteger()
	constList()
	printIota()
	benderaBerkibar()
	printIlegibleConstant()
	untypedConstant()
	untypedConstant1()
	untypedConstant2()
	constOverflow()
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func stringConversion() {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))             // "123 123"
	fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
	s := fmt.Sprintf("x=%b", x)                 // "x=1111011"
	fmt.Println(s)                              // "x=1111011"
}

func parseInteger() {
	x, err := strconv.Atoi("123")
	fmt.Println(x, err)                       // x is an int
	y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	fmt.Println(y, err)
}

//!-

func constList() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // "1 1 2 2"
}

type Weekday int

const (
	Sunday    Weekday = iota // 0
	Monday                   // 1
	Tuesday                  // 2
	Wednesday                // 3
	Thursday                 // 4
	Friday                   // 5
	Saturday                 // 6
)

func printIota() {
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)
}

type Flags uint

const (
	FlagUp Flags = 1 << iota // is up
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func benderaBerkibar() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 							(exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 		(exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

func printIlegibleConstant() {
	fmt.Println(YiB / ZiB) // "1024"
}

func untypedConstant() {
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Printf("%f\n", x)
	fmt.Printf("%f\n", y)
	fmt.Printf("%f\n", z)
}

func untypedConstant1() {
	const Pi64 float64 = math.Pi
	var x float32 = float32(Pi64)
	var y float64 = Pi64
	var z complex128 = complex128(Pi64)
	fmt.Printf("%f\n", x)
	fmt.Printf("%f\n", y)
	fmt.Printf("%f\n", z)
}

func untypedConstant2() {
	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9) // "100"; (f - 32) * 5 is a float64
	// the integer division '5 / 9' results in zero (SA4025)
	fmt.Println(float64(5) / 9 * (f - 32)) //_// "0";   5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32))      // "100"; 5.0/9.0 is an untyped float

	var ff float64 = 3 + 0i // untyped complex -> float64
	ff = 2                  // untyped integer -> float64
	ff = 1e123              // untyped floating-point -> float64
	ff = 'a'                // untyped rune -> float64
	fmt.Printf("%f\n", ff)

	// The statements above are thus equivalent to these:
	// var f float64 = float64(3 + 0i)
	ff = float64(2)
	ff = float64(1e123)
	ff = float64('a')

}

func constOverflow() {
	const (
		deadbeef = 0xdeadbeef        // untyped int with value 3735928559
		a        = uint32(deadbeef)  // uint32 with value 3735928559
		b        = float32(deadbeef) // float32 with value 3735928576 (rounded up)
		c        = float64(deadbeef) // float64 with value 3735928559 (exact)
		// d        = int32(deadbeef)   // compile error: constant overflows int32
		// e        = float64(1e309)    // compile error: constant overflows float64
		// f        = uint(-1)          // compile error: constant underflows uint

	)

	fmt.Printf("%d\n", deadbeef)
	fmt.Printf("%d\n", a)
	fmt.Printf("%f\n", b)
	fmt.Printf("%f\n", c)

	i := 0      // untyped integer;        implicit int(0)
	r := '\000' // untyped rune;           implicit rune('\000')
	f := 0.0    // untyped floating-point; 	implicit float64(0.0)
	co := 0i    // untyped complex;        implicit complex128(0i)

	fmt.Printf("%d\n", i)
	fmt.Printf("%d\n", r)
	fmt.Printf("%f\n", f)
	fmt.Printf("%f\n", co)

	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)
}
