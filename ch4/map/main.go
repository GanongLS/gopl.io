// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

func main() {
	var data map[string]string
	fmt.Printf("%v\n", data)
	person := map[string]string{
		"name":    "alice",
		"address": "sokaraja",
	}
	fmt.Printf("%v\n", person)

	ages()
}

func ages() {
	ages := make(map[string]int) // mapping from strings to ints
	fmt.Printf("%v\n", ages)
	ages["alice"] = 32
	fmt.Println(ages["alice"]) // "32"

	aAges := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Printf("%v\n", aAges)

	mage := make(map[string]int)
	mage["alice"] = 31
	mage["charlie"] = 34

	fmt.Printf("%v\n", mage)
}
