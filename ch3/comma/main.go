// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"
	"bytes"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", commaBuf(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func commaBuf(s string) string {
	var buf bytes.Buffer
	var j = 0
	var count = len(s) / 3
	var place = len(s) % 3
	for i := 0; i < len(s); i++ {
		// fmt.Printf("place = %d j = %d i = %d\n", place, j, i)
		if count > 0 && place == i  && place != 0 {
			buf.WriteString(",")
			j = 0
		} else if j >= 3 && count > 0 {
			buf.WriteString(",")
			j = 0
		}
		j++
		buf.WriteByte(s[i])
	}
	return buf.String()
}

//!-
