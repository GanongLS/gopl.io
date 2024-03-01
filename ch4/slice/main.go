// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 88.

// Append illustrates the behavior of the built-in append function.
package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
	// "strings"
)

func main() {
	arrayOfMonths()
	reverseInt()
	testEqual()
}

func testEqual() {
	hello := []string{"hello", "world"} // init
	name := []string{"hello", "world"}
	isEqual := equal(hello, name)

	// checking
	fmt.Println(isEqual)
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func reverseInt() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice

	// Interactive test of reverse.
	// 	input := bufio.NewScanner(os.Stdin)
	// outer:
	// 	for input.Scan() {
	// 		var ints []int
	// 		for _, s := range strings.Fields(input.Text()) {
	// 			x, err := strconv.ParseInt(s, 10, 64)
	// 			if err != nil {
	// 				fmt.Fprintln(os.Stderr, err)
	// 				continue outer
	// 			}
	// 			ints = append(ints, int(x))
	// 		}
	// 		reverse(ints)
	// 		fmt.Printf("%v\n", ints)
	// 	}
	// NOTE: ignoring potential errors from input.Err()
}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func arrayOfMonths() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Printf("%v\n%x\n", months, months)

	Q2 := months[4:7]     // notasi slice, slice itu dynamically sized slice of array.
	summer := months[6:9] // notasi slice
	fmt.Println(Q2)       // ["April" "May" "June"]
	fmt.Println(summer)   // ["June" "July" "August"]

	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}

	// Slicing beyond cap(s) causes a panic, but slicing beyond len(s) extends the slice, so the result may be longer than the original:

	// fmt.Println(summer[:20])    // panic: out of range
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"
}
