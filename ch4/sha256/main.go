// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"fmt"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
)

var (
	sha = flag.String("sha", "", "digest 512")
)

func main() {
	flag.Parse()

	for _, arg := range flag.Args() {
		fmt.Printf("arg = %s\n", arg)
		if len(*sha) > 0 {
			switch *sha {
			case "SHA256":
				fmt.Printf("SHA256: %x\n", sha256.Sum256([]byte(arg)))
				break
			case "SHA384":
				fmt.Printf("SHA384: %x\n", sha512.Sum384([]byte(arg)))
				break
			case "SHA512":
				fmt.Printf("SHA512: %x\n", sha512.Sum512([]byte(arg)))
				break
			default: return
			}	
		}			
	}

	// go run main.go --sha SHA512 a

	// c1 := sha256.Sum256([]byte("x"))
	// c2 := sha256.Sum256([]byte("y"))
	// fmt.Printf("%b\n%b\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// count := 0
	// for i := 0; i < len(c1); i++ {

	// 	b1 := c1[i]
 //        b2 := c2[i]
 //        // fmt.Printf("%b %b\n", b1, b2)
 //        for j := 0; j < 8; j++ {
 //            mask := byte(1 << uint(j))
 //            // fmt.Printf("%b %b %b\n", b1 & mask, b2 & mask, mask)
 //            if (b1 & mask) != (b2 & mask) {
 //                count++
 //            }
 //        }
	// }
	// fmt.Printf("%d\n", count)

	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}

//!-
