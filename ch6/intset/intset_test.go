// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "fmt"
// import "testing"
// import "gobook/ch6/intset"
//go test -v
func main() {

}

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	y.Add(43)
	y.Add(44)
	y.Add(444)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123), x.Has(444)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42 43 44 444}
	// {1 9 42 43 44 144 444}
	// true false true
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

func Example_three() {
	var x IntSet
	x.Add(1)
	fmt.Println(x.Len())
	x.Add(2)
	fmt.Println(x.Len())
	x.Add(1)
	fmt.Println(x.Len())

	x.Remove(3)
	fmt.Println(x.String())

	x.Remove(2)
	fmt.Println(x.String())

	x.Remove(1)
	fmt.Println(x.String())

	x.Add(1)
	x.Add(10)
	x.Add(4)
	x.Clear()
	fmt.Println(x.String())

	x.Add(1)
	x.Add(10)
	x.Add(4)
	y := x.Copy()
	fmt.Println(y.String())

	// Output:
	// 1
	// 2
	// 2
	// {1 2}
	// {1}
	// {}
	// {}
	// {1 4 10}
}
