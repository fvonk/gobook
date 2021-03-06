// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	// "sort"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization", "programming languages"},
	"programming languages": {"data structures", "computer organization", "operating systems"},
}

//!-table

//!+main
func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func contains(mass []string, key string) (ok bool) {
	check := make(map[string]bool)
	for _, v := range mass {
		check[v] = true
	}
	_, ok = check[key]
	return
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	// var visitAll func(items []string)
	var visitAll func(items map[string][]string)
	visitAll = func(items map[string][]string) {
		for key, values := range items {
			if !seen[key] {
				seen[key] = true
				for _, value := range values {
					if contains(m[value], key) {
						fmt.Printf("cycle: %s and %s\n", key, value)
					}
					visitAll(map[string][]string{
							value: m[value],
						}) 
					if !seen[value] {
						order = append(order, value)
					}
				}
				order = append(order, key)
			}
		}
	}
	// visitAll = func(items []string) {
	// 	for _, item := range items {
	// 		if !seen[item] {
	// 			seen[item] = true
	// 			visitAll(m[item])
	// 			order = append(order, item)
	// 		}
	// 	}
	// }

	// var keys []string
	// for key := range m {
	// 	keys = append(keys, key)
	// }
	// sort.Strings(keys)
	// visitAll(keys)
	visitAll(m)
	return order
}

//!-main
