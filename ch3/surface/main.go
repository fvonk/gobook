// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"log"
	"net/http"
	"strings"
	"strconv"
	"sync"	
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)
var colors 		  = [...]string {"#0000ff", "#006eff", "#00bbff", "#00ffe1", "#00ff88", "#62ff00", "#d0ff00", "#ffe100", "#ff7700", "#ff5500", "#ff0000"}
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)
var scale = 1
var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/scale/", reScale) 
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
func reScale(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.String(), "/scale/") {
		num, err := strconv.Atoi(strings.TrimPrefix(r.URL.String(), "/scale/"))
		if err != nil {
			// fmt.Fprintf(os.Stderr, "can't get number %d: %v\n", num, err)
		} else {
			scale = num
		}
	}
	w.Header().Set("ContentType", "image/svg+xml")
	getSVG(w)
	// fmt.Fprintf(w, "scale %d\n", scale)
}
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("ContentType", "image/svg+xml")
	scale = 1
	getSVG(w)
}

func getSVG(w http.ResponseWriter) {
	fmt.Fprintf(w, "<!DOCTYPE html><html><body><h1>SVG</h1><svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width * scale, height * scale)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, _ := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, color := corner(i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintf(w, "</svg></body></html>")
}

func corner(i, j int) (float64, float64, string) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, _ := f(x, y)
	z *= float64(scale)
	h := int(z * 100)
	// fmt.Printf("h %d", h);
	if h > 10 {
		h = 10
	} else if h < 0 {
		h = 0
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 * float64(scale) + (x-y)*cos30*xyscale * float64(scale)
	sy := height/2 * float64(scale) + (x+y)*sin30*xyscale * float64(scale) - z*zscale
	return sx, sy, colors[h]
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	if r != 0 {
		return math.Sin(r) / r, true
	} else {
		return 0, false
	}
}

//!-
