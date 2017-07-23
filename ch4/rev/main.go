// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//!+array
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a) // "[5 4 3 2 1 0]"
	// reverse(&a, 6)
	rotate(a, -1)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	fmt.Printf("%T\n", a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	// reverse(s[:2])
	// reverse(s[2:])
	// reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		// reverse(ints)
		rotate(ints, -12)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!+rev
// reverse reverses a slice of ints in place.
func rotate(a []int, shift int) {
	s := shift % len(a)
	if shift < 0 {
		s = len(a) + s
	}

	b := make([]int, len(a), cap(a))
	copy(b, a)
	j := 0
	for i := s; i < len(a) + s; i++ {
		pos := i
		if i >= len(a) {
			pos = i - len(a)
		}
		a[pos] = b[j]
		j++
	}
}
// func reverse(s []int) {
// 	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
// 		s[i], s[j] = s[j], s[i]
// 	}
// }
func reverse(s *[32]int, l int) {
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev
