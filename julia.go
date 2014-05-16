package main

import (
	"fmt"
	"image/png"
	"os"
)

func main() {
	f := NewPoly(-1+.1i, 0, 1)
	m := Draw(f, -2, 2, -2, 2, 0.001)

	out, err := os.Create("julia.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	png.Encode(out, m)
}
