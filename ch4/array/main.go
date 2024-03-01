// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 88.

// Append illustrates the behavior of the built-in append function.
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	arrayElement()
	arrayLiteral()
	sizeOfArray()
	arrayFormConst()
	hundredArray()
	comparingArray()
	sha256Array()
	zeroing()
}

func zeroing() {
	c1 := sha256.Sum256([]byte("x"))
	zero(&c1)
	fmt.Printf("%x\n", c1) // 0000 to the end.
	c2 := sha256.Sum256([]byte("X"))
	ezzero(&c2)
	fmt.Printf("%x\n", c2) // 0000 to the end.
}

func ezzero(ptr *[32]byte) { // type iki sof.
	*ptr = [32]byte{}
}

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}

}

func sha256Array() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8s
}

func comparingArray() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	// d := [3]int{1, 2} // declared but not used.
	// fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}

func hundredArray() {
	// defines an array r with 100 elements, all zero except for the last, which has value −1.
	r := [...]int{99: -1}
	fmt.Println(r) // 3 ¥

}

func arrayFormConst() {
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB]) // 3 ¥
}

func sizeOfArray() { // is type
	q := [3]int{1, 2, 3}
	// q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int

	fmt.Printf("%T\n", q) // "[3]int"

}

func arrayLiteral() {
	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2} // element yang kosong default zero.
	fmt.Println(r[2], q[2])     // 0 3

	// type determined by initialization

	s := [...]int{1, 2, 3}
	fmt.Printf("%T\n", s) // "[3]int"
}

func arrayElement() {
	var a [5]int = [5]int{1, 2, 3, 4, 5}

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("index: %d, elemt: %d\n", i, v)
	}
	// Print the elements only.
	for _, v := range a {
		fmt.Printf("element: %d\n", v)
	}
}
