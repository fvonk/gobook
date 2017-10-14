package main

import (
  "fmt"
  "time"
)

func main() {
  start := time.Now()
  for i := 0; i < 10000; i++ {
    go fmt.Print(0)
    fmt.Print(1)
  }

  fmt.Printf("\ntime: %s\n", time.Since(start))
}
