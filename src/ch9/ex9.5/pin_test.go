package pin_test

import (
  "ch9/ex9.5"
	"testing"
  "fmt"
  "time"
)

func Test(t *testing.T) {
  start := time.Now()
	pin.Pinpong()
  fmt.Println(">>> %s", time.Since(start))
}
