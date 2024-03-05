// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 91.

//!+nonempty

// Nonempty is an example of an in-place slice algorithm.
package main

import (
	"fmt"
	"sort"
)

func main() {
	var data map[string]string
	fmt.Printf("%v\n", data)
	person := map[string]string{
		"name":    "alice",
		"address": "sokaraja",
	}
	fmt.Printf("%v\n", person)

	ages()
	sortingAges()
	zeroMap()
	// True if equal is written incorrectly.
	isEqual := equalMap(map[string]int{"A": 0}, map[string]int{"B": 42})
	fmt.Printf("is map equal: %t\n", isEqual)

}

func equalMap(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func zeroMap() {
	// The zero value for a map typ e is nil, that is, a reference to no hash table at all.

	var ages map[string]int
	fmt.Println(ages == nil)    // "true"
	fmt.Println(len(ages) == 0) // "true"

	// storing to nil map cause panic.
	// ages["carol"] = 21 // panic: assignment to entry in nil map

	age, ok := ages["bob"]
	if !ok { /* "bob" is not a key in this map; age == 0. */
		fmt.Println("\"bob\" is not a key in this map; age == 0.")
	}
	fmt.Printf("\"bob\"  age == %d.", age)

	if age, ok := ages["bob"]; !ok {
		fmt.Println("\"bob\" is not a key in this map; age == 0.")
	} else {
		fmt.Printf("\"bob\"  age == %d.", age)
	}

}

func sortingAges() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}

func ages() {
	ages := make(map[string]int) // mapping from strings to ints
	fmt.Printf("%v\n", ages)
	ages["alice"] = 32
	fmt.Println(ages["alice"]) // "32"

	ages["bob"] = ages["bob"] + 1 // happy birthday!
	fmt.Printf("%v\n", ages)

	ages["bob"] += 1
	fmt.Printf("%v\n", ages)

	ages["bob"]++
	fmt.Printf("%v\n", ages)

	// But a map element is not a var iable, and we cannot t ake its address:
	// _ = &ages["bob"] // compile error: cannot take address of map element

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	aAges := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Printf("%v\n", aAges)

	mage := make(map[string]int)
	mage["alice"] = 31
	mage["charlie"] = 34

	fmt.Printf("%v\n", mage)

	delete(mage, "alice") // remove element mage["alice"]
	fmt.Printf("%v\n", mage)

}
