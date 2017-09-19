// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 246.

// Countdown implements the countdown for a rocket launch.
package main

// NOTE: the ticker goroutine never terminates if the launch is aborted.
// This is a "goroutine leak".

import (
	"fmt"
	"os"
	"time"
)

//!+
var abort = make(chan struct{})

func gofunc() {
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
}

func main() {
	// ...create abort channel...

	//!-


	gofunc()

	//!+
	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			//return
			countdown = 10
			gofunc()
		}
	}
	launch()
}

//!-

func launch() {
	fmt.Println("Lift off!")
}
