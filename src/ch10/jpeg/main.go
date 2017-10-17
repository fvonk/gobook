// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 287.

//!+main

// The jpeg command reads a PNG image from the standard input
// and writes it as a JPEG image to the standard output.
package main

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png" // register PNG decoder
	"image/gif"
	"flag"
	"io"
	"os"
)

func main() {
	outputKind := flag.String("output", "jpeg", "")
	flag.Parse()
	img, kind, err := image.Decode(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind, "Output format =", *outputKind)
	switch *outputKind {
	case "jpeg":
		err = toJPEG(img, os.Stdout)
	case "png":
		err = toPNG(img, os.Stdout)
	case "gif":
		err = toGif(img, os.Stdout)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", kind, err)
		os.Exit(1)
	}
}

func toJPEG(img image.Image, out io.Writer) error {
	return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
}

func toPNG(img image.Image, out io.Writer) error {
	return png.Encode(out, img)
}

func toGif(img image.Image, out io.Writer) error {
	return gif.Encode(out, img, &gif.Options{NumColors: 8})
}

//!-main

/*
//!+with
$ go build gopl.io/ch3/mandelbrot
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
Input format = png
//!-with

//!+without
$ go build gopl.io/ch10/jpeg
$ ./mandelbrot | ./jpeg >mandelbrot.jpg
jpeg: image: unknown format
//!-without
*/
