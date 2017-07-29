// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 12.

//!+

// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "expand: %v\n", err)
			continue
		}
		res := "\n-----------\n"
		res += expand(string(data), func (s string) string {
			if s == "$foo" {
				return "baa"
			} else {
				return s
			}
		})
		f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Fprintf(os.Stderr, "expand: %v\n", err)
			continue
		}
		if _, err := f.WriteString(res); err != nil {
			f.Close()
			fmt.Fprintf(os.Stderr, "expand: %v\n", err)
			continue
		}
		f.Close()
	}
}

//"$foo"
func expand(s string, f func(string) string) string {
	r := strings.NewReplacer("$foo", f("$foo"))
	return r.Replace(s)
}

//!-
