// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"log"
	"net/http"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	// "math/rand"
	"strings"
	"strconv"
	// "fmt"	
	// "os"
)

var palette = []color.Color{color.Black, color.White}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		handler(w, r)	
		lissajous(w)	
	})
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// cycles
	// })
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "%s\n", r.URL.String())
	if strings.HasPrefix(r.URL.String(), "/?cycles=") {
		num, err := strconv.Atoi(strings.TrimPrefix(r.URL.String(), "/?cycles="))
		if err != nil {
			// fmt.Fprintf(os.Stderr, "can't get number %d: %v\n", num, err)
		} else {
			// fmt.Fprintf(w, "num: %v\n", num)
			nframes = num//float64()
			// return
		}
	}
	// fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// for k, v := range r.Header {
	// 	fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	// }
	// fmt.Fprintf(w, "Host = %q\n", r.Host)
	// fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	// if err := r.ParseForm(); err != nil {
	// 	log.Print(err)
	// }
	// for k, v := range r.Form {
	// 	fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	// }
}

var cycles float64 = 5     // number of complete x oscillator revolutions
var nframes = 64    // number of animation frames

func lissajous(out io.Writer) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		// nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
		temo = 234
	)
	freq := 10.0//rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference

	// for i := 1; i < nframes; i++ {
	// 	palette = append(palette, color.RGBA{uint8(i * 2 + 100), uint8(i * 3 + 50), uint8(i * 4), 1})
	// }

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// col := i % len(palette)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), 1)//uint8(i))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-handler
