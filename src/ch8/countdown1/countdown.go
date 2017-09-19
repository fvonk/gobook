// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 244.

// Countdown implements the countdown for a rocket launch.
package main

import (
	"fmt"
	//"time"
)

//!+
func main() {
	//fmt.Println("Commencing countdown.")
	//tick := time.Tick(1 * time.Second)
	//for countdown := 10; countdown > 0; countdown-- {
	//	fmt.Println(countdown)
	//	fmt.Println(<-tick)
	//}
	//launch()

	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		fmt.Println("i", i)
		select {
		case x := <- ch:
			fmt.Println("case 1", i, x)
		case ch <-i:
			fmt.Println("case 2", i)
		}
	}
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
