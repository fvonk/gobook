package pin

import (
  "sync"
  // "fmt"
)

func Pinpong() {
  in := make(chan int)
	out := make(chan int)

	var n sync.WaitGroup
  var count int

  n.Add(1)
  go func(in chan int, out chan int) {
    out <- 1
    for v := range in {
      // fmt.Println("count = %d", count)
      if count > 1000000 {
        break
      } else {
        out <- v
        count++
      }
    }
    close(out)
    close(in)
    n.Done()
  }(in, out)

  n.Add(1)
  go func(in chan int, out chan int) {
    for v := range in {
      out <- v
    }
    n.Done()
  }(out, in)

  n.Wait()
}
