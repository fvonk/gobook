// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 33.
//!+

// Echo4 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", "_", "separator")

func main() {
	flag.Parse()
	fmt.Println("sep value is", *sep)
	fmt.Print(strings.Join(flag.Args(), *sep))
	fmt.Printf("\n%s\n%t\n", *sep, *n)
	if !*n {
		fmt.Println()
	}
}

//!-
// go run main.go --n --s /  a bc def
