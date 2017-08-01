// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import "fmt"

//!+
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(val int, vals ...int) int {
	max := val
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}
func min(val int, vals ...int) int {
	min := val
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

//!-

func main() {
	//!+main
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
	fmt.Println(max(1, 7, 3, 4)) 
	fmt.Println(max(1, 4, -4)) 
	fmt.Println(min(1, 7, 3, 4)) 
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	//!-slice
}
