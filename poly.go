package main

import (
	"bytes"
	"fmt"
)

// pow computes z to the (nonnegative integer) power i.
func pow(z complex128, i int) complex128 {
	if i < 0 {
		panic("negative power")
	}

	var res complex128 = 1

	for i > 0 {
		res *= z
		i--
	}

	return res
}

// A Function can be evaluated at a complex point.
type Function interface {
	// evaluate computes the value of the function at z.
	Evaluate(z complex128) complex128
}

// A poly is a (complex) polynomial.
type poly []complex128

// Degree returns the degree of the polynomial.
//
// The degree of the zero polynomial is 0.
func (f poly) Degree() int {
	if len(f) == 0 {
		return 0
	}

	var deg int

	for i, v := range f {
		if v != 0 {
			deg = i
		}
	}

	return deg
}

// Normalise returns a new poly with any excess zeroes culled from the end.
func (f poly) Normalise() poly {
	deg := f.Degree()

	if deg == 0 && len(f) > 0 && f[0] == 0 {
		return poly([]complex128{})
	}

	if deg < len(f) {
		return f[:deg+1]
	}

	return f
}

// IsConstant checks whether the (normalised) polynomial is constant.
func (f poly) IsConstant() bool {
	return len(f) <= 1
}

func (f poly) Add(g poly) poly {
	deg1 := f.Degree()
	deg2 := g.Degree()
	var max_deg int

	if deg1 >= deg2 {
		max_deg = deg1
	} else {
		max_deg = deg2
	}

	res := make([]complex128, max_deg)

	for i := range res {
		var fi, gi complex128

		if i >= len(f) {
			fi = 0
		} else {
			fi = f[i]
		}

		if i >= len(g) {
			gi = 0
		} else {
			gi = g[i]
		}

		res[i] = fi + gi
	}

	return poly(res)
}

func (f poly) Negative() poly {
	res := make([]complex128, len(f))

	for i := range res {
		res[i] = -f[i]
	}

	return res
}

func (f poly) Subtract(g poly) poly {
	return f.Add(g.Negative())
}

func (f poly) Multiply(g poly) poly {
	deg := f.Degree() + g.Degree()

	res := make([]complex128, deg)

	for i := range res {
		for j := 0; j <= i; j++ {
			var fj, gi_j complex128

			if j >= len(f) {
				fj = 0
			} else {
				fj = f[j]
			}

			if i-j >= len(g) {
				gi_j = 0
			} else {
				gi_j = g[i-j]
			}

			res[i] += fj * gi_j
		}
	}

	return res
}

func (f poly) Equal(g poly) bool {
	deg1 := f.Degree()
	deg2 := g.Degree()

	if deg1 != deg2 {
		return false
	}

	if deg1 < len(f) {
		f = f[:deg1+1]
	}

	for i := range f {
		if f[i] != g[i] {
			return false
		}
	}

	return true
}

func (f poly) String() string {
	var b bytes.Buffer

	if len(f) == 0 {
		return "0"
	}

	b.WriteString(fmt.Sprint(f[0]))

	if len(f) >= 2 {
		part := fmt.Sprintf(" + %vt", f[1])
		b.WriteString(part)
	} else {
		return b.String()
	}

	for i, z := range f[2:] {
		part := fmt.Sprintf(" + %vt^%v", z, i+2)
		b.WriteString(part)
	}

	return b.String()
}

func (f poly) Evaluate(z complex128) complex128 {
	var res complex128

	for i, coeff := range f {
		res += coeff * pow(z, i)
	}

	return res
}

// NewPoly returns a Function coefficients coeffs, starting with the constant term.
func NewPoly(coeffs ...complex128) Function {
	return poly(coeffs)
}

// A rat is a rational function.
type rat struct {
	numerator   poly
	denominator poly
}

func (f rat) Evaluate(z complex128) complex128 {
	num := f.numerator.Evaluate(z)
	denom := f.denominator.Evaluate(z)
	return num / denom
}
