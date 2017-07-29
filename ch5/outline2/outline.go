// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
func forEachNode(n *html.Node, pre func(n *html.Node, close bool), post func(n *html.Node)) {
	if pre != nil {
		if n.FirstChild == nil {
			pre(n, true)
		} else {
			pre(n, false)
		}
	}

	var count = 0
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		count++
		forEachNode(c, pre, post)
	}

	if post != nil {
		if count > 0 {
			post(n)
		}
	}
}

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node, close bool) {
	if n.Type == html.ElementNode {
		if len(n.Attr) > 0 {
			var s string
			for _, value := range n.Attr {
				s += fmt.Sprintf("%s ", value)
			}
			if close {
				fmt.Printf("%*s<%s %s/>\n", depth*2, "", n.Data, s)			
			} else {
				fmt.Printf("%*s<%s %s>\n", depth*2, "", n.Data, s)			
				depth++
			}
		} else {
			if close {
				fmt.Printf("%*s<%s/>\n", depth*2, "", n.Data)
			} else {
				fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
				depth++
			}
		}
	} else if n.Type == html.TextNode && n.Data != "\n" {
		fmt.Printf("%*s %s\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

//!-startend
