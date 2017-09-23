// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"fmt"
	"log"
	"os"

	"ch5/links"
	"flag"
)

func crawl(url string, cancel <-chan struct{}) []string {
	select {
	case <-cancel:
		return nil
	default:
		fmt.Println(url)
		list, err := links.Extract(url, cancel)
		if err != nil {
			log.Print(err)
		}
		return list
	}
}

var depth = flag.Int("depth", 1, "depth")

type Link struct {
	string
	int
}

//!+
func main() {
	worklist := make(chan []Link)  // lists of URLs, may have duplicates
	unseenLinks := make(chan Link) // de-duplicated URLs

	flag.Parse()
	fmt.Println("depth value is", *depth)

	cancel := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(cancel)
	}()

	// Add command-line arguments to worklist.
	go func() {
		var links []Link
		for _, arg := range os.Args[1:] {
			links = append(links, Link{arg, 1})
		}
		worklist <- links
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for {
				select {
				case link := <-unseenLinks:
					if link.int <= *depth {
						var links []Link
						for _, crawlLink := range crawl(link.string, cancel) {
							links = append(links, Link{crawlLink, link.int + 1})
						}
						foundLinks := links
						go func() { worklist <- foundLinks }()
					}
				case <-cancel:
					return
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for {
		select {
		case list := <-worklist:
			for _, link := range list {
				if !seen[link.string] {
					seen[link.string] = true
					unseenLinks <- link
				}
			}
		case <-cancel:
			//panic("the end")
			return
		}
	}
}

//!-
