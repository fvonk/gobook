package main

import (
	"fmt"
	// "io"
	"reflect"
)



func main() {
	defer func() {
		p := recover()
		if reflect.TypeOf(p).Kind() == reflect.Int {
			fmt.Printf("p = %d\n", p)
		} else {
			panic(p)
		}
	}()
	noReturn(5)
}

func noReturn(val int) {
	res := val * val
	panic(res)//fmt.Sprintf("%s", res))
}