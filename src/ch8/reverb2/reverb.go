// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
	"os"
	"sync"
)

func echo(c net.Conn, shout string, delay time.Duration, done chan bool, wg sync.WaitGroup) {
	defer wg.Done()

	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	fmt.Fprintln(os.Stdout, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	fmt.Fprintln(os.Stdout, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	fmt.Fprintln(os.Stdout, "\t", strings.ToLower(shout))

	done <- true
}


//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup

	done := make(chan bool)

	for input.Scan() {

		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, done, wg)
		<-done
		fmt.Println("Done!")
	}
	fmt.Println("After Done!")
	go func() {
		fmt.Println("Wait!")
		wg.Wait()
		if err := c.(*net.TCPConn).CloseWrite(); err != nil {
			panic(1)
		}
	}()
	fmt.Println("c.Close!")
	// NOTE: ignoring potential errors from input.Err()
	c.Close()
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
