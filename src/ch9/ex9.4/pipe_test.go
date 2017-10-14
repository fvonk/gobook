package pipe_test

import (
  "ch9/ex9.4"
	"testing"
  "fmt"
  "time"
)

func BenchmarkPipeline(b *testing.B) {
	in, out := pipe.Pipeline(100000)
  start := time.Now()
	for i := 0; i < b.N; i++ {
		in <- 1
		<-out
	}
  fmt.Println("%s", time.Since(start))
	close(in)
}
