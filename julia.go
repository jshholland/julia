// Copyright © 2014 Josh Holland
//
// Part of julia, released under the MIT licence.
// Full source available at https://github.com/jshholland/julia

package main

import (
	"flag"
	"fmt"
	"image/png"
	"os"
)

var (
	re   = flag.Float64("real", 0.0, "real part of the constant whose quadratic Julia set to plot")
	im   = flag.Float64("imag", 0.0, "imaginary part of the constant whose quadratic Julia set to plot")
	size = flag.Int("size", 4000, "size in pixels to draw fractal at")
	out  = flag.String("out", "julia.png", "filename to save output to")
)

func main() {
	flag.Parse()
	f := NewPoly(complex(*re, *im), 0, 1)
	m := Draw(f, -2, 2, -2, 2, 4/float64(*size))

	out, err := os.Create(*out)
	if err != nil {
		fmt.Println(err)
		return
	}

	png.Encode(out, m)
}
