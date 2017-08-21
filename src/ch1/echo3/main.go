// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
    // See http://golang.org/pkg/time/#Parse
    timeFormat = "2006-01-02 15:04 MST"
)

//!+
func main() {
	t := time.Now()
	fmt.Println(strings.Join(os.Args[0:], " "))
	duration := time.Since(t)
	fmt.Printf("%f\n", duration.Seconds())
}

//!-
