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

// A Function can be evaluated at a complex point
type Function interface {
	// evaluate computes the value of the function at z
	Evaluate(z complex128) complex128
}

// A constant represents a constant function
type constant complex128

func (c constant) String() string {
	return fmt.Sprint(complex128(c))
}

func (f constant) Evaluate(z complex128) complex128 {
	return complex128(f)
}

// A poly is a (complex) polynomial.
type poly []complex128

func (f poly) String() string {
	var b bytes.Buffer

	if len(f) == 0 {
		return "0"
	}

	b.WriteString(fmt.Sprint(f[0]))

	if len(f) >= 2 {
		part := fmt.Sprintf(" + %vt", f[1])
		b.WriteString(part)
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

// NewPoly returns a Function coefficients coeffs, starting with the constant term
func NewPoly(coeffs ...complex128) Function {
	return poly(coeffs)
}

// A rat is a rational function
type rat struct {
	// coeffs for the numerator, starting with the constant term
	numerator poly
	// coeffs for the denominator, as for numerator
	denominator poly
}

func (f rat) Evaluate(z complex128) complex128 {
	num := f.numerator.Evaluate(z)
	denom := f.denominator.Evaluate(z)
	return num / denom
}
