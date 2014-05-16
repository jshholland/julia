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

	if deg != 0 && f[deg] == 0 {
		return deg - 1
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
