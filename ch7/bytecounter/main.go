// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"bufio"
	"strings"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

//!-bytecounter

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	str := string(p)
	fmt.Println(str)
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	*c += WordCounter(count)
	return len(p), nil
}

type LineCounter int

func (c *LineCounter ) Write(p []byte) (int, error) {
	str := string(p)
	fmt.Println(str)
	scanner := bufio.NewScanner(strings.NewReader(str))
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	*c += LineCounter(count)
	return len(p), nil
}

func main() {
	//!+main
	var c LineCounter
	c.Write([]byte("asd"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hel \nlo lo, \n%s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main
}
