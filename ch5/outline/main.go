// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 123.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"os"
	// "encoding/json"

	"golang.org/x/net/html"
)

var res = map[string]int{}

//!+
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	// outline(nil, doc)
	//5.2
	// outlineMap(doc)
	// b, err := json.MarshalIndent(res, "", "  ")
	// if err != nil {
	//     fmt.Println("error:", err)
	// }
	// fmt.Println(string(b))
	findTexts(doc)
}

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func outlineMap(n *html.Node) {
	if n.Type == html.ElementNode {
		res[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outlineMap(c)
	}
}

func findTexts(n *html.Node) {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "type" && a.Val == "image/png" {
				// return
				// fmt.Printf("%v %v\n", a.Key, a.Val)
				fmt.Printf("%v\n", n.Data)
				break
			}
		}
		// fmt.Printf("%v\n", n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findTexts(c)
	}
}

//!-
