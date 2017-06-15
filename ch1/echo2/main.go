// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	t := time.Now()
	s := ""
	for idx, arg := range os.Args[0:] {
		s += string(idx) + " " + arg + "\n"
	}
	fmt.Println(s)
	duration := time.Since(t)
	fmt.Printf("%f\n", duration.Seconds())
}

//!-
