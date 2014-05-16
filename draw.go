// Copyright © 2014 Josh Holland
//
// Part of julia, released under the MIT licence.
// Full source available at https://github.com/jshholland/julia

package main

import (
	"image"
	"image/color"
)

// coords handles conversion between pixels and complex values
type coords struct {
	pix_size float64
	origin   complex128
}

// At gets the complex number at the centre of the given pixel.
func (c coords) At(x, y int) complex128 {
	re := float64(x)*c.pix_size + real(c.origin)
	im := float64(y)*c.pix_size + imag(c.origin)
	return complex(re, im)
}

func Draw(f Function, re_from, re_to, im_from, im_to float64, pix_size float64) image.Image {
	w := re_to - re_from
	h := im_to - im_from
	if w < 0 || h < 0 {
		panic("negative width or height")
	}

	if pix_size < 0 {
		panic("negative pixel size")
	}

	c := coords{pix_size, complex(re_from, im_from)}

	resx := int(w / pix_size)
	resy := int(h / pix_size)

	m := image.NewGray16(image.Rect(0, 0, resx, resy))

	bounds := m.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			z := c.At(x, y)

			o := Orbit(f, z, 2, 500)

			if o < 0 {
				m.Set(x, y, color.Black)
			} else {
				m.Set(x, y, color.White)
			}
		}
	}

	return m
}
