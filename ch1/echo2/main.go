// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, sep := "", ""
	for id, arg := range os.Args[1:] { // os.Args[1:] mulai dari index ke 1, param ke 0 itu alamat file.
		s += sep + arg
		sep = strconv.Itoa(id+1) + ", "
	}
	fmt.Println(s)
}

//!-
