package main 

import (
	"fmt"
	"os"
	// "bytes"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Printf("%t\n", isAnagram(os.Args[1], os.Args[2]))
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func isAnagram(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	var count = 0
	for i := 0; i < len(s1); i++ {
		if found := strings.LastIndex(s1, string(s2[i])); found >= 0 {
			count++
		}
	}
	// fmt.Printf("count %d\n", count)
	return count == len(s1)
}