// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 12.

//!+

// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			// fmt.Println("line:" + line)
			counts[line]++
		}
	}
	for line, n := range counts {
		// fmt.Println(line + "\n")
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line+"\n")
		}
	}
	// fmt.Println(counts)
}

//!-
